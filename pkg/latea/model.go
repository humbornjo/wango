package latea

import (
	"image/color"
)

type WinSize struct {
	width, height int
}

type Palette struct {
	clrs []color.RGBA
}

func InitDefaultPalette() Palette {
	return Palette{
		clrs: []color.RGBA{
			color.RGBA{0xff, 0, 0, 0xff},
			color.RGBA{0, 0, 0, 0xff},
		},
	}
}

type model struct {
	width, height int
	path          string
	mode          string
	texture       string
	palette       Palette
	seed          uint
	ws            WinSize
}

func InitDefaultModel() model {
	return model{
		width:   1600,
		height:  1000,
		texture: "vanilla",
		palette: InitDefaultPalette(),
		path:    "./",
		seed:    3407,
		ws:      WinSize{},
	}
}
