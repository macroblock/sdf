package ui

import (
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

// DrawScheme -
func (o *KernelNode) DrawScheme(offset geom.Point2i, clip geom.Rect2i) {
	self := o.self
	r := o.Renderer()
	rect := self.Rect().Add(offset)
	r.SetViewport(rect)
	self.DrawNC()

	rect = self.ClientRect().Add(rect.A)
	r.SetViewport(rect)
	self.Draw()

	offset = rect.A
	for _, child := range o.Children {
		child := child.UIKernelNode().self
		child.DrawScheme(offset, clip)
	}
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
	r.Print(11, 11, "Test")
	r.SetTextColor(color.RGBA{255, 255, 255, 255})
	r.Print(10, 10, "Test")
}

// Rect -
func (o *KernelNode) Rect() geom.Rect2i {
	return o.Bounds
}

// ClientRect -
func (o *KernelNode) ClientRect() geom.Rect2i {
	return o.Bounds.Sub(o.Bounds.A)
	// return geom.Rect2i{B2: o.Bounds.B2}
	// return geom.InitRect2i(0, 0, o.Bounds.W, o.Bounds.H)
	// return shrinkRect(o.Bounds, o.border)
}

// Renderer -
func (o *KernelNode) Renderer() *gfx.Renderer {
	if o.Root == nil {
		return nil
	}
	return o.Root.renderer
}
