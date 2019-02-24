package sdf

import (
	"fmt"

	"github.com/macroblock/sdf/pkg/geom"
)

type (
	// TileSheet -
	TileSheet struct {
		tex       *Texture
		tileSize  geom.Point2i
		tilesPerW int
		// tiles     map[string]*Tile
	}
)

// LoadTileSheet - if tileW or/and tileH is negative values,
// it will be interpreted as a number of tiles per width
// or/and per height respectively.
func LoadTileSheet(name string, tileW, tileH int, path string) *TileSheet {
	if !Ok() {
		return nil
	}
	if tileSheetExists(name) {
		setError(fmt.Errorf("tilesheet %q already exists", name))
		return nil
	}
	tex := LoadTexture(path)
	if !Ok() {
		return nil
	}
	tilesPerW := tex.W / tileW
	if tileW < 0 {
		tilesPerW = -tileW
		tileW = tex.W / -tileW
	}
	if tileH < 0 {
		tileH = tex.H / -tileH
	}
	ret := &TileSheet{}
	ret.tex = tex
	ret.tileSize = geom.InitPoint2i(tileW, tileH)
	ret.tilesPerW = tilesPerW
	// ret.tiles = map[string]*Tile{}
	assets.sheets[name] = ret
	assets.currentTileSheet = ret
	return ret
}

// initTile -
func (o *TileSheet) initTile(name string, x0, y0 int, extend *geom.Rect2i, flip FlipMode) *Tile {
	if !Ok() {
		return nil
	}
	if tileExists(name) {
		setError(fmt.Errorf("tile %q already exists", name))
		return nil
	}
	if extend == nil {
		extend = &geom.Rect2i{}
		*extend = geom.InitRect2i(0, 0, 1, 1)
	}

	origin := geom.InitPoint2i(x0, y0)
	bounds := extend.Add(origin)
	bearing := origin.SubInt(bounds.X, bounds.Y)
	tile := &Tile{
		tex:    o.tex,
		bounds: bounds.Mul(o.tileSize),
		pivot:  bearing.Mul(o.tileSize),
		flip:   flip,
	}

	assets.tiles[name] = tile
	return tile
}

// Tile -
// func (o *TileSheet) Tile(name string) *Tile {
// 	if !Ok() {
// 		return nil
// 	}
// 	tile, ok := assets.tiles[name]
// 	if !ok {
// 		setError(fmt.Errorf("elem %q was not found", name))
// 		return nil
// 	}
// 	return tile
// }
