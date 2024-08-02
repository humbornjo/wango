package render

import (
	"image/color"
)

// bottom left top right
func parsePattern(pattern uint8) []int {
	var b, l, t, r int
	b = int(pattern & 0b11)
	l = int((pattern >> 2) & 0b11)
	t = int((pattern >> 4) & 0b11)
	r = int((pattern >> 6) & 0b11)
	return []int{b, l, t, r}
}

type FooShader struct{ palette color.Palette }

func (s *FooShader) Render(u, v float64, pattern uint8) color.RGBA {
	return color.RGBA{uint8(u * 0xff), uint8(v * 0xff), 0, 0xff}
}

type JapanShader struct{}

func (s *JapanShader) Render(u, v float64, pattern uint8) color.RGBA {
	dist := Vec2f{u, v}.Dist(Vec2f{0.5, 0.5})
	var mask uint8 = 0xff
	if dist < 0.3 {
		mask = 0
	}

	return color.RGBA{0xff, mask & 0xff, mask & 0xff, 0xff}
}

type MoistShader struct {
	palette color.Palette
}

func (s *MoistShader) Render(p Vec2f, pattern uint8, clr color.RGBA) color.RGBA {
	rgba := Vec4f{float64(clr.R), float64(clr.G), float64(clr.B), float64(clr.A)}.Div(0xff)
	radius := 0.4
	bltr := parsePattern(pattern)
	sides := []Vec2f{{0.5, 1}, {0, 0.5}, {0.5, 0}, {1, 0.5}}

	for i, side := range sides {
		if p.Dist(side) >= radius {
			continue
		}
		t := 1 - p.Dist(side)/radius
		pclr := s.palette[bltr[i]%len(s.palette)].(color.RGBA)
		pclr4f := Vec4f{float64(pclr.R), float64(pclr.G), float64(pclr.B), float64(pclr.A)}.Div(0xff)
		rgba = rgba.Fmaxf(rgba.Lerp(pclr4f, t)) // TODO: how to blend color
	}
	rgba = rgba.Gamma(1.0 / 2.2).Mul(0xff)
	return color.RGBA{uint8(rgba.X), uint8(rgba.Y), uint8(rgba.Z), uint8(rgba.A)}
}
