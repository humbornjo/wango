package render

import (
	"image"
	"image/color"
	"math/rand"
	"sync"
	"unsafe"
)

const (
	WIDTH  = 2048
	HEIGHT = 1536
	SIZE   = 256
)

var DefaultShader = &MoistShader{DefaultPalette}
var DefaultClrNum = len(DefaultPalette)
var DefaultPalette = color.Palette{
	color.RGBA{0xff, 0, 0, 0xff},
	color.RGBA{0, 0xff, 0xff, 0xff},
}

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
	hash |= uint32(tp.b) << 24
	hash |= uint32(tp.l) << 16
	hash |= uint32(tp.t) << 8
	hash |= uint32(tp.r)
	return hash
}

func GenPattern(tilem TileMask, n int) (tilep TilePattern) {
	tilep.b = (Pattern(rand.Intn(n))) &
		(*(*Pattern)(unsafe.Pointer(&tilem.b)) - 1)
	tilep.l = (Pattern(rand.Intn(n))) &
		(*(*Pattern)(unsafe.Pointer(&tilem.l)) - 1)
	tilep.t = (Pattern(rand.Intn(n))) &
		(*(*Pattern)(unsafe.Pointer(&tilem.t)) - 1)
	tilep.r = (Pattern(rand.Intn(n))) &
		(*(*Pattern)(unsafe.Pointer(&tilem.r)) - 1)
	return tilep
}

type Task struct {
	rect  image.Rectangle
	tilep TilePattern
}

type Wang struct {
	width, height int
	tile          Tile
	img           *image.RGBA
	clrNum        int
	clrBg         color.RGBA
	tasks         chan Task
	cache         *sync.Map
}

type WangOption func(*Wang)

func InitWangWithOptions(options ...WangOption) (w Wang) {
	w.width = WIDTH
	w.height = HEIGHT
	w.tile = Tile{SIZE, DefaultShader}
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

func WithSize(size int) WangOption {
	return func(w *Wang) {
		w.tile.size = size
	}
}

func WithBgColor(clr color.RGBA) WangOption {
	return func(w *Wang) {
		w.clrBg = clr
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

	patternGrid[0][0] = GenPattern(TileMask{}, w.clrNum)
	for j := 1; j < tw; j++ {
		tilep := GenPattern(
			TileMask{false, true, false, false},
			w.clrNum,
		)
		tilep.l = patternGrid[0][j-1].r
		patternGrid[0][j] = tilep
	}

	for i := 1; i < th; i++ {
		tilep := GenPattern(
			TileMask{false, false, true, false},
			w.clrNum,
		)
		tilep.t = patternGrid[i-1][0].b
		patternGrid[i][0] = tilep
	}

	for i := 1; i < th; i++ {
		for j := 1; j < tw; j++ {
			tilep := GenPattern(
				TileMask{false, true, true, false},
				w.clrNum,
			)
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
					w.img.SubImage(task.rect).(*image.RGBA),
					task.tilep,
				)
			}
		}()
	}
	wg.Wait()
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
