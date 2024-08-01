package render

import (
	"image/color"
	"math"
)

type FooShader struct{ palette color.Palette }

// bottom left top right
func parsePattern(pattern uint8) (int, int, int, int) {
	var b, l, t, r int
	b = int(pattern | 0b11)
	l = int((pattern >> 2) | 0b11)
	t = int((pattern >> 4) | 0b11)
	r = int((pattern >> 6) | 0b11)
	return b, l, t, r
}

func (s *FooShader) Render(u, v float64, pattern uint8) color.RGBA {
	return color.RGBA{uint8(u * 0xff), uint8(v * 0xff), 0, 0xff}
}

type MoistShader struct{ palette color.Palette }

func (s *MoistShader) Render(u, v float64, pattern uint8) {
	b, l, t, r := parsePattern(pattern)
	sides := [][]float64{{0.5, 1}, {0, 0.5}, {0.5, 0}, {1, 0.5}}

}

type JapanShader struct{}

func (s *JapanShader) Render(u, v float64, pattern uint8) color.RGBA {
	dist := math.Sqrt(
		(u-0.5)*(u-0.5) + (v-0.5)*(v-0.5),
	)
	var mask uint8 = 0xff
	if dist < 0.3 {
		mask = 0
	}

	return color.RGBA{0xff, mask & 0xff, mask & 0xff, 0xff}
}
