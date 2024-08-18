package latea

import (
	"fmt"
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
	var err error
	var wang render.Wang
	var imgWidth, imgHeight int
	var size, width, height int
	var mode, path string
	var color color.RGBA
	var shader render.Shader

	width, err = ParseInt(&inputWidth, config.WIDTH)
	if err != nil {
		goto ret
	}

	height, err = ParseInt(&inputHeight, config.HEIGHT)
	if err != nil {
		goto ret
	}

	size, err = ParseInt(&inputSize, config.SIZE)
	if err != nil {
		goto ret
	}

	color, err = ParseColor(&inputClrBg, config.ClrBackground)
	if err != nil {
		goto ret
	}

	path, err = ParseStr(&inputPath, config.PATH)
	if err != nil {
		goto ret
	}

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

	wang = render.InitWangWithOptions(
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

ret:
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

func ParseStr(ti *textinput.Model, def string) (string, error) {
	str := ti.Value()
	if str == "" {
		str = def
	}
	return str, nil
}

func ParseColor(ti *textinput.Model, def color.RGBA) (color.RGBA, error) {
	var err error
	s := ti.Value()
	if s == "" {
		return def, nil
	}

	c := def
	switch len(s) {
	case 7:
		_, err = fmt.Sscanf(s, "#%02x%02x%02x", &c.R, &c.G, &c.B)
	case 4:
		_, err = fmt.Sscanf(s, "#%1x%1x%1x", &c.R, &c.G, &c.B)
		// Double the hex digits:
		c.R *= 17
		c.G *= 17
		c.B *= 17
	default:
		err = fmt.Errorf("invalid length, must be 7 or 4")

	}
	return c, err
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
