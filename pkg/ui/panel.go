package ui

import (
	"image/color"

	"github.com/macroblock/sdf/pkg/geom"
)

type (
	// Panel -
	Panel struct {
		KernelNode
		fgColor, bgColor color.RGBA
	}
)

// NewPanel -
func NewPanel() *Panel {
	o := &Panel{}
	o.self = o
	// o.border = geom.InitRect2i(1, 1, 1, 1)
	return o
}

// SetBounds -
func (o *Panel) SetBounds(rect geom.Rect2i) *Panel {
	o.Bounds = rect
	return o
}

// ClientRect -
func (o *Panel) ClientRect() geom.Rect2i {
	return shrinkToClientRect(o.Bounds, geom.InitRect2iAbs(1, 1, 1, 1))
}
