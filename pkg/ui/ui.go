package ui

import (
	"image/color"

	"github.com/macroblock/sdf/pkg/gfx"
)

type (
	// UI -
	UI struct {
		Panel
		renderer *gfx.Renderer
	}
)

// NewUI -
func NewUI(renderer *gfx.Renderer) *UI {
	o := &UI{}
	o.self = o
	o.KernelNode.Root = o
	o.renderer = renderer
	// o.border = geom.InitRect2i(1, 1, 1, 1)
	return o
}

// Draw -
func (o *UI) Draw() {
	self := o.self
	r := o.Renderer()
	rect := self.Rect()
	r.ClearAll()
	r.SetViewport(rect)
	r.SetColor(color.RGBA{50, 150, 100, 255})
	r.Clear()
	for _, child := range o.Children {
		child := child.UIKernelNode().self
		clip := o.Bounds
		child.DrawScheme(rect.A, clip)
	}
	r.ResetViewport()
}
