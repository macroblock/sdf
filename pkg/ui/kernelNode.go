package ui

import (
	"image"

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
		rect    image.Rectangle
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
func ClipRect(rect, clip image.Rectangle) (image.Point, image.Rectangle) {
	switch {
	case clip.Min.X < rect.Min.X:
		clip.Min.X = rect.Min.X
	case clip.Min.X > rect.Max.X:
		clip.Min.X = rect.Max.X
	}
	switch {
	case clip.Min.Y < rect.Min.Y:
		clip.Min.Y = rect.Min.Y
	case clip.Min.Y > rect.Max.Y:
		clip.Min.Y = rect.Max.Y
	}
	switch {
	case clip.Max.X > rect.Max.X:
		clip.Max.X = rect.Max.X
	case clip.Max.X < rect.Min.X:
		clip.Max.X = rect.Min.X
	}
	switch {
	case clip.Max.Y > rect.Max.Y:
		clip.Max.Y = rect.Max.Y
	case clip.Max.Y < rect.Min.Y:
		clip.Max.Y = rect.Min.Y
	}
	offset := image.Pt(clip.Min.X-rect.Min.X, clip.Min.Y-rect.Min.Y)
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
func (o *KernelNode) RectNC() image.Rectangle {
	return o.rect
}

// Rect -
func (o *KernelNode) Rect() image.Rectangle {
	return o.rect //.Sub(o.Bounds.A)
	// return geom.Rect2i{B2: o.Bounds.B2}
	// return geom.InitRect2i(0, 0, o.Bounds.W, o.Bounds.H)
	// return shrinkRect(o.Bounds, o.border)
}

// SizeNC -
func (o *KernelNode) SizeNC() image.Point {
	rect := o.self.RectNC()
	return image.Pt(rect.Dx(), rect.Dy())
}

// Size -
func (o *KernelNode) Size() image.Point {
	rect := o.self.Rect()
	return image.Pt(rect.Dx(), rect.Dy())
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
