package render

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"runtime"
	"sync"
	"testing"
)

func defaultSave(path string, img image.Image) {
	f, err := os.Create(path)
	if err != nil {
		panic("SaveDefault")
	}
	defer f.Close()
	png.Encode(f, img)
}

func TestSingle(t *testing.T) {
	var width = SIZE
	var height = SIZE
	w := Wang{
		width, height,
		Tile{SIZE, DefaultShader},
		image.NewRGBA(image.Rect(0, 0, width, height)),
		defaultClrNum,
		color.RGBA{0, 0, 0, 0xff},
		make(chan Task, 10),
		&sync.Map{},
	}

	WithBgColor(color.RGBA{0, 0, 0, 0xff})(&w)

	go w.Map()
	w.Reduce(runtime.NumCPU())
	defaultSave("../../single.png", w.img)
}

func TestGrid(t *testing.T) {
	var width = WIDTH
	var height = HEIGHT
	w := Wang{
		width, height,
		Tile{SIZE, DefaultShader},
		image.NewRGBA(image.Rect(0, 0, width, height)),
		defaultClrNum,
		color.RGBA{0, 0, 0, 0xff},
		make(chan Task, 10),
		&sync.Map{},
	}

	go w.Map()
	w.Reduce(runtime.NumCPU())
	defaultSave("../../atlas.png", w.img)
}
