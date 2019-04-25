package ui

import (
	"github.com/macroblock/sdf/pkg/geom"
	"github.com/macroblock/sdf/pkg/gfx"
)

type (
	// IKernelNode -
	IKernelNode interface {
		UIKernelNode() *KernelNode

		Renderer() *gfx.Renderer

		Objects() []IKernelNode

		RectNC() geom.Rect2i
		Rect() geom.Rect2i

		// SizeNC() geom.Point2i
		// Size() geom.Point2i
		// BoundsNC() geom.Rect2i
		// Bounds() geom.Rect2i

		DrawNC()
		Draw()

		// HandleEventScheme(geom.Point2i, event.IEvent) bool
		// HandleEvent(event.IEvent) bool
		// HandleEventNC(event.IEvent) bool
	}
)
