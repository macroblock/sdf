package font

import (
	"fmt"
	"unicode"

	"github.com/macroblock/sdf/pkg/geom"
	"github.com/macroblock/sdf/pkg/gfx"
	"github.com/macroblock/sdf/pkg/misc"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

type (
	// PixelFontFace -
	PixelFontFace struct {
		tex *gfx.Texture
		PixelFontSettings
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

func glyphMetrics(face font.Face, r rune) {
	dr, mask, mpt, adv, ok := face.Glyph(fixed.Point26_6{}, r)
	if ok
}

func parseRanges(face font.Face) {
	maxAdvance := fixed.Int26_6(-1)
	// if dr, mask, mpt, adv, ok := face.Glyph(fixed.Point26_6{}, r); ok {
	for r := rune(0); r <= unicode.MaxRune; r++ {
		if 0xe000 <= r && r <= 0xf8ff ||
			0xf0000 <= r && r <= 0xffffd ||
			0x100000 <= r && r <= 0x10fffd {
			continue
		}
		// if dr, mask, mpt, adv, ok := face.Glyph(fixed.Point26_6{}, r); ok {
		// 	_ = mask
		// 	_ = dr
		// 	_ = mpt
		// 	_ = adv
		// 	// fmt.Printf("valid skip %U %v %v %v\n", r, dr, mpt, adv)
		// 	// continue
		// }
		if ttf.Index(r) == 0 {
			// fmt.Printf("skiped %U\n", r)
			continue
		}
		dr, _, maskp, adv, ok := face.Glyph(fixed.Point26_6{}, r)
		if !ok {
			return nil, -1, -1, fixed.Int26_6(0), fmt.Errorf("could not load glyph %q %U", r, r)
		}
		maxAdvance = fixed.Int26_6(misc.MaxInt(int(maxAdvance), int(adv)))
		volume += dr.Dx() * dr.Dy()
		maskData := tMask{
			r:         r,
			destRect:  dr,
			maskPoint: maskp,
			advance:   adv,
		}
		slice = append(slice, maskData)
	}
}

// ParseFontFace -
func ParseFontFace(face font.Face) {

}
