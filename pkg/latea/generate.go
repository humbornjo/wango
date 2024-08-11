package latea

import (
	"image/color"
	"runtime"
	"strconv"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/humbornjo/wango/pkg/config"
	"github.com/humbornjo/wango/pkg/render"
)

var (
	Success       = false
	Err     error = nil
	Path    string
)

func (m *model) Generate() {
	width, err := ParseInt(&inputWidth, config.WIDTH)
	if err != nil {
		panic("fail parse width")
	}

	height, err := ParseInt(&inputHeight, config.HEIGHT)
	if err != nil {
		panic("fail parse height")
	}

	size, err := ParseInt(&inputSize, config.SIZE)
	if err != nil {
		panic("fail parse size")
	}

	color, err := ParseColor(&inputClrBg, config.ClrBackground)
	if err != nil {
		panic("fail parse color")
	}

	path, err := ParseStr(&inputPath, config.PATH)
	if err != nil {
		panic("fail parse path")
	}

	var imgWidth int
	var imgHeight int
	var mode string
	var shader render.Shader

	for _, choice := range config.ChoicesShader {
		if choice.Choosen {
			shader = render.DefaultShader
			break
		}
	}

	for _, choice := range config.ChoicesMode {
		if choice.Choosen {
			mode = "up"
			break
		}
	}

	switch mode {
	case "up":
		imgWidth = CeilX(width, size)
		imgHeight = CeilX(height, size)
		m.width = imgWidth
		m.height = imgHeight
	case "down":
		imgWidth = FloorX(width, size)
		imgHeight = FloorX(height, size)
		m.width = imgWidth
		m.height = imgHeight
	case "exact":
		imgWidth = CeilX(width, size)
		imgHeight = CeilX(height, size)
		m.width = width
		m.height = height
	}

	wang := render.InitWangWithOptions(
		render.WithWidth(imgWidth),
		render.WithHeight(imgHeight),
		render.WithSize(size),
		render.WithBgColor(color),
		render.WithShader(shader),
		render.WithNumColor(render.DefaultClrNum),
	)

	go wang.Map()
	wang.Reduce(runtime.NumCPU())
	err = wang.Save(path, m.width, m.height)
	if err != nil {
		Err = err
		return
	}
	Path = path
	Success = true
}

func ParseInt(ti *textinput.Model, def int) (int, error) {
	str := ti.Value()
	if str == "" {
		return def, nil
	}
	return strconv.Atoi(str)
}

func ParseColor(ti *textinput.Model, def color.RGBA) (color.RGBA, error) {
	str := ti.Value()
	if str == "" {
		return def, nil
	}
	return color.RGBA{}, nil
}

func ParseStr(ti *textinput.Model, def string) (string, error) {
	str := ti.Value()
	if str == "" {
		str = def
	}
	return str, nil
}

func CeilX(x int, y int) int {
	if x < y {
		return y
	}
	remainder := x % y
	if remainder == 0 {
		return x
	}
	return x + y - remainder
}

func FloorX(x int, y int) int {
	if x < y {
		return y
	}
	remainder := x % y
	if remainder == 0 {
		return x
	}
	return x - remainder
}
