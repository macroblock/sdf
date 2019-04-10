package gfx

import "image/color"

// FlipMode -
type FlipMode int

// can be ORed
const (
	FlipNone FlipMode = iota
	FlipHorizontal
	FlipVertical
)

type (
	// IFont -
	IFont interface {
		Print(x0, y0 int, text string)
		SetColor(c color.RGBA)
	}
)
