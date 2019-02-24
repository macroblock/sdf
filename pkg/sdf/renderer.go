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

// Size -
func Size() (int, int) {
	if !Ok() {
		return -1, -1
	}
	w, h, err := sdf.renderer.GetOutputSize()
	if err != nil {
		setError(err)
		return -1, -1
	}
	return int(w), int(h)
}
