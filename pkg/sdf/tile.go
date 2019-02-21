package sdf

import (
	"fmt"
	"time"

	"github.com/macroblock/sdf/pkg/geom"
	"github.com/veandco/go-sdl2/sdl"
)

// FlipMode -
type FlipMode int

// -
const (
	FlipNone FlipMode = iota
	FlipHorizontal
	FlipVertical
)

type (
	// Tile -
	Tile struct {
		tex     *Texture
		bounds  geom.Rect2i
		bearing geom.Point2i
		flip    FlipMode
	}
)

// Copy -
func (o *Tile) Copy(x, y int, delta time.Duration) {
	if !Ok() || o == nil {
		return
	}
	bounds := o.bounds.Normalize()
	src := sdl.Rect{X: int32(bounds.X), Y: int32(bounds.Y), W: int32(bounds.W), H: int32(bounds.H)}
	x -= o.bearing.X
	y -= o.bearing.Y
	dst := sdl.Rect{X: int32(x), Y: int32(y), W: src.W, H: src.H}
	fmt.Printf("src: %v\ndst: %v\n", src, dst)
	// err := sdf.renderer.Copy(o.tex.sdltex, &src, &dst)
	err := sdf.renderer.CopyEx(o.tex.sdltex, &src, &dst, 0, nil, sdl.RendererFlip(o.flip))
	// err := sdf.renderer.Copy(o.tex.sdltex, &src, &dst)
	setError(err)
}
