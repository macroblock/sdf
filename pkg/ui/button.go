package ui

import (
	"image/color"

	"github.com/macroblock/sdf/pkg/geom"
)

type (
	// Button -
	Button struct {
		KernelNode
		pressed bool
		color   color.Color
	}
)

// NewButton -
func NewButton(bounds geom.Rect2i, c color.Color) *Button {
	return nil
}
