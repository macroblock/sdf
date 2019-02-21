package sdf

import (
	"github.com/macroblock/sdf/pkg/fonts"
	"github.com/veandco/go-sdl2/sdl"
)

// PixelFont -
type PixelFont struct {
	tex *Texture
	fonts.PixelFontSettings
	scale int32
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
	font := &PixelFont{PixelFontSettings: settings, scale: 1}
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

// SetScale -
func (o *PixelFont) SetScale(scale int) {
	if !Ok() {
		return
	}
	o.scale = int32(scale)
}

// Print -
func (o *PixelFont) Print(x0, y0 int, text string) {
	if !Ok() {
		return
	}
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
		err := o.printRune(int32(x), int32(y), r)
		setError(err)
		x += o.AdvanceX * int(o.scale)
	}
}

func (o *PixelFont) printRune(x, y int32, r rune) error {
	if r < o.MinRune || r > o.MaxRune {
		r = o.UnavailableRune
	}
	offsRune := int(r - o.MinRune)
	gw := int32(o.GlyphW)
	gh := int32(o.GlyphH)
	offsX := int32(offsRune % o.TilesX)
	offsY := int32(offsRune / o.TilesX)
	offsX *= gw
	offsY *= gh
	src := sdl.Rect{X: offsX, Y: offsY, W: gw, H: gh}
	dst := sdl.Rect{X: x, Y: y, W: gw * o.scale, H: gh * o.scale}

	// src = sdl.Rect{X: 0, Y: 9, W: 5, H: 9}
	// dst = sdl.Rect{X: x, Y: y, W: 5, H: 9}
	err := sdf.renderer.Copy(o.tex.sdltex, &src, &dst)
	return err
}
