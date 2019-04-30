package gfx

import (
	"fmt"
	"image"

	"github.com/macroblock/sdf/pkg/geom"
	"github.com/macroblock/sdf/pkg/theme"
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

// DrawText -
func (o *Renderer) DrawText(x, y int, text string) {
	font := o.font
	if font == nil {
		font = defaultFont
	}
	x0 := x
	for _, r := range text {
		bounds, tex, bearing, advance, ok := font.Glyph(fixed.Point26_6{}, r)
		_ = ok
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

		tex.SetColorMod(o.textColor)
		o.CopyRegion(tex, bounds, dst)

		x += advance
	}
}

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
			s += fmt.Sprintf("%v", c)
			rgba.Set(x, y, c)
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
	s.SaveBMP("s.bmp") // Byte order is incorrect under mac os

	return s, err
}

// DrawText2 -
func (o *Renderer) DrawText2(x, y int, text string) {
	// font := o.font
	// if font == nil {
	// 	font = defaultFont
	// }
	// info, err := o.SDLRenderer().GetInfo()
	// if err != nil {
	// 	panic(fmt.Sprint("sdl renderer get info: ", err))
	// }
	// info.RendererInfoData
	face := theme.AcquireFontFace()
	x0 := x
	for _, r := range text {
		// bounds, tex, bearing, advance, ok := font.Glyph(fixed.Point26_6{}, r)
		bounds, img, bearing, advance, ok := face.Glyph(fixed.Point26_6{}, r)
		surf, err := image2Surface(img)
		if err != nil {
			panic(fmt.Sprint("sdl image2surface: ", err))
		}
		sdltex, err := o.SDLRenderer().CreateTextureFromSurface(surf)
		if err != nil {
			panic(fmt.Sprint("sdl texture from surface: ", err))
		}
		tex := &Texture{
			W:      surf.Bounds().Dx(),
			H:      surf.Bounds().Dy(),
			sdltex: sdltex,
		}

		// err = o.SDLRenderer().Copy(sdltex, nil, nil)
		// if err != nil {
		// 	panic(fmt.Sprint("sdl texture copy: ", err))
		// }

		_ = ok
		switch r {
		case '\n':
			x = x0
			y += bounds.Dy() //advanceY * int(o.scale)
			continue
		case '\r':
			x = x0
			continue
		}
		dst := geom.InitRect2i(x-bearing.X, y-bearing.Y, bounds.Dx(), bounds.Dy())
		// _ = bearing
		// dst := geom.InitRect2i(x+10, y+10, 8, 16)

		// src = sdl.Rect{X: 0, Y: 9, W: 5, H: 9}
		// dst = sdl.Rect{X: x, Y: y, W: 5, H: 9}
		// err := sdf.renderer.Copy(o.tex.sdltex, &src, &dst)

		box := geom.InitRect2iAbs(bounds.Min.X, bounds.Min.Y, bounds.Max.X, bounds.Max.Y)
		// box := geom.InitRect2iAbs(0, 16, 8, 16)
		tex.SetColorMod(o.textColor)
		o.SDLRenderer().SetDrawBlendMode(sdl.BLENDMODE_NONE)
		o.CopyRegion(tex, box, dst)

		x += advance.Round()
	}
}
