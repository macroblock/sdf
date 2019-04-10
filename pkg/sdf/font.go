package sdf

import (
	"image/color"

	"github.com/macroblock/sdf/pkg/fonts"
	"github.com/macroblock/sdf/pkg/geom"
	"github.com/macroblock/sdf/pkg/gfx"
)

// PixelFont -
type PixelFont struct {
	tex *gfx.Texture
	fonts.PixelFontSettings
	color color.RGBA
	scale int
}

// // PixelFontSettings -
// type PixelFontSettings struct {
// 	FileName           string
// 	MinRune, MaxRune   rune
// 	UnavailableRune    rune
// 	TilesX, TilesY     int
// 	GlyphW, GlyphH     int
// 	AdvanceX, AdvanceY int
// 	// OriginX, OriginY   int
// 	BearingX, BearingY int
// }

// CreatePixelFont -
func CreatePixelFont(settings fonts.PixelFontSettings) *PixelFont {
	if !Ok() {
		return nil
	}
	font := &PixelFont{PixelFontSettings: settings, color: color.RGBA{255, 255, 255, 255}, scale: 1}
	tex := LoadTexture(font.FileName)
	if !Ok() {
		return nil
	}
	// surf, err := img.Load(font.FileName)
	// defer surf.Free()
	// if err != nil {
	// 	return nil, err
	// }
	// bound := surf.Bounds()

	// tex, err := sdf.renderer.CreateTextureFromSurface(surf)
	font.tex = tex
	// font.texW = bound.Max.X - bound.Min.X
	// font.texH = bound.Max.Y - bound.Min.Y
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

// NewFace -
func (o *PixelFont) NewFace(scale int) gfx.IFontFace {
	face := PixelFontFace{tex: o.tex, PixelFontSettings: o.PixelFontSettings}
	return &face
}

// SetScale -
func (o *PixelFont) SetScale(scale int) {
	if !Ok() {
		return
	}
	o.scale = scale
}

// SetColor -
func (o *PixelFont) SetColor(c color.RGBA) {
	o.color = c
}

// Print -
func (o *PixelFont) Print(x0, y0 int, text string) {
	if !Ok() {
		return
	}
	o.tex.SetColorMod(o.color)
	x0 -= o.BearingX
	y0 -= o.BearingY
	x, y := x0, y0
	for _, r := range text {
		switch r {
		case '\n':
			x = x0
			y += o.AdvanceY * int(o.scale)
			continue
		case '\r':
			x = x0
			continue
		}
		err := o.printRune(x, y, r)
		setError(err)
		x += o.AdvanceX * int(o.scale)
	}
}

func (o *PixelFont) printRune(x, y int, r rune) error {
	if r < o.MinRune || r > o.MaxRune {
		r = o.UnavailableRune
	}
	offsRune := int(r - o.MinRune)
	gw := o.GlyphW
	gh := o.GlyphH
	offsX := offsRune % o.TilesX
	offsY := offsRune / o.TilesX
	offsX *= gw
	offsY *= gh
	src := geom.InitRect2i(offsX, offsY, gw, gh)
	dst := geom.InitRect2i(x, y, gw*o.scale, gh*o.scale)

	// src = sdl.Rect{X: 0, Y: 9, W: 5, H: 9}
	// dst = sdl.Rect{X: x, Y: y, W: 5, H: 9}
	// err := sdf.renderer.Copy(o.tex.sdltex, &src, &dst)
	sdf.renderer.CopyRegion(o.tex, src, dst)
	return nil
}
