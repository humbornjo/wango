package render

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"runtime"
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
	wang := InitWangWithOptions(
		WithWidth(width),
		WithHeight(height),
		WithSize(SIZE),
		WithBgColor(color.RGBA{}),
		WithShader(DefaultShader),
		WithNumColor(DefaultClrNum),
	)

	go wang.Map()
	wang.Reduce(runtime.NumCPU())
	defaultSave("../../single.png", wang.img)
}

func TestGrid(t *testing.T) {
	var width = WIDTH
	var height = HEIGHT
	wang := InitWangWithOptions(
		WithWidth(width),
		WithHeight(height),
		WithSize(SIZE),
		WithBgColor(color.RGBA{}),
		WithShader(DefaultShader),
		WithNumColor(DefaultClrNum),
	)

	go wang.Map()
	wang.Reduce(runtime.NumCPU())
	defaultSave("../../grid.png", wang.img)
}
