package ui

import "github.com/macroblock/sdf/pkg/geom"

func shrink(rect, offsets geom.Rect2i) geom.Rect2i {
	rect.A.X += offsets.A.X
	rect.A.Y += offsets.A.Y
	rect.B.X -= offsets.A.X
	rect.B.Y -= offsets.A.Y
	return rect
}
