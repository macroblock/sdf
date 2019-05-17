package gfx

import (
	"image/color"
)

type (
	// ITexture -
	ITexture interface {
		SetColorMod(c color.Color)
	}
)
