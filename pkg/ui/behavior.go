package ui

import (
	"image"
	"image/color"
)

// DrawScheme -
func DrawScheme(o IKernelNode, zp image.Point, clip image.Rectangle) {
	r := o.Renderer()

	rect := o.RectNC()
	// fmt.Println("----------------")
	// fmt.Println("zp  : ", zp)
	// fmt.Println("clip: ", clip)
	// fmt.Println("rect: ", rect)

	offset, clip := ClipRect(rect.Add(zp), clip)
	r.SetViewport(clip)
	r.SetOffset(offset)
	o.DrawNC()

	size := o.RectNC().Size()
	r.SetColor(color.RGBA{255, 255, 255, 255})
	r.DrawRect(0, 0, size.X, size.Y)

	rect = o.Rect()
	offset, clip = ClipRect(rect.Add(zp), clip)
	r.SetViewport(clip)
	r.SetOffset(offset)
	o.Draw()

	zp = zp.Add(rect.Min)
	for _, child := range o.Objects() {
		// child := child.UIKernelNode().self
		DrawScheme(child, zp, clip)
	}
}
