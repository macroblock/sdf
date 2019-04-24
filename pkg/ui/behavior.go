package ui

import (
	"github.com/macroblock/sdf/pkg/geom"
	"github.com/macroblock/sdf/pkg/gfx"
)

type (
	// IKernelNode2 -
	IKernelNode2 interface {
		// UIKernelNode() *KernelNode
		CoreBehavior() interface{}
		DrawBehavior() interface{}
		HandleEventBehavior() interface{}
	}

	coreBehavior struct {
		// Self    func(IKernelNode2) IKernelNode2
		// Parent  func(IKernelNode2) IKernelNode2
		Objects func(IKernelNode2) []IKernelNode2
	}

	drawBehavior struct {
		Renderer func(IKernelNode2) *gfx.Renderer
		RectNC   func(IKernelNode2) geom.Rect2i
		Rect     func(IKernelNode2) geom.Rect2i

		SizeNC func(IKernelNode2) geom.Point2i
		Size   func(IKernelNode2) geom.Point2i

		DrawNC func(IKernelNode2)
		Draw   func(IKernelNode2)
	}
)

// DrawScheme -
func DrawScheme(o IKernelNode2, zp geom.Point2i, clip geom.Rect2i) {
	core := o.CoreBehavior().(*coreBehavior)
	draw := o.DrawBehavior().(*drawBehavior)
	r := draw.Renderer(o)

	rect := draw.RectNC(o)
	// fmt.Println("----------------")
	// fmt.Println("zp  : ", zp)
	// fmt.Println("clip: ", clip)
	// fmt.Println("rect: ", rect)

	offset, clip := ClipRect(rect.Add(zp), clip)
	r.SetViewport(clip)
	r.SetOffset(offset)
	draw.DrawNC(o)

	// size := self.SizeNC()
	// r.SetColor(color.RGBA{255, 255, 255, 255})
	// r.DrawRect(0, 0, size.X, size.Y)

	rect = draw.Rect(o)
	offset, clip = ClipRect(rect.Add(zp), clip)
	r.SetViewport(clip)
	r.SetOffset(offset)
	draw.Draw(o)

	zp = zp.Add(rect.A)
	for _, child := range core.Objects(o) {
		DrawScheme(child, zp, clip)
	}
}
