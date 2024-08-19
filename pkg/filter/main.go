package filter

import (
	"fmt"
	"image"
)

type Filter func(image.Image) image.Image

func FilterSelect(name string) (Filter, error) {
	switch name {
	case "identical":
		return Identical, nil
	case "invert":
		return Invert, nil
	case "sepia":
		return Sepia, nil
	default:
		return nil, fmt.Errorf("undefined filter")
	}
}
