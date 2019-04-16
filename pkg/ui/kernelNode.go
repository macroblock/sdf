package ui

import (
	"fmt"
	"image/color"

	"github.com/macroblock/sdf/pkg/geom"
	"github.com/macroblock/sdf/pkg/gfx"
)

type (
	// KernelNode -
	KernelNode struct {
		self     IKernelNode
		Root     *UI
		Parent   IKernelNode
		Children []IKernelNode
		Bounds   geom.Rect2i
	}
)

// UIKernelNode -
func (o *KernelNode) UIKernelNode() *KernelNode {
	return o
}

func setRootRecursive(src IKernelNode, root *UI) {
	kernel := src.UIKernelNode()
	for _, child := range kernel.Children {
		setRootRecursive(child, root)
	}
	kernel.Root = root
}

// AddChildren -
func (o *KernelNode) AddChildren(children ...IKernelNode) {
	for _, child := range children {
		child.UIKernelNode().Parent = o.self
		setRootRecursive(child, o.Root)
		o.Children = append(o.Children, child)
	}
}

// ClipRect -
func ClipRect(rect, clip geom.Rect2i) (geom.Point2i, geom.Rect2i) {
	switch {
	case clip.A.X < rect.A.X:
		clip.A.X = rect.A.X
	case clip.A.X > rect.B.X:
		clip.A.X = rect.B.X
	}
	switch {
	case clip.A.Y < rect.A.Y:
		clip.A.Y = rect.A.Y
	case clip.A.Y > rect.B.Y:
		clip.A.Y = rect.B.Y
	}
	switch {
	case clip.B.X > rect.B.X:
		clip.B.X = rect.B.X
	case clip.B.X < rect.A.X:
		clip.B.X = rect.A.X
	}
	switch {
	case clip.B.Y > rect.B.Y:
		clip.B.Y = rect.B.Y
	case clip.B.Y < rect.A.Y:
		clip.B.Y = rect.A.Y
	}
	offset := geom.InitPoint2i(clip.A.X-rect.A.X, clip.A.Y-rect.A.Y)
	return offset, clip
}

// DrawScheme -
func (o *KernelNode) DrawScheme(zp geom.Point2i, clip geom.Rect2i) {
	self := o.self
	r := o.Renderer()

	rect := self.RectNC()
	fmt.Println("----------------")
	fmt.Println("zp  : ", zp)
	fmt.Println("clip: ", clip)
	fmt.Println("rect: ", rect)

	offset, clip := ClipRect(rect.Add(zp), clip)
	r.SetViewport(clip)
	r.SetOffset(offset)

	self.DrawNC()
	size := self.SizeNC()
	r.SetColor(color.RGBA{255, 255, 255, 255})
	r.DrawRect(0, 0, size.X, size.Y)

	rect = self.Rect()
	offset, clip = ClipRect(rect.Add(zp), clip)
	r.SetViewport(clip)
	r.SetOffset(offset)

	self.Draw()

	zp = zp.Add(rect.A)
	// clip = clip.Add(offset)
	for _, child := range o.Children {
		child := child.UIKernelNode().self
		child.DrawScheme(zp, clip)
	}

	// self := o.self
	// r := o.Renderer()
	// rect := self.Rect().Add(zp)
	// r.SetViewport(rect)
	// self.DrawNC()

	// rect = self.ClientRect().Add(rect.A)
	// r.SetViewport(rect)
	// self.Draw()

	// zp = rect.A
	// for _, child := range o.Children {
	// 	child := child.UIKernelNode().self
	// 	child.DrawScheme(zp, clip)
	// }
}

// DrawNC -
func (o *KernelNode) DrawNC() {
	// fmt.Println("draw nc")
	// self := o.self
	c1 := color.RGBA{200, 100, 100, 255}
	// c2 := color.RGBA{150, 100, 200, 255}
	r := o.Renderer()
	r.SetColor(c1)
	r.Clear()
}

// Draw -
func (o *KernelNode) Draw() {
	// self := o.self
	// c1 := color.RGBA{200, 150, 100, 255}
	// fmt.Println("draw")
	c2 := color.RGBA{100, 100, 200, 255}
	r := o.Renderer()
	r.SetColor(c2)
	r.Clear()
	r.SetTextColor(color.RGBA{0, 0, 0, 100})
	// r.Print(10, 9, "Test")
	// r.Print(10, 11, "Test")
	// r.Print(9, 10, "Test")
	// r.Print(11, 10, "Test")
	// r.Print(11, 11, "Test")
	r.DrawText(11, 11, "Test")
	r.SetTextColor(color.RGBA{255, 255, 255, 255})
	r.DrawText(10, 10, "Test")
}

// RectNC -
func (o *KernelNode) RectNC() geom.Rect2i {
	return o.Bounds
}

// Rect -
func (o *KernelNode) Rect() geom.Rect2i {
	return o.Bounds //.Sub(o.Bounds.A)
	// return geom.Rect2i{B2: o.Bounds.B2}
	// return geom.InitRect2i(0, 0, o.Bounds.W, o.Bounds.H)
	// return shrinkRect(o.Bounds, o.border)
}

// SizeNC -
func (o *KernelNode) SizeNC() geom.Point2i {
	rect := o.self.RectNC()
	return geom.InitPoint2i(rect.W(), rect.H())
}

// Size -
func (o *KernelNode) Size() geom.Point2i {
	rect := o.self.Rect()
	return geom.InitPoint2i(rect.W(), rect.H())
}

// // Bounds -
// func (o *KernelNode) Bounds() geom.Rect2i {
// 	return geom.InitRect2iAbs(1,1,, B: o.Bounds.B}
// }

// Renderer -
func (o *KernelNode) Renderer() *gfx.Renderer {
	if o.Root == nil {
		return nil
	}
	return o.Root.renderer
}
