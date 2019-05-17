package gfx

import (
	"fmt"
	"image"
	"image/color"

	"github.com/veandco/go-sdl2/sdl"
	"golang.org/x/image/font"
)

type (
	// Renderer -
	Renderer struct {
		offset   image.Point
		bounds   image.Rectangle
		viewport image.Rectangle
		font     font.Face //*HWFace
		// face        IFace
		defaultFont font.Face //*HWFace
		textColor   color.Color
		r           *sdl.Renderer
	}
)

func sdlRectI(x, y, w, h int) sdl.Rect {
	return sdl.Rect{X: int32(x), Y: int32(y), W: int32(w), H: int32(h)}
}

func sdlRect(r image.Rectangle) sdl.Rect {
	return sdl.Rect{X: int32(r.Min.X), Y: int32(r.Min.Y), W: int32(r.Dx()), H: int32(r.Dy())}
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
func (o *Renderer) SetViewport(rect image.Rectangle) {
	o.viewport = rect
	r := sdlRect(rect)
	err := o.r.SetViewport(&r)
	if err != nil {
		fmt.Println(err)
	}
	_ = err
	fmt.Println("viewport: ", rect)
}

// SetOffset -
func (o *Renderer) SetOffset(offset image.Point) {
	o.offset = offset
}

// SetDefaultFont -
func (o *Renderer) SetDefaultFont(font font.Face) {
	o.defaultFont = font
}

// ResetViewport -
func (o *Renderer) ResetViewport() {
	o.viewport = image.Rectangle{}
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
func (o *Renderer) SetColor(c color.Color) {
	o.r.SetDrawColor(RGBA8(c))
}

// SetTextColor -
func (o *Renderer) SetTextColor(c color.Color) {
	o.textColor = c
}

// // SetTextColor -
// func (o *Renderer) SetTextColor(c color.RGBA) {
// 	font := o.font
// 	if font == nil {
// 		font = defaultFont
// 	}
// 	font.SetColor(c)
// }

// ClearAll -
func (o *Renderer) ClearAll() {
	err := o.r.Clear()
	_ = err
}

// Clear -
func (o *Renderer) Clear() {
	rect := sdl.Rect{X: 0, Y: 0, W: int32(o.viewport.Dx()), H: int32(o.viewport.Dy())}
	err := o.r.FillRect(&rect)
	_ = err
}

// DrawLine -
func (o *Renderer) DrawLine(x1, y1, x2, y2 int) {
	o.r.DrawLine(int32(x1-o.offset.X), int32(y1-o.offset.Y), int32(x2), int32(y2))
}

// DrawRect -
func (o *Renderer) DrawRect(x, y, w, h int) {
	rect := sdlRectI(x-o.offset.X, y-o.offset.Y, w, h)
	o.r.DrawRect(&rect)
}

// FillRect -
func (o *Renderer) FillRect(x, y, w, h int) {
	rect := sdlRectI(x-o.offset.X, y-o.offset.Y, w, h)
	o.r.FillRect(&rect)
}

// DrawRect2i -
func (o *Renderer) DrawRect2i(rect image.Rectangle) {
	r2 := sdlRect(rect.Sub(o.offset))
	o.r.DrawRect(&r2)
}

// FillRect2i -
func (o *Renderer) FillRect2i(rect image.Rectangle) {
	r2 := sdlRect(rect.Sub(o.offset))
	o.r.FillRect(&r2)
}

// Print -
func (o *Renderer) Print(x, y int, text string) {
	font := o.font
	if font == nil {
		font = o.defaultFont
	}
	// font.Print(x, y, text)
}
