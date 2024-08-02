package render

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"runtime"
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

func TestSingleBlock(t *testing.T) {
	var w = 512
	var h = 512
	wang := Wang{
		w, h,
		Tile{512, defaultShader, color.RGBA{0, 0, 0, 0xff}},
		*image.NewRGBA(image.Rect(0, 0, w, h)),
		make(chan image.Rectangle, 10),
	}

	WithBgColor(color.RGBA{0, 0, 0, 0xff})(&wang)

	go wang.Map()
	wang.Reduce(runtime.NumCPU())
	defaultSave("../../public/single.png", &wang.img)
}

func TestAtlas(t *testing.T) {
	var w = 2048
	var h = 2048
	wang := Wang{
		w, h,
		Tile{512, defaultShader, color.RGBA{0, 0, 0, 0xff}},
		*image.NewRGBA(image.Rect(0, 0, w, h)),
		make(chan image.Rectangle, 10),
	}

	go wang.Map()
	wang.Reduce(runtime.NumCPU())
	defaultSave("../../public/atlas.png", &wang.img)
}
