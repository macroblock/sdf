package ui

import "github.com/macroblock/sdf/pkg/geom"

type (
	// IKernelNode -
	IKernelNode interface {
		UIKernelNode() *KernelNode

		Rect() geom.Rect2i
		ClientRect() geom.Rect2i

		Draw()
		DrawNC()
		DrawScheme(geom.Point2i, geom.Rect2i)
	}
)
