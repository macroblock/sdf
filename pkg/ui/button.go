package ui

import (
	"image/color"

	"github.com/macroblock/sdf/pkg/geom"
)

type (
	// InnerNode -
	InnerNode struct {
		bounds geom.Rect2i
		client geom.Rect2i
		parent *InnerNode
		childs []interface{}
		self   interface{}
	}
	// Button -
	Button struct {
		InnerNode
		pressed bool
		color   color.Color
	}
)

// NewButton -
func NewButton(bounds geom.Rect2i, c color.Color) *Button {
	return nil
}
