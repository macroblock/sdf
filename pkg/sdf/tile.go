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

var tilePrefix = ""

type (
	// Tile -
	Tile struct {
		tex    *Texture
		bounds geom.Rect2i
		pivot  geom.Point2i
		flip   FlipMode
	}
)

// SetTilePrefix -
func SetTilePrefix(prefix string) {
	tilePrefix = prefix
}

// CreateTile -
func CreateTile(name string, x0, y0 int, extend *geom.Rect2i, flip FlipMode) *Tile {
	if !Ok() {
		return nil
	}
	if assets.currentTileSheet == nil {
		setError(fmt.Errorf("current tile sheet is nil"))
		return nil
	}
	if len(tilePrefix) > 0 {
		name = tilePrefix + constNameSeparator + name
	}
	name = assets.currentTileSheet.name + constNameSeparator + name
	if tileExists(name) {
		setError(fmt.Errorf("tile %q already exists", name))
		return nil
	}
	tile := assets.currentTileSheet.initTile(name, x0, y0, extend, flip)
	if tile == nil {
		return nil
	}
	assets.tiles[name] = tile
	return tile
}

// Update -
func (o *Tile) Update(delta time.Duration) bool {
	return false
}

// Copy -
func (o *Tile) Copy(x, y int) {
	if !Ok() || o == nil {
		return
	}
	bounds := o.bounds.Normalize()
	src := sdl.Rect{X: int32(bounds.X), Y: int32(bounds.Y), W: int32(bounds.W), H: int32(bounds.H)}
	x -= o.pivot.X
	y -= o.pivot.Y
	dst := sdl.Rect{X: int32(x), Y: int32(y), W: src.W, H: src.H}
	// fmt.Printf("src: %v\ndst: %v\n", src, dst)
	// err := sdf.renderer.Copy(o.tex.sdltex, &src, &dst)
	err := sdf.renderer.CopyEx(o.tex.sdltex, &src, &dst, 0, nil, sdl.RendererFlip(o.flip))
	// err := sdf.renderer.Copy(o.tex.sdltex, &src, &dst)
	setError(err)
}
