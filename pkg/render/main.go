package render

import (
	"image"
	"image/color"
	"math/rand"
	"sync"
)

// render img parallely using map-reduce
type ParaWang interface {
	Map()
	Reduce(peer int)
}

type Wang struct {
	width, height int
	tile          Tile
	img           image.RGBA
	tasks         chan image.Rectangle
}

type WangOption func(*Wang)

func WithBgColor(clr color.RGBA) WangOption {
	return func(w *Wang) {
		for i := range w.width {
			for j := range w.height {
				w.img.SetRGBA(i, j, clr)
			}
		}
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
				w.tile.Draw(w.img.SubImage(task).(*image.RGBA))
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
	bgclr  color.RGBA
}

func (t *Tile) Draw(img *image.RGBA) {
	rect := img.Bounds()
	posMin := rect.Min
	posMax := rect.Max
	w := posMax.X - posMin.X
	h := posMax.Y - posMin.Y

	pattern := uint8(rand.Uint32())

	for i := range w {
		for j := range h {
			u := float64(i) / float64(w)
			v := float64(j) / float64(h)
			img.SetRGBA(
				i+posMin.X,
				j+posMin.Y,
				t.shader.Render(Vec2f{u, v}, pattern, t.bgclr),
			)
		}
	}
}
