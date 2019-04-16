package ui

import "github.com/macroblock/sdf/pkg/geom"

type (
	// IKernelNode -
	IKernelNode interface {
		UIKernelNode() *KernelNode

		RectNC() geom.Rect2i
		Rect() geom.Rect2i

		SizeNC() geom.Point2i
		Size() geom.Point2i
		// BoundsNC() geom.Rect2i
		// Bounds() geom.Rect2i

		Draw()
		DrawNC()
		DrawScheme(geom.Point2i, geom.Rect2i)
	}
)
