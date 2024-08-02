package render

import (
	"image/color"
)

// bottom left top right
func parsePattern(pattern uint8) [4]int {
	var b, l, t, r int
	b = int(pattern & 0b11)
	l = int((pattern >> 2) & 0b11)
	t = int((pattern >> 4) & 0b11)
	r = int((pattern >> 6) & 0b11)
	return [4]int{b, l, t, r}
}

func rgbaToVec4f(clr *color.RGBA) Vec4f {
	return Vec4f{float64(clr.R), float64(clr.G), float64(clr.B), float64(clr.A)}
}

type FooShader struct{}

func (s *FooShader) Render(p Vec2f, pattern uint8, clr color.RGBA) color.RGBA {
	return color.RGBA{uint8(p.X * 0xff), uint8(p.Y * 0xff), 0, 0xff}
}

type JapanShader struct{}

func (s *JapanShader) Render(p Vec2f, pattern uint8, clr color.RGBA) color.RGBA {
	dist := p.Dist(Vec2f{0.5, 0.5}, 2)
	var mask uint8 = 0xff
	if dist < 0.3 {
		mask = 0
	}

	return color.RGBA{0xff, mask & 0xff, mask & 0xff, 0xff}
}

type MoistShader struct {
	palette color.Palette
}

func (s *MoistShader) Render(p Vec2f, pattern uint8, bgclr color.RGBA) color.RGBA {
	rgba := rgbaToVec4f(&bgclr).Div(0xff)
	radius := 0.5
	bltr := parsePattern(pattern)
	sides := []Vec2f{{0.5, 1}, {0, 0.5}, {0.5, 0}, {1, 0.5}}

	for i, side := range sides {
		t := 1 - min(p.Dist(side, 2)/radius, 1)
		clr := s.palette[bltr[i]%len(s.palette)].(color.RGBA)
		clr4f := rgbaToVec4f(&clr).Div(0xff)
		rgba = rgba.Fmax(rgba.Lerp(clr4f, t)) // TODO: how to blend color
	}

	rgba = rgba.Gamma(1.0 / 2.2).Mul(0xff)
	return color.RGBA{uint8(rgba.X), uint8(rgba.Y), uint8(rgba.Z), uint8(rgba.A)}
}
