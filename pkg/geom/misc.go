package geom

import "github.com/veandco/go-sdl2/sdl"

// Rect2iToSdl -
func Rect2iToSdl(r Rect2i) sdl.Rect {
	return sdl.Rect{
		X: int32(r.A.X),
		Y: int32(r.A.Y),
		W: int32(r.B.X - r.A.X),
		H: int32(r.B.Y - r.A.Y),
	}
}
