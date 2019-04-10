package sdf

import (
	"github.com/macroblock/sdf/pkg/fonts"
	"github.com/macroblock/sdf/pkg/geom"
	"github.com/macroblock/sdf/pkg/gfx"
	"golang.org/x/image/math/fixed"
)

type (
	// PixelFontFace -
	PixelFontFace struct {
		tex *gfx.Texture
		fonts.PixelFontSettings
	}
)

// Glyph -
func (o *PixelFontFace) Glyph(dot fixed.Point26_6, r rune) (bounds geom.Rect2i, mask *gfx.Texture, bearing geom.Point2i, advance int, ok bool) {
	ok = true
	if r < o.MinRune || r > o.MaxRune {
		r = o.UnavailableRune
		ok = false
	}
	offsRune := int(r - o.MinRune)
	gw := o.GlyphW
	gh := o.GlyphH
	offsX := offsRune % o.TilesX
	offsY := offsRune / o.TilesX
	offsX *= gw
	offsY *= gh

	bounds = geom.InitRect2i(offsX, offsY, gw, gh)
	mask = o.tex
	bearing = geom.InitPoint2i(o.BearingX, o.BearingY)
	advance = o.AdvanceX
	// dst := geom.InitRect2i(x, y, gw*o.scale, gh*o.scale)
	// sdf.renderer.CopyRegion(o.tex, src, dst)
	return
}
