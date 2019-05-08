package gfx

import (
	"image/color"

	"github.com/macroblock/sdf/pkg/geom"
)

type (
	// ITexture -
	ITexture interface {
		SetColorMod(c color.Color)
	}
	// IFace -
	IFace interface {
		Glyph(rune) (ITexture, geom.Rect2i, geom.Point2i, int, bool)
	}
)
