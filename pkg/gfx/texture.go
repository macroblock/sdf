package gfx

import (
	"fmt"
	"image"
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
	r := sdlRect(x-o.offset.X, y-o.offset.Y, tex.W, tex.H)
	err := o.r.Copy(tex.sdltex, nil, &r)
	_ = err
	// setError(err)
}

// CopyRegion -
func (o *Renderer) CopyRegion(tex *Texture, src, dst geom.Rect2i) {
	r1 := geom.Rect2iToSdl(src)
	r2 := geom.Rect2iToSdl(dst.Sub(o.offset))
	err := o.r.Copy(tex.sdltex, &r1, &r2)
	_ = err
	// setError(err)
}

// CopyRegionEx -
func (o *Renderer) CopyRegionEx(tex *Texture, src, dst geom.Rect2i, flip FlipMode) {
	r1 := geom.Rect2iToSdl(src)
	r2 := geom.Rect2iToSdl(dst.Sub(o.offset))
	// err := o.r.Copy(tex.sdltex, &r1, &r2)
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
		s := ""
		for x := 0; x < size.X; x++ {
			c := img.At(x, y)
			r, g, b, a := c.RGBA()
			_, _, _, _ = r, g, b, a
			// newColor := color.Gray{uint8(r >> 8)}
			newColor := color.RGBA{255, 255, 255, uint8(r)}
			s += fmt.Sprintf("%v", newColor)
			rgba.Set(x, y, newColor)
		}
		_ = s
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
		W: int(w),
		H: int(h),
		// W: img.Bounds().Dx(),
		// H: img.Bounds().Dy(),
		// W:      surf.Bounds().Dx(),
		// H:      surf.Bounds().Dy(),
		sdltex: sdltex,
	}
	return tex, err
}

// SurfaceToTexture - it's hack
func (o *Renderer) SurfaceToTexture(surf *sdl.Surface) (*Texture, error) {
	surf.SaveBMP("test.bmp")
	sdltex, err := o.SDLRenderer().CreateTextureFromSurface(surf)
	// defer surf.Free() // TODO: WHY ?
	if err != nil {
		return nil, err
	}
	_, _, w, h, err := sdltex.Query()
	if err != nil {
		return nil, err
	}
	tex := &Texture{
		W:      int(w),
		H:      int(h),
		sdltex: sdltex,
	}
	return tex, nil
}
