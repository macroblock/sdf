package ui

import (
	"github.com/macroblock/sdf/pkg/geom"
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
func (o *Panel) SetPos(pos geom.Point2i) {
	o.rect = o.rect.SetPos(pos)
	// return o
}

// SetBounds -
func (o *Panel) SetBounds(rect geom.Rect2i) {
	o.rect = rect
	// return o
}

// Rect -
func (o *Panel) Rect() geom.Rect2i {
	// return shrink(o.Bounds, geom.InitRect2iAbs(1, 1, 1, 1))
	return o.rect
}
