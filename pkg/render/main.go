package render

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"sync"
	"unsafe"

	"github.com/humbornjo/wango/pkg/config"
	"github.com/humbornjo/wango/pkg/filter"
)

const (
	WIDTH  = 2560
	HEIGHT = 1536
	SIZE   = 256
)

// tile parts
//
// +---------+
// |+   t   +|
// |  +   +  |
// | l  *  r |
// |  +   +  |
// |+   b   +|
// +---------+

type Mask bool

type TileMask struct {
	b Mask
	l Mask
	t Mask
	r Mask
}

type Pattern uint8

type TilePattern struct {
	b Pattern
	l Pattern
	t Pattern
	r Pattern
}

func (tp *TilePattern) Hash() (hash uint32) {
	hash = (uint32(tp.b) << 24) | uint32(tp.l)<<16 | uint32(tp.t)<<8 | uint32(tp.r)
	return hash
}

func GenPattern(tilem TileMask, n int) (tilep TilePattern) {
	if n <= 0 {
		return tilep
	}
	tilep.b = (Pattern(config.Rng.Intn(n))) & (*(*Pattern)(unsafe.Pointer(&tilem.b)) - 1)
	tilep.l = (Pattern(config.Rng.Intn(n))) & (*(*Pattern)(unsafe.Pointer(&tilem.l)) - 1)
	tilep.t = (Pattern(config.Rng.Intn(n))) & (*(*Pattern)(unsafe.Pointer(&tilem.t)) - 1)
	tilep.r = (Pattern(config.Rng.Intn(n))) & (*(*Pattern)(unsafe.Pointer(&tilem.r)) - 1)
	return tilep
}

type Task struct {
	rect  image.Rectangle
	tilep TilePattern
}

type Wang struct {
	width, height int
	tile          Tile
	img           image.Image
	filters       []filter.Filter
	psize         int
	clrBg         color.RGBA
	tasks         chan Task
	cache         *sync.Map
}

type WangOption func(*Wang)

func InitWangWithOptions(options ...WangOption) (w Wang) {
	w.width = WIDTH
	w.height = HEIGHT
	w.cache = &sync.Map{}
	w.tile = Tile{SIZE, &MoistShader{}}
	for _, option := range options {
		option(&w)
	}
	w.tasks = make(chan Task, 10)
	w.img = image.NewRGBA(image.Rect(0, 0, w.width, w.height))
	return w
}

func WithWidth(width int) WangOption {
	return func(w *Wang) {
		w.width = width
	}
}

func WithHeight(height int) WangOption {
	return func(w *Wang) {
		w.height = height
	}
}

func WithBgColor(clr color.RGBA) WangOption {
	return func(w *Wang) {
		w.clrBg = clr
	}
}

func WithPatternSize(num int) WangOption {
	return func(w *Wang) {
		w.psize = num
	}
}

func WithSize(size int) WangOption {
	return func(w *Wang) {
		w.tile.size = size
	}
}

func WithShader(shader Shader) WangOption {
	return func(w *Wang) {
		w.tile.shader = shader
	}
}

func WithFilters(filters []filter.Filter) WangOption {
	return func(w *Wang) {
		w.filters = filters
	}
}

func (w *Wang) Map() {
	span := w.tile.size
	tw := w.width / w.tile.size
	th := w.height / w.tile.size
	patternGrid := make([][]TilePattern, th)
	for i := range th {
		patternGrid[i] = make([]TilePattern, tw)
	}

	patternGrid[0][0] = GenPattern(TileMask{}, w.psize)
	w.tasks <- Task{
		rect:  image.Rect(0, 0, span, span),
		tilep: patternGrid[0][0],
	}
	for j := 1; j < tw; j++ {
		tilep := GenPattern(TileMask{false, true, false, false}, w.psize)
		tilep.l = patternGrid[0][j-1].r
		patternGrid[0][j] = tilep
		w.tasks <- Task{
			rect:  image.Rect(j*span, 0, j*span+span, span),
			tilep: patternGrid[0][j],
		}
	}

	for i := 1; i < th; i++ {
		tilep := GenPattern(TileMask{false, false, true, false}, w.psize)
		tilep.t = patternGrid[i-1][0].b
		patternGrid[i][0] = tilep
		w.tasks <- Task{
			rect:  image.Rect(0, i*span, span, i*span+span),
			tilep: patternGrid[i][0],
		}
	}

	for i := 1; i < th; i++ {
		for j := 1; j < tw; j++ {
			tilep := GenPattern(TileMask{false, true, true, false}, w.psize)
			tilep.l = patternGrid[i][j-1].r
			tilep.t = patternGrid[i-1][j].b
			patternGrid[i][j] = tilep
			w.tasks <- Task{
				rect:  image.Rect(j*span, i*span, j*span+span, i*span+span),
				tilep: patternGrid[i][j],
			}
		}
	}
	close(w.tasks)
}

func (w *Wang) Reduce(peer int) {
	var wg sync.WaitGroup
	wg.Add(peer)
	for range peer {
		go func() {
			defer wg.Done()
			task, ok := <-w.tasks
			for ; ok; task, ok = <-w.tasks {
				w.Draw(
					w.img.(*image.RGBA).SubImage(task.rect).(*image.RGBA),
					task.tilep,
				)
			}
		}()
	}
	wg.Wait()

	for _, filter := range w.filters {
		w.img = filter(w.img)
	}
}

func (w *Wang) Save(path string, width, height int) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	err = png.Encode(f, w.img) // .SubImage(image.Rect(0, 0, width, height))
	if err != nil {
		return err
	}
	return nil
}

type Tile struct {
	size   int
	shader Shader
}

func (w *Wang) Draw(img *image.RGBA, tilep TilePattern) {
	rect := img.Bounds()
	posMin := rect.Min
	posMax := rect.Max
	width := posMax.X - posMin.X
	height := posMax.Y - posMin.Y

	if block, ok := w.cache.Load(tilep.Hash()); ok {
		span := width * 4
		srcImg := block.(*image.RGBA)
		srcPosMin := srcImg.Bounds().Min
		destImg := img
		destPosMin := posMin

		for i := range height {
			srcPixPtr := &destImg.Pix[srcImg.PixOffset(srcPosMin.X, srcPosMin.Y+i)]
			srcBytePtr := (*byte)(unsafe.Pointer(srcPixPtr))
			destPixPtr := &srcImg.Pix[destImg.PixOffset(destPosMin.X, destPosMin.Y+i)]
			destBytePtr := (*byte)(unsafe.Pointer(destPixPtr))
			copy(unsafe.Slice(srcBytePtr, span), unsafe.Slice(destBytePtr, span))
		}
		return
	}

	for j := range width {
		for i := range height {
			u := float64(j) / float64(width)
			v := float64(i) / float64(height)
			img.SetRGBA(
				j+posMin.X,
				i+posMin.Y,
				w.tile.shader.Render(Vec2f{u, v}, tilep, w.clrBg),
			)
		}
	}
	w.cache.Store(tilep.Hash(), img)
}
