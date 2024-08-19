package render

import (
	"fmt"
	"image/color"
	"unsafe"

	"github.com/humbornjo/wango/pkg/config"
)

func ShaderSelect(name string) (Shader, error) {
	switch name {
	case "moist":
		return &MoistShader{}, nil
	default:
		return nil, fmt.Errorf("undefined shader")
	}
}

func (v *Vec4f) FromRGBA(clr *color.RGBA) *Vec4f {
	return &Vec4f{float64(clr.R), float64(clr.G), float64(clr.B), float64(clr.A)}
}

func (v *Vec4f) ToRGBA() color.RGBA {
	return color.RGBA{uint8(v.X), uint8(v.Y), uint8(v.Z), uint8(v.A)}
}

type Shader interface {
	Render(Vec2f, TilePattern, color.RGBA) color.RGBA
}

// foo shader
type FooShader struct{}

func (s *FooShader) Render(p Vec2f, pattern TilePattern, clr color.RGBA) color.RGBA {
	return color.RGBA{uint8(p.X * 0xff), uint8(p.Y * 0xff), 0, 0xff}
}

// japan shader
type JapanShader struct{}

func (s *JapanShader) Render(p Vec2f, pattern TilePattern, clr color.RGBA) color.RGBA {
	radius := 0.3
	flag := p.Dist(Vec2f{0.5, 0.5}, 2) < radius
	mask := *(*uint8)(unsafe.Pointer(&flag)) - 1
	return color.RGBA{0xff, mask & 0xff, mask & 0xff, 0xff}
}

// moist shader
type MoistShader struct{}

func (s *MoistShader) Render(p Vec2f, tilep TilePattern, bgclr color.RGBA) color.RGBA {
	radius := 0.5
	rgba4f := (&Vec4f{}).FromRGBA(&bgclr).Div(0xff)
	bltr := []Pattern{tilep.b, tilep.l, tilep.t, tilep.r}
	sides := []Vec2f{{0.5, 1}, {0, 0.5}, {0.5, 0}, {1, 0.5}}

	for i, side := range sides {
		t := 1 - min(p.Dist(side, 2)/radius, 1)
		clr := config.Palette[bltr[i]].(*color.RGBA)
		clr4f := (&Vec4f{}).FromRGBA(clr).Div(0xff)
		rgba4f = rgba4f.Fmax(rgba4f.Lerp(clr4f, t))
	}

	rgba4f = rgba4f.Gamma(1.0 / 2.2).Mul(0xff)
	return rgba4f.ToRGBA()
}

// path shader
type PathShader struct{}

func (s *PathShader) Render(p Vec2f, tilep TilePattern, bgclr color.RGBA) color.RGBA {
	// Read block get meta data width, height
	// return color.RGBA(img.get(width*p.X, height.Y))
	return color.RGBA{}
}
