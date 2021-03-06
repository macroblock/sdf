package unifont

import (
	"image"
	"image/color"
	"unicode"

	"github.com/golang/freetype/truetype"
	"github.com/macroblock/sdf/pkg/gfx"
	"github.com/macroblock/sdf/pkg/misc"
	"github.com/veandco/go-sdl2/sdl"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

type (
	// HWFace -
	HWFace struct {
		font   *truetype.Font
		tex    *gfx.Texture
		coords map[rune]*RuneCoords
	}
	// RuneCoords -
	RuneCoords struct {
		Rect    image.Rectangle
		ZpOffs  image.Point
		Advance fixed.Int26_6
	}
)

// Close -
func (o *HWFace) Close() error {
	// TODO: fix me
	return nil
}

// Glyph -
func (o *HWFace) Glyph(dot fixed.Point26_6, r rune) (
	dr image.Rectangle, mask image.Image, maskp image.Point, advance fixed.Int26_6, ok bool) {
	coords, ok := o.coords[r]
	if !ok {
		return image.Rectangle{}, nil, image.Point{}, 0, false
	}
	dr = coords.Rect
	mask = o.tex
	maskp = coords.ZpOffs
	advance = coords.Advance
	ok = true
	return
}

// GlyphBounds -
func (o *HWFace) GlyphBounds(r rune) (bounds fixed.Rectangle26_6, advance fixed.Int26_6, ok bool) {
	coords, ok := o.coords[r]
	if !ok {
		return fixed.Rectangle26_6{}, 0, false
	}
	bounds = fixed.R(coords.Rect.Min.X, coords.Rect.Min.Y, coords.Rect.Max.X, coords.Rect.Max.Y) // TODO: is it correct ?
	advance = coords.Advance
	ok = true
	return
}

// GlyphAdvance -
func (o *HWFace) GlyphAdvance(r rune) (advance fixed.Int26_6, ok bool) {
	coords, ok := o.coords[r]
	if !ok {
		return 0, false
	}
	return coords.Advance, true
}

// Kern -
func (o *HWFace) Kern(r0, r1 rune) fixed.Int26_6 {
	// TODO: fix it
	coords, ok := o.coords[r0]
	if !ok {
		return 0
	}
	return coords.Advance
}

// Metrics -
func (o *HWFace) Metrics() font.Metrics {
	// TODO: fix it
	panic("not implemented yet")
}

// NewHWFace -
func NewHWFace(renderer *gfx.Renderer, font *truetype.Font, opts *truetype.Options) (*HWFace, error) {

	face := HWFace{font: font, coords: map[rune]*RuneCoords{}}
	f := truetype.NewFace(font, opts)
	defer f.Close()

	size := image.Point{}
	sizeV := image.Point{}

	for r := rune(0); r <= unicode.MaxRune; r++ {
		if 0xe000 <= r && r <= 0xf8ff ||
			0xf0000 <= r && r <= 0xffffd ||
			0x100000 <= r && r <= 0x10fffd ||
			font.Index(r) == 0 {
			continue
		}

		dr, _, _, _, ok := f.Glyph(fixed.Point26_6{}, r)
		if !ok {
			continue
		}
		coords := &RuneCoords{}
		face.coords[r] = coords

		size.X += dr.Dx()
		size.Y = misc.MaxInt(size.Y, dr.Dy())
		sizeV.X = misc.MaxInt(sizeV.X, dr.Dx())
		sizeV.Y += dr.Dy()
	}

	delta := image.Pt(1, 0)
	if size.X*size.Y > sizeV.X*sizeV.Y {
		size = sizeV
		delta = image.Pt(0, 1)
	}

	img := image.NewRGBA(image.Rect(0, 0, size.X, size.Y))

	dstOffs := image.Point{}
	for r, coords := range face.coords {
		dr, mask, maskpoint, advance, ok := f.Glyph(fixed.Point26_6{}, r)
		if !ok {
			panic("unreachable")
		}
		coords.Rect = dr
		coords.ZpOffs = dstOffs
		coords.Advance = advance

		srcOffs := maskpoint
		blitRune(mask, img, srcOffs, dstOffs, coords.Rect)

		dstOffs.X += coords.Rect.Dx() * delta.X
		dstOffs.Y += coords.Rect.Dy() * delta.Y
	}

	tex, err := renderer.ImageToTexture(img)
	if err != nil {
		return nil, err
	}
	tex.SetBlendMode(sdl.BLENDMODE_BLEND)

	_ = tex
	face.tex = tex

	return &face, nil
}

func blitRune(mask image.Image, dst *image.RGBA, srcOffs, dstOffs image.Point, rect image.Rectangle) {
	// fmt.Println("---------------------------------------")
	for y := 0; y < rect.Dy(); y++ {
		// s := ""
		for x := 0; x < rect.Dx(); x++ {
			c := mask.At(x+srcOffs.X, y+srcOffs.Y)
			r, g, b, a := c.RGBA()
			_, _, _, _ = r, g, b, a
			newColor := color.RGBA{255, 255, 255, uint8(r)}
			// setPixel(surf, x+dstOffs.X, y+dstOffs.Y, newColor)
			dst.Set(x+dstOffs.X, y+dstOffs.Y, newColor)
			// r, g, b, a = newColor.RGBA()
			// s += fmt.Sprintf("#%v %v %v %v - %v", r, g, b, a, newColor)
		}
		// _ = s
		// fmt.Println(s)
	}
	// fmt.Println(sdl.GetPixelFormatName(uint(surf.Format.Format)))
}
