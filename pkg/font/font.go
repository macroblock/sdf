package font

import (
	"image/color"

	"github.com/macroblock/sdf/pkg/gfx"
)

// PixelFont -
type PixelFont struct {
	PixelFontSettings
	tex   *gfx.Texture
	color color.RGBA
	scale int
}

// PixelFontSettings -
type PixelFontSettings struct {
	FileName           string
	MinRune, MaxRune   rune
	UnavailableRune    rune
	TilesX, TilesY     int
	GlyphW, GlyphH     int
	AdvanceX, AdvanceY int
	BearingX, BearingY int
}

// CreatePixelFont -
func CreatePixelFont(settings PixelFontSettings, tex *gfx.Texture) *PixelFont {
	// if !Ok() {
	// 	return nil
	// }
	font := &PixelFont{PixelFontSettings: settings, color: color.RGBA{255, 255, 255, 255}, scale: 1}

	font.tex = tex
	if font.TilesX <= 0 {
		font.TilesX = font.tex.W / font.GlyphW
	} else {
		font.GlyphW = font.tex.W / font.TilesX
	}
	if font.TilesY <= 0 {
		font.TilesY = font.tex.H / font.GlyphH
	} else {
		font.GlyphH = font.tex.H / font.TilesY
	}
	return font
}

// // NewFace -
// func (o *PixelFont) NewFace(scale int) gfx.IFontFace {
// 	face := PixelFontFace{tex: o.tex, PixelFontSettings: o.PixelFontSettings}
// 	return &face
// }
