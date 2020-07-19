package rgba32

import (
	"image"
	"image/color"
)

func (receiver Slice) At(x, y int) color.Color {
	return receiver
}

func (receier Slice) Bounds() image.Rectangle {
	const max int = int(^uint(0) >> 1)
	const min int = (-max - 1)

	return image.Rectangle{
		Min: image.Point{
			X: min,
			Y: min,
		},
		Max: image.Point{
			X: max,
			Y: max,
		},
	}
}

func (receiver Slice) ColorModel() color.Model {
	return color.NRGBAModel
}
