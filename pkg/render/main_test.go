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

// var defaultShader = &FooShader{defaultPalette()}
var defaultShader = &MoistShader{defaultPalette()}

func defaultPalette() color.Palette {
	return color.Palette{
		color.RGBA{0xff, 0, 0, 0xff},
		color.RGBA{0, 0xff, 0xff, 0xff},
		// color.RGBA{0, 0, 0, 0xff},
	}
}

func defaultSave(path string, img image.Image) {
	f, err := os.Create(path)
	if err != nil {
		panic("SaveDefault")
	}
	defer f.Close()
	png.Encode(f, img)
}

func TestSingle(t *testing.T) {
	var width = 512
	var height = 512
	w := Wang{
		width, height,
		Tile{512, defaultShader},
		image.NewRGBA(image.Rect(0, 0, width, height)),
		color.RGBA{0, 0, 0, 0xff},
		make(chan image.Rectangle, 10),
		&sync.Map{},
	}

	WithBgColor(color.RGBA{0, 0, 0, 0xff})(&w)

	go w.Map()
	w.Reduce(runtime.NumCPU())
	defaultSave("../../single.png", w.img)
}

func TestAtlas(t *testing.T) {
	var width = 8000
	var height = 8000
	w := Wang{
		width, height,
		Tile{100, defaultShader},
		image.NewRGBA(image.Rect(0, 0, width, height)),
		color.RGBA{0, 0, 0, 0xff},
		make(chan image.Rectangle, 10),
		&sync.Map{},
	}

	go w.Map()
	w.Reduce(runtime.NumCPU())
	defaultSave("../../atlas.png", w.img)
}
