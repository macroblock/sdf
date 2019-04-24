package gfx

import "image/color"

// RGBA8 -
func RGBA8(c color.Color) (uint8, uint8, uint8, uint8) {
	if c, ok := c.(color.RGBA); ok {
		return c.R, c.G, c.B, c.A
	}
	r, g, b, a := c.RGBA()
	return uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)
}
