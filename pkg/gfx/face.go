package gfx

import (
	"github.com/macroblock/sdf/pkg/geom"
	"golang.org/x/image/math/fixed"
)

// IFontFace -
type IFontFace interface {
	// io.Closer

	// Glyph returns the draw.DrawMask parameters (dr, mask, maskp) to draw r's
	// glyph at the sub-pixel destination location dot, and that glyph's
	// advance width.
	//
	// It returns !ok if the face does not contain a glyph for r.
	//
	// The contents of the mask image returned by one Glyph call may change
	// after the next Glyph call. Callers that want to cache the mask must make
	// a copy.
	Glyph(dot fixed.Point26_6, r rune) (
		bounds geom.Rect2i, tex *Texture, bearing geom.Point2i, advance int, ok bool)

	// GlyphBounds returns the bounding box of r's glyph, drawn at a dot equal
	// to the origin, and that glyph's advance width.
	//
	// It returns !ok if the face does not contain a glyph for r.
	//
	// The glyph's ascent and descent equal -bounds.Min.Y and +bounds.Max.Y. A
	// visual depiction of what these metrics are is at
	// https://developer.apple.com/library/mac/documentation/TextFonts/Conceptual/CocoaTextArchitecture/Art/glyph_metrics_2x.png
	// GlyphBounds(r rune) (bounds fixed.Rectangle26_6, advance fixed.Int26_6, ok bool)

	// GlyphAdvance returns the advance width of r's glyph.
	//
	// It returns !ok if the face does not contain a glyph for r.
	// GlyphAdvance(r rune) (advance fixed.Int26_6, ok bool)

	// Kern returns the horizontal adjustment for the kerning pair (r0, r1). A
	// positive kern means to move the glyphs further apart.
	// Kern(r0, r1 rune) fixed.Int26_6

	// Metrics returns the metrics for this Face.
	// Metrics() Metrics
}

// DrawText -
func (o *Renderer) DrawText(x, y int, text string) {
	x0, y0 := x, y
	for _, r := range text {
		bounds, tex, bearing, advance, ok := o.font.Glyph(fixed.Point26_6{}, r)
		switch r {
		case '\n':
			x = x0
			y += bounds.H() //advanceY * int(o.scale)
			continue
		case '\r':
			x = x0
			continue
		}
		dst := geom.InitRect2i(x-bearing.X, y-bearing.Y, bounds.W(), bounds.H())

		// src = sdl.Rect{X: 0, Y: 9, W: 5, H: 9}
		// dst = sdl.Rect{X: x, Y: y, W: 5, H: 9}
		// err := sdf.renderer.Copy(o.tex.sdltex, &src, &dst)
		o.CopyRegion(tex, bounds, dst)

		x += advance
	}
}
