package gfx

import (
	"fmt"
	"image"
	"image/color"

	"github.com/macroblock/sdf/pkg/general"
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

// ColorModel -
func (o *Texture) ColorModel() color.Model {
	panic("this is just a stub to use as Image interface")
}

// Bounds -
func (o *Texture) Bounds() image.Rectangle {
	panic("this is just a stub to use as Image interface")
}

// At -
func (o *Texture) At(x, y int) color.Color {
	panic("this is just a stub to use as Image interface")
}

// SetColorMod -
func (o *Texture) SetColorMod(c color.Color) {
	r, g, b, a := RGBA8(c)
	o.sdltex.SetColorMod(r, g, b)
	o.sdltex.SetAlphaMod(a)
}

// SetBlendMode -
func (o *Texture) SetBlendMode(mode sdl.BlendMode) {
	o.sdltex.SetBlendMode(mode)
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
	r := sdlRectI(x-o.offset.X, y-o.offset.Y, tex.W, tex.H)
	err := o.r.Copy(tex.sdltex, nil, &r)
	_ = err
	// setError(err)
}

// CopyRegion -
func (o *Renderer) CopyRegion(tex *Texture, src, dst image.Rectangle) {
	r1 := sdlRect(src)
	r2 := sdlRect(dst.Sub(o.offset))
	err := o.r.Copy(tex.sdltex, &r1, &r2)
	_ = err
	// setError(err)
}

// CopyRegionEx -
func (o *Renderer) CopyRegionEx(tex *Texture, src, dst image.Rectangle, flip FlipMode) {
	r1 := sdlRect(src)
	r2 := sdlRect(dst.Sub(o.offset))
	err := o.r.CopyEx(tex.sdltex, &r1, &r2, 0, nil, sdl.RendererFlip(flip))
	_ = err
	// setError(err)
}

// ImageToTexture -
func (o *Renderer) ImageToTexture(img image.Image) (*Texture, error) {
	rgba := image.NewRGBA(img.Bounds())
	size := img.Bounds().Size()
	surf, err := sdl.CreateRGBSurface(0, int32(size.X), int32(size.Y), 32, 0x000000ff, 0x0000ff00, 0x00ff0000, 0xff000000)
	if err != nil {
		return nil, err
	}
	rgba.Pix = surf.Pixels()

	for y := 0; y < size.Y; y++ {
		for x := 0; x < size.X; x++ {
			c := img.At(x, y)
			rgba.Set(x, y, c)
		}
	}
	sdltex, err := o.SDLRenderer().CreateTextureFromSurface(surf)
	surf.Free()
	if err != nil {
		return nil, fmt.Errorf("Renderer.ImageToTexture CeateTuxtureFromSurface: %v", err)
	}
	_, _, w, h, err := sdltex.Query()
	if err != nil {
		return nil, fmt.Errorf("Renderer.ImageToTexture Query: %v", err)
	}
	tex := &Texture{
		W:      int(w),
		H:      int(h),
		sdltex: sdltex,
	}
	return tex, err
}
