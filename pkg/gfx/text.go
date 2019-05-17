package gfx

import (
	"image"

	"golang.org/x/image/math/fixed"
)

// DrawText -
func (o *Renderer) DrawText(x, y int, text string) {
	face := o.font
	if face == nil {
		face = o.defaultFont
	}
	x0 := x
	for _, r := range text {
		dr, mask, maskp, advance, ok := face.Glyph(fixed.Point26_6{}, r)
		if !ok {
			dr, mask, maskp, advance, ok = face.Glyph(fixed.Point26_6{}, '?')
		}

		switch r {
		case '\n':
			y += dr.Dy() //advanceY * int(o.scale)
			fallthrough
		case '\r':
			x = x0
			continue
		}

		src := image.Rect(0, 0, dr.Dx(), dr.Dy()).Add(maskp)
		dst := dr.Add(image.Pt(x, y))
		// fmt.Printf("src %q: %v\n", r, src)
		// fmt.Printf("dst %q: %v\n", r, dst)
		_ = dst
		_ = src
		tex, ok := mask.(*Texture)
		if !ok {
			panic("it is not a texture")
		}
		o.CopyRegion(tex, src, dst)
		x += advance.Round()
	}
}
