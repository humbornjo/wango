package filter

import (
	"image"
	"image/color"
)

func Invert(img image.Image) image.Image {
	rect := img.Bounds()
	for i := rect.Min.X; i < rect.Max.X; i++ {
		for j := rect.Min.Y; j < rect.Max.Y; j++ {
			clr := img.(*image.RGBA).RGBAAt(i, j)
			invertClr := color.RGBA{
				0xff - clr.R,
				0xff - clr.G,
				0xff - clr.B,
				clr.A,
			}
			img.(*image.RGBA).SetRGBA(i, j, invertClr)
		}
	}
	return img
}
