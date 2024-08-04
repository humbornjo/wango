package render

import (
	"image"
	"image/color"
	"math/rand"
	"sync"
	"unsafe"
)

// render img parallely using map-reduce
type ParaWang interface {
	Map()
	Reduce(peer int)
}

type Wang struct {
	width, height int
	tile          Tile
	img           *image.RGBA
	bgclr         color.RGBA
	tasks         chan image.Rectangle
	cache         *sync.Map
}

type WangOption func(*Wang)

func InitWangWithOptions(width, height, size int, options ...WangOption) (w Wang) {
	w.width = width
	w.height = height
	w.img = image.NewRGBA(image.Rect(0, 0, width, height))
	w.tasks = make(chan image.Rectangle, 10)
	w.tile = Tile{size, nil}
	for _, option := range options {
		option(&w)
	}
	return w
}

func WithBgColor(clr color.RGBA) WangOption {
	return func(w *Wang) {
		w.bgclr = clr
	}
}

func (w *Wang) Map() {
	stride := w.tile.size
	for j := 0; j < w.height; j += stride {
		for i := 0; i < w.width; i += stride {
			w.tasks <- image.Rect(i, j, i+stride, j+stride)
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
					w.img.SubImage(task).(*image.RGBA),
					uint8(((rand.Uint32()%2)<<6)|((rand.Uint32()%2)<<4)|((rand.Uint32()%2)<<2)|(rand.Uint32()%2)),
				)
			}
		}()
	}
	wg.Wait()
}

type Shader interface {
	Render(Vec2f, uint8, color.RGBA) color.RGBA
}

type Tile struct {
	size   int
	shader Shader
}

func (w *Wang) Draw(img *image.RGBA, pattern uint8) {
	rect := img.Bounds()
	posMin := rect.Min
	posMax := rect.Max
	width := posMax.X - posMin.X
	height := posMax.Y - posMin.Y

	if block, ok := w.cache.Load(pattern); ok {
		stride := width * 4

		srcImg := block.(*image.RGBA)
		srcPosMin := srcImg.Bounds().Min
		destImg := img
		destPosMin := posMin

		for j := range height {
			srcPixPtr := &destImg.Pix[destImg.PixOffset(destPosMin.X, destPosMin.Y+j)]
			srcBytePtr := (*byte)(unsafe.Pointer(srcPixPtr))
			destPixPtr := &srcImg.Pix[srcImg.PixOffset(srcPosMin.X, srcPosMin.Y+j)]
			destBytePtr := (*byte)(unsafe.Pointer(destPixPtr))
			copy(unsafe.Slice(srcBytePtr, stride), unsafe.Slice(destBytePtr, stride))
		}
		return
	}

	for i := range width {
		for j := range height {
			u := float64(i) / float64(width)
			v := float64(j) / float64(height)
			img.SetRGBA(i+posMin.X, j+posMin.Y, w.tile.shader.Render(Vec2f{u, v}, pattern, w.bgclr))
		}
	}
	w.cache.Store(pattern, img)
}
