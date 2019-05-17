package ui

import (
	"image"
)

type (
	// Panel -
	Panel struct {
		KernelNode
	}
)

// InitPanel -
func InitPanel() Panel {
	o := Panel{}
	// o.border = geom.InitRect2i(1, 1, 1, 1)
	return o
}

// NewPanel -
func NewPanel() *Panel {
	o := InitPanel()
	o.self = &o
	// o.border = geom.InitRect2i(1, 1, 1, 1)
	return &o
}

// SetPos -
func (o *Panel) SetPos(pos image.Point) {
	pos = pos.Sub(o.rect.Min)
	o.rect = o.rect.Add(pos)
	// return o
}

// SetBounds -
func (o *Panel) SetBounds(rect image.Rectangle) {
	o.rect = rect
	// return o
}

// Rect -
func (o *Panel) Rect() image.Rectangle {
	// return shrink(o.Bounds, geom.InitRect2iAbs(1, 1, 1, 1))
	return o.rect
}
