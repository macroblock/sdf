package ui

import (
	"image"

	"github.com/macroblock/sdf/pkg/gfx"
)

type (
	// IKernelNode -
	IKernelNode interface {
		UIKernelNode() *KernelNode

		Renderer() *gfx.Renderer

		Objects() []IKernelNode

		RectNC() image.Rectangle
		Rect() image.Rectangle

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
