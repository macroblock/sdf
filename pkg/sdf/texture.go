package sdf

import (
	"github.com/macroblock/sdf/pkg/gfx"
)

// LoadTexture -
func LoadTexture(path string) *gfx.Texture {
	if !Ok() {
		return nil
	}
	tex, err := assets.loadResource(path)
	setError(err)
	return tex
}

// func loadTexture(path string) (*gfx.Texture, error) {
// 	surf, err := img.Load(path)
// 	defer surf.Free()
// 	if err != nil {
// 		return nil, err
// 	}
// 	bounds := surf.Bounds()

// 	tex, err := sdf.renderer.CreateTextureFromSurface(surf)
// 	if err != nil {
// 		return nil, err
// 	}
// 	ret := &gfx.Texture{sdltex: tex}
// 	ret.W = bounds.Max.X - bounds.Min.X
// 	ret.H = bounds.Max.Y - bounds.Min.Y
// 	return ret, nil
// }

// func unloadTexture(tex *gfx.Texture) {
// 	err := sdl.GetError()
// 	_ = tex.sdltex.Destroy()
// 	sdl.SetError(err)
// 	tex.sdltex = nil
// }

// // Copy -
// func (o *Texture) Copy(x, y int) {
// 	if !Ok() {
// 		return
// 	}
// 	r := sdl.Rect{X: int32(x), Y: int32(y), W: int32(o.W), H: int32(o.H)}
// 	err := sdf.renderer.Copy(o.sdltex, nil, &r)
// 	setError(err)
// }
