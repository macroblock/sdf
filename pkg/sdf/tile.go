package sdf

import (
	"fmt"
	"image"
	"time"

	"github.com/macroblock/sdf/pkg/gfx"
)

type (
	// Tile -
	Tile struct {
		tex    *gfx.Texture
		bounds image.Rectangle
		pivot  image.Point
		flip   gfx.FlipMode
	}
)

// CreateTile - see TileSheet.InitTile function
func CreateTile(name string, x, y int, extend *image.Rectangle, flip gfx.FlipMode) *Tile {
	if !Ok() {
		return nil
	}
	if assets.currentTileSheet == nil {
		setError(fmt.Errorf("current tile sheet is nil"))
		return nil
	}
	name = joinPaths("/", name)
	if tileExists(name) {
		setError(fmt.Errorf("tile %q already exists", name))
		return nil
	}
	tile := assets.currentTileSheet.InitTile(name, x, y, extend, flip)
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
	src := o.bounds.Canon()
	x -= o.pivot.X
	y -= o.pivot.Y
	dst := image.Rect(x, y, src.Dx(), src.Dy())
	sdf.renderer.CopyRegionEx(o.tex, src, dst, o.flip)
	err := error(nil)
	setError(err)
}
