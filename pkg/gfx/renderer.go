package gfx

import (
	"fmt"
	"image/color"

	"github.com/macroblock/sdf/pkg/geom"
	"github.com/veandco/go-sdl2/sdl"
)

var defaultFont IFont

type (
	// Renderer -
	Renderer struct {
		offset   geom.Point2i
		bounds   geom.Rect2i
		viewport geom.Rect2i
		font     IFontFace
		r        *sdl.Renderer
	}
)

func sdlRect(x, y, w, h int) sdl.Rect {
	return sdl.Rect{X: int32(x), Y: int32(y), W: int32(w), H: int32(h)}
}

// NewRenderer -
func NewRenderer(r *sdl.Renderer) *Renderer {
	r.SetDrawBlendMode(sdl.BLENDMODE_BLEND)
	return &Renderer{r: r}
}

// Destroy -
func (o *Renderer) Destroy() error {
	err := o.r.Destroy()
	return err
}

// SetViewport -
func (o *Renderer) SetViewport(rect geom.Rect2i) {
	o.viewport = rect
	r := geom.Rect2iToSdl(rect)
	err := o.r.SetViewport(&r)
	if err != nil {
		fmt.Println(err)
	}
	_ = err
}

// SetDefaultFont -
func (o *Renderer) SetDefaultFont(font IFont) {
	defaultFont = font
}

// ResetViewport -
func (o *Renderer) ResetViewport() {
	o.viewport = geom.Rect2i{}
	err := o.r.SetViewport(nil)
	_ = err
}

// SetScale -
func (o *Renderer) SetScale(x, y float64) {
	o.r.SetScale(float32(x), float32(y))
}

// SDLRenderer -
func (o *Renderer) SDLRenderer() *sdl.Renderer {
	return o.r
}

// Present -
func (o *Renderer) Present() {
	o.r.Present()
}

// Size -
func (o *Renderer) Size() (int, int) {
	w, h, err := o.r.GetOutputSize()
	if err != nil {
		// setError(err)
		return -1, -1
	}
	return int(w), int(h)
}

// SetColor -
func (o *Renderer) SetColor(c color.RGBA) {
	r, g, b, a := c.RGBA()
	o.r.SetDrawColor(uint8(r), uint8(g), uint8(b), uint8(a))
}

// SetTextColor -
func (o *Renderer) SetTextColor(c color.RGBA) {
	font := o.font
	if font == nil {
		font = defaultFont
	}
	font.SetColor(c)
}

// ClearAll -
func (o *Renderer) ClearAll() {
	err := o.r.Clear()
	_ = err
}

// Clear -
func (o *Renderer) Clear() {
	rect := sdl.Rect{X: 0, Y: 0, W: int32(o.viewport.W()), H: int32(o.viewport.H())}
	err := o.r.FillRect(&rect)
	_ = err
}

// DrawLine -
func (o *Renderer) DrawLine(x1, y1, x2, y2 int) {
	o.r.DrawLine(int32(o.offset.X+x1), int32(o.offset.Y+y1), int32(x2), int32(y2))
}

// DrawRect -
func (o *Renderer) DrawRect(x, y, w, h int) {
	rect := sdlRect(o.offset.X+x, o.offset.Y+y, w, h)
	o.r.DrawRect(&rect)
}

// FillRect -
func (o *Renderer) FillRect(x, y, w, h int) {
	rect := sdlRect(o.offset.X+x, o.offset.Y+y, w, h)
	o.r.FillRect(&rect)
}

// DrawRect2i -
func (o *Renderer) DrawRect2i(rect geom.Rect2i) {
	r2 := geom.Rect2iToSdl(rect.Add(o.offset))
	o.r.DrawRect(&r2)
}

// FillRect2i -
func (o *Renderer) FillRect2i(rect geom.Rect2i) {
	r2 := geom.Rect2iToSdl(rect.Add(o.offset))
	o.r.FillRect(&r2)
}

// Print -
func (o *Renderer) Print(x, y int, text string) {
	font := o.font
	if font == nil {
		font = defaultFont
	}
	font.Print(x, y, text)
}