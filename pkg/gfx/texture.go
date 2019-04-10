package gfx

import (
	"image/color"

	"github.com/macroblock/sdf/pkg/general"
	"github.com/macroblock/sdf/pkg/geom"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

// Texture -
type Texture struct {
	general.RefCounter
	general.URI
	sdltex *sdl.Texture
	W, H   int
}

// // Copy -
// func (o *Texture) Copy(x, y int) {
// 	if !Ok() {
// 		return
// 	}
// 	r := sdl.Rect{X: int32(x), Y: int32(y), W: int32(o.W), H: int32(o.H)}
// 	err := sdf.renderer.Copy(o.sdltex, nil, &r)
// 	setError(err)
// }

// SetColorMod -
func (o *Texture) SetColorMod(c color.RGBA) {
	o.sdltex.SetColorMod(c.R, c.G, c.B)
	o.sdltex.SetAlphaMod(c.A)
}

// LoadTexture -
func (o *Renderer) LoadTexture(path string) (*Texture, error) {
	surf, err := img.Load(path)
	defer surf.Free()
	if err != nil {
		return nil, err
	}
	bounds := surf.Bounds()

	tex, err := o.r.CreateTextureFromSurface(surf)
	if err != nil {
		return nil, err
	}
	ret := &Texture{sdltex: tex}
	ret.W = bounds.Max.X - bounds.Min.X
	ret.H = bounds.Max.Y - bounds.Min.Y
	return ret, nil
}

// UnloadTexture -
func (o *Renderer) UnloadTexture(tex *Texture) {
	err := sdl.GetError()
	_ = tex.sdltex.Destroy()
	sdl.SetError(err)
	tex.sdltex = nil
}

// Copy -
func (o *Renderer) Copy(tex *Texture, x, y int) {
	r := sdlRect(o.offset.X+x, o.offset.Y+y, tex.W, tex.H)
	err := o.r.Copy(tex.sdltex, nil, &r)
	_ = err
	// setError(err)
}

// CopyRegion -
func (o *Renderer) CopyRegion(tex *Texture, src, dst geom.Rect2i) {
	r1 := geom.Rect2iToSdl(src)
	r2 := geom.Rect2iToSdl(dst.Add(o.offset))
	err := o.r.Copy(tex.sdltex, &r1, &r2)
	_ = err
	// setError(err)
}

// CopyRegionEx -
func (o *Renderer) CopyRegionEx(tex *Texture, src, dst geom.Rect2i, flip FlipMode) {
	r1 := geom.Rect2iToSdl(src)
	r2 := geom.Rect2iToSdl(dst.Add(o.offset))
	// err := o.r.Copy(tex.sdltex, &r1, &r2)
	err := o.r.CopyEx(tex.sdltex, &r1, &r2, 0, nil, sdl.RendererFlip(flip))
	_ = err
	// setError(err)
}
