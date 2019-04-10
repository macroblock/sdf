package ui

import "github.com/macroblock/sdf/pkg/geom"

func shrinkToClientRect(rect, offsets geom.Rect2i) geom.Rect2i {
	rect.B.X -= rect.A.X + offsets.B.X
	rect.B.Y -= rect.A.Y + offsets.B.Y
	rect.A = offsets.A
	return rect
}
