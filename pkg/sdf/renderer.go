package sdf

import "github.com/veandco/go-sdl2/sdl"

// SetScale -
func SetScale(x, y float64) {
	if !Ok() {
		return
	}
	sdf.renderer.SetScale(float32(x), float32(y))
}

// Renderer -
func Renderer() *sdl.Renderer {
	return sdf.renderer
}
