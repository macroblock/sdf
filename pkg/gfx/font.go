package gfx

import (
	"fmt"
	"image"
	"image/color"
	"unicode"

	"github.com/golang/freetype/truetype"
	"github.com/macroblock/sdf/pkg/geom"
	"github.com/macroblock/sdf/pkg/misc"
	"github.com/veandco/go-sdl2/sdl"
	"golang.org/x/image/math/fixed"
)

type (
	// HWFace -
	HWFace struct {
		font   *truetype.Font
		tex    *Texture
		coords map[rune]*RuneCoords
	}
	// RuneCoords -
	RuneCoords struct {
		Rect    geom.Rect2i
		ZpOffs  geom.Point2i
		Advance int
	}
)

// Glyph -
func (o *HWFace) Glyph(r rune) (*Texture, *RuneCoords, bool) {
	coords, ok := o.coords[r]
	return o.tex, coords, ok
}

// NewHWFace -
func (o *Renderer) NewHWFace(font *truetype.Font, opts *truetype.Options) (*HWFace, error) {

	face := HWFace{font: font, coords: map[rune]*RuneCoords{}}
	f := truetype.NewFace(font, opts)

	size := geom.Point2i{}
	sizeV := geom.Point2i{}

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
	delta := geom.InitPoint2i(1, 0)
	if size.X*size.Y > sizeV.X*sizeV.Y {
		size = sizeV
		delta = geom.InitPoint2i(0, 1)
		fmt.Println("vertical texture: ", size)
	} else {
		fmt.Println("horizontal texture: ", size)
	}

	surf, err := sdl.CreateRGBSurface(0, int32(size.X), int32(size.Y), 32, 0x00000ff, 0x0000ff00, 0x00ff0000, 0xff000000)
	defer surf.Free()
	if err != nil {
		return nil, err
	}

	dstOffs := geom.Point2i{}
	for r, coords := range face.coords {
		dr, mask, maskpoint, advance, ok := f.Glyph(fixed.Point26_6{}, r)
		if !ok {
			panic("unreachable")
		}
		coords.Rect = geom.InitRect2iAbs(0, 0, dr.Dx(), dr.Dy()).Add(dstOffs)
		coords.ZpOffs = geom.InitPoint2i(dr.Min.X, dr.Min.Y)
		coords.Advance = advance.Round() // TODO: Fix me

		srcOffs := geom.InitPoint2i(maskpoint.X, maskpoint.Y)
		blitRune(surf, dstOffs, mask, srcOffs, coords.Rect)

		dstOffs.X += coords.Rect.W() * delta.X
		dstOffs.Y += coords.Rect.H() * delta.Y
	}

	tex, err := o.SurfaceToTexture(surf)
	tex.SetBlendMode(sdl.BLENDMODE_BLEND)
	if err != nil {
		return nil, err
	}

	_ = tex
	face.tex = tex

	return &face, nil
}

func blitRune(surf *sdl.Surface, dstOffs geom.Point2i, mask image.Image, srcOffs geom.Point2i, rect geom.Rect2i) {
	// fmt.Println("---------------------------------------")
	for y := 0; y < rect.H(); y++ {
		// s := ""
		for x := 0; x < rect.W(); x++ {
			c := mask.At(x+srcOffs.X, y+srcOffs.Y)
			r, g, b, a := c.RGBA()
			_, _, _, _ = r, g, b, a
			newColor := color.NRGBA{255, 255, 255, uint8(r)}
			setPixel(surf, x+dstOffs.X, y+dstOffs.Y, newColor)
			r, g, b, a = newColor.RGBA()
			// s += fmt.Sprintf("#%v %v %v %v - %v", r, g, b, a, newColor)
		}
		// _ = s
		// fmt.Println(s)
	}
	// fmt.Println(sdl.GetPixelFormatName(uint(surf.Format.Format)))
}

func setPixel(surf *sdl.Surface, x, y int, c color.Color) {
	pix := surf.Pixels()
	i := int32(y)*surf.Pitch + int32(x)*int32(surf.Format.BytesPerPixel)
	r, g, b, a := c.RGBA()
	if sdl.BYTEORDER == sdl.LIL_ENDIAN {
		r, g, b, a = a, b, g, r
	}
	switch surf.Format.Format {
	case sdl.PIXELFORMAT_ARGB8888:
		pix[i+0] = uint8(a >> 8)
		pix[i+1] = uint8(r >> 8)
		pix[i+2] = uint8(g >> 8)
		pix[i+3] = uint8(b >> 8)
	case sdl.PIXELFORMAT_ABGR8888:
		pix[i+0] = uint8(a >> 8)
		pix[i+1] = uint8(b >> 8)
		pix[i+2] = uint8(g >> 8)
		pix[i+3] = uint8(r >> 8)
	case sdl.PIXELFORMAT_RGB24, sdl.PIXELFORMAT_RGB888:
		pix[i+0] = uint8(r >> 8)
		pix[i+1] = uint8(g >> 8)
		pix[i+2] = uint8(b >> 8)
	case sdl.PIXELFORMAT_BGR24, sdl.PIXELFORMAT_BGR888:
		pix[i+0] = uint8(b >> 8)
		pix[i+1] = uint8(g >> 8)
		pix[i+2] = uint8(r >> 8)
	default:
		panic("unknown pixel format!")
	}
}
