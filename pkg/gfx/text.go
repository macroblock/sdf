package gfx

import (
	"fmt"
	"image"
	"image/color"

	"github.com/macroblock/sdf/pkg/geom"
	"github.com/veandco/go-sdl2/sdl"
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

// // DrawText -
// func (o *Renderer) DrawText(x, y int, text string) {
// 	font := o.font
// 	if font == nil {
// 		font = defaultFont
// 	}
// 	x0 := x
// 	for _, r := range text {
// 		bounds, tex, bearing, advance, ok := font.Glyph(fixed.Point26_6{}, r)
// 		_ = ok
// 		switch r {
// 		case '\n':
// 			x = x0
// 			y += bounds.H() //advanceY * int(o.scale)
// 			continue
// 		case '\r':
// 			x = x0
// 			continue
// 		}
// 		dst := geom.InitRect2i(x-bearing.X, y-bearing.Y, bounds.W(), bounds.H())

// 		// src = sdl.Rect{X: 0, Y: 9, W: 5, H: 9}
// 		// dst = sdl.Rect{X: x, Y: y, W: 5, H: 9}
// 		// err := sdf.renderer.Copy(o.tex.sdltex, &src, &dst)

// 		tex.SetColorMod(o.textColor)
// 		o.CopyRegion(tex, bounds, dst)

// 		x += advance
// 	}
// }

func image2Surface(img image.Image) (*sdl.Surface, error) {
	rgba := image.NewRGBA(img.Bounds())
	// rgba := image.NewGray(img.Bounds())
	size := img.Bounds().Size()
	s, err := sdl.CreateRGBSurface(0, int32(size.X), int32(size.Y), 32, 0x000000ff, 0x0000ff00, 0x00ff0000, 0xff000000)
	if err != nil {
		return s, err
	}
	rgba.Pix = s.Pixels()

	for y := 0; y < size.Y; y++ {
		s := ""
		for x := 0; x < size.X; x++ {
			c := img.At(x, y)
			r, g, b, a := c.RGBA()
			_, _, _, _ = r, g, b, a
			// newColor := color.Gray{uint8(r >> 8)}
			newColor := color.RGBA{255, 255, 255, uint8(r)}
			s += fmt.Sprintf("%v", newColor)
			rgba.Set(x, y, newColor)
		}
		_ = s
		// fmt.Println(s)
	}
	// for y := 0; y < size.Y; y++ {
	// 	for x := 0; x < size.X; x++ {
	// 		c := img.At(x, y)
	// 		rgba.Set(x, y, c)
	// 	}
	// }

	// tmp := image.NewRGBA(img.Bounds())
	// tmp.Pix = s.Pixels()
	// fmt.Println("Pixels ", s.Pixels()[0:10], " VS ", rgba.Pix[0:10])
	// saveTemp(tmp)      // save correct
	// s.SaveBMP("s.bmp") // Byte order is incorrect under mac os

	return s, err
}

// // DrawText2 -
// func (o *Renderer) DrawText2(x, y int, text string) {
// 	face := theme.AcquireFontFace()
// 	x0 := x
// 	for _, r := range text {
// 		// bounds, tex, bearing, advance, ok := font.Glyph(fixed.Point26_6{}, r)
// 		bounds, img, maskpoint, advance, ok := face.Glyph(fixed.Point26_6{}, r)
// 		tex, err := o.ImageToTexture(img)
// 		if err != nil {
// 			panic(fmt.Sprint("imagToSurface: ", err))
// 		}

// 		_ = ok
// 		switch r {
// 		case '\n':
// 			y += bounds.Dy() //advanceY * int(o.scale)
// 			fallthrough
// 		case '\r':
// 			x = x0
// 			continue
// 		}

// 		zp := geom.InitPoint2i(maskpoint.X, maskpoint.Y)
// 		src := geom.InitRect2iAbs(0, 0, bounds.Dx(), bounds.Dy()).Add(zp)
// 		offset := geom.InitPoint2i(bounds.Min.X, bounds.Min.Y)
// 		dst := geom.InitRect2i(x, y, bounds.Dx(), bounds.Dy()).Add(offset)
// 		_ = dst
// 		_ = src
// 		tex.SetColorMod(o.textColor)
// 		o.CopyRegion(tex, src, dst)
// 		// fmt.Println("maskpoint: ", maskpoint, " bounds: ", bounds)
// 		// fmt.Println("src: ", src, " dst: ", dst)
// 		x += advance.Round()
// 	}
// }

// DrawText -
func (o *Renderer) DrawText(x, y int, text string) {
	face := o.font
	if face == nil {
		face = o.defaultFont
	}
	x0 := x
	for _, r := range text {
		// bounds, tex, bearing, advance, ok := font.Glyph(fixed.Point26_6{}, r)
		tex, coords, ok := face.Glyph(r)
		if !ok {
			tex, coords, ok = face.Glyph('?')
		}

		switch r {
		case '\n':
			y += coords.Rect.H() //advanceY * int(o.scale)
			fallthrough
		case '\r':
			x = x0
			continue
		}

		src := coords.Rect
		// fmt.Printf("src %q: %v\n", r, src)
		offset := coords.ZpOffs
		dst := geom.InitRect2iAbs(0, 0, src.W(), src.H()).AddInt(x, y).Add(offset)
		// fmt.Printf("dst %q: %v\n", r, dst)
		_ = dst
		_ = src
		// tex.SetColorMod(o.textColor)
		o.CopyRegion(tex, src, dst)
		x += coords.Advance
	}
}
