package latea

import (
	"image/color"

	"github.com/humbornjo/wango/pkg/render"
)

type WinSize struct {
	width, height int
}

type Palette struct {
	clrs []color.RGBA
}

type model struct {
	wang          render.Wang
	width, height int
	path          string
	mode          string
	texture       string
	palette       Palette
	seed          int
	winsize       WinSize
}

func InitModel() model {
	return model{}
}
