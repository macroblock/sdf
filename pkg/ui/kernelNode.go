package ui

import (
	"github.com/macroblock/sdf/pkg/geom"
	"github.com/macroblock/sdf/pkg/gfx"
	"github.com/macroblock/sdf/pkg/theme"
)

type (
	// KernelNode -
	KernelNode struct {
		self    IKernelNode
		Root    *UI
		Parent  IKernelNode
		objects []IKernelNode
		rect    geom.Rect2i
	}

	// HasObjects -
	HasObjects struct {
		objects []interface{}
	}
)

// UIKernelNode -
func (o *KernelNode) UIKernelNode() *KernelNode {
	return o
}

// Objects -
func (o *KernelNode) Objects() []IKernelNode {
	return o.objects
}

func setRootRecursive(src IKernelNode, root *UI) {
	kernel := src.UIKernelNode()
	for _, child := range kernel.objects {
		setRootRecursive(child, root)
	}
	kernel.Root = root
}

// AddObjects -
func (o *KernelNode) AddObjects(objects ...IKernelNode) {
	for _, obj := range objects {
		obj.UIKernelNode().Parent = o.self
		setRootRecursive(obj, o.Root)
		o.objects = append(o.objects, obj)
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

// DrawNC -
func (o *KernelNode) DrawNC() {
	// fmt.Println("draw nc")
	// self := o.self

	// r := o.Renderer()
	// r.SetColor(theme.PlaceHolder.Color())
	// r.Clear()
}

// Draw -
func (o *KernelNode) Draw() {
	// self := o.self

	r := o.Renderer()
	r.SetColor(theme.Background.Color())
	r.Clear()
	r.SetTextColor(theme.Text.Color())
	// r.DrawText2(0, 16, " !jTyestxyw")
	r.DrawText(0, 16, " !jTyestxyw")
}

// RectNC -
func (o *KernelNode) RectNC() geom.Rect2i {
	return o.rect
}

// Rect -
func (o *KernelNode) Rect() geom.Rect2i {
	return o.rect //.Sub(o.Bounds.A)
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

// Renderer -
func (o *KernelNode) Renderer() *gfx.Renderer {
	if o.Root == nil {
		return nil
	}
	return o.Root.renderer
}

// HandleEventScheme -
// func (o *KernelNode) HandleEventScheme(zp geom.Point2i, ev event.IEvent) bool {
// 	self := o.self
// 	r := o.Renderer()
// 	rect := self.RectNC().Add(zp)
// 	ok := true
// 	switch ev := ev.(type) {
// 	case *event.MouseClick:
// 		ok = rect.ConsistsInt(ev.X, ev.Y)
// 	case *event.MouseMotion:
// 		ok = rect.ConsistsInt(ev.X, ev.Y)
// 	}
// 	if !ok {
// 		return false
// 	}
// 	zp = zp.Add(rect.A)
// 	for _, child := range o.Children {
// 		child := child.UIKernelNode().self
// 		ok := child.HandleEventScheme(zp, ev)
// 		if ok {
// 			return true
// 		}
// 	}

// 	return true
// }
