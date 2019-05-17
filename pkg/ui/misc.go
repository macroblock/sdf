package ui

import "image"

func shrink(rect, offsets image.Rectangle) image.Rectangle {
	rect.Min.X += offsets.Min.X
	rect.Min.Y += offsets.Min.Y
	rect.Max.X -= offsets.Min.X
	rect.Max.Y -= offsets.Min.Y
	return rect
}
