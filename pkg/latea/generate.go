package latea

import (
	"fmt"
	"image/color"
	"runtime"
	"strconv"
	"strings"

	"github.com/humbornjo/wango/pkg/config"
	"github.com/humbornjo/wango/pkg/filter"
	"github.com/humbornjo/wango/pkg/render"
)

var (
	Success       = false
	Err     error = nil
	Path    string
)

func (m *model) Generate() {
	var (
		err                 error
		wang                render.Wang
		imgWidth, imgHeight int
		size, width, height int
		mode, path          string
		color               color.RGBA
		shader              render.Shader
		filters             []filter.Filter
	)

	width, err = ParseInt(inputWidth.Value(), config.WIDTH)
	if err != nil {
		goto ret
	}

	height, err = ParseInt(inputHeight.Value(), config.HEIGHT)
	if err != nil {
		goto ret
	}

	size, err = ParseInt(inputSize.Value(), config.SIZE)
	if err != nil {
		goto ret
	}

	color, err = ParseColor(inputClrBg.Value(), config.ClrBackground)
	if err != nil {
		goto ret
	}

	path, err = ParseStr(inputPath.Value(), config.PATH)
	if err != nil {
		goto ret
	}

	for _, choice := range config.ChoicesShader {
		if choice.Choosen {
			sdr, err := render.ShaderSelect(
				strings.TrimSpace(choice.Label),
			)
			if err != nil {
				goto ret
			}
			shader = sdr
			break
		}
	}

	for _, choice := range config.ChoicesFilter {
		if choice.Choosen {
			flt, err := filter.FilterSelect(
				strings.TrimSpace(choice.Label),
			)
			if err != nil {
				goto ret
			}
			filters = append(filters, flt)
		}
	}

	for _, choice := range config.ChoicesMode {
		if choice.Choosen {
			mode = strings.TrimSpace(choice.Label)
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
		render.WithFilters(filters),
		render.WithPatternSize(len(config.Palette)),
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

func ParseInt(str string, def int) (int, error) {
	if str == "" {
		return def, nil
	}
	return strconv.Atoi(str)
}

func ParseStr(str string, def string) (string, error) {
	if str == "" {
		str = def
	}
	return str, nil
}

func ParseColor(str string, def color.RGBA) (color.RGBA, error) {
	var err error
	if str == "" {
		return def, nil
	}
	c := color.RGBA{0, 0, 0, 0xff}
	switch len(str) {
	case 7:
		_, err = fmt.Sscanf(str, "#%02x%02x%02x", &c.R, &c.G, &c.B)
	case 4:
		_, err = fmt.Sscanf(str, "#%1x%1x%1x", &c.R, &c.G, &c.B)
		// Double the hex digits:
		c.R *= 17
		c.G *= 17
		c.B *= 17
	default:
		err = fmt.Errorf("invalid hex length, must be 7 or 4")
	}
	return c, err
}

func CeilX(x int, y int) int {
	if x < y {
		return y
	}
	rem := x % y
	if rem == 0 {
		return x
	}
	return x + y - rem
}

func FloorX(x int, y int) int {
	if x < y {
		return y
	}
	rem := x % y
	if rem == 0 {
		return x
	}
	return x - rem
}
