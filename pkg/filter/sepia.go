package filter

import (
	"image"
	"image/color"
)

func Sepia(img image.Image) image.Image {
	bounds := img.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			rgba := img.(*image.RGBA).RGBAAt(x, y)

			// Apply sepia filter
			newR := uint8(0.393*float32(rgba.R) + 0.769*float32(rgba.G) + 0.189*float32(rgba.B))
			newG := uint8(0.349*float32(rgba.R) + 0.686*float32(rgba.G) + 0.168*float32(rgba.B))
			newB := uint8(0.272*float32(rgba.R) + 0.534*float32(rgba.G) + 0.131*float32(rgba.B))

			// Clamp values to 0-255
			newR = clamp(newR, 0, 255)
			newG = clamp(newG, 0, 255)
			newB = clamp(newB, 0, 255)

			img.(*image.RGBA).SetRGBA(x, y, color.RGBA{newR, newG, newB, rgba.A})
		}
	}
	return img
}

func clamp(v, min, max uint8) uint8 {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}
