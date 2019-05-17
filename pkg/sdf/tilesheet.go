package sdf

import (
	"fmt"
	"image"

	"github.com/macroblock/sdf/pkg/gfx"
)

type (
	// TileSheet -
	TileSheet struct {
		name      string
		tex       *gfx.Texture
		tileSize  image.Point
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
	ret := &TileSheet{name: name}
	ret.tex = tex
	ret.tileSize = image.Pt(tileW, tileH)
	ret.tilesPerW = tilesPerW
	// ret.tiles = map[string]*Tile{}
	assets.sheets[name] = ret
	assets.currentTileSheet = ret
	return ret
}

func scaleRect(r image.Rectangle, scale image.Point) image.Rectangle {
	r.Min.X *= scale.X
	r.Min.Y *= scale.Y
	r.Max.X *= scale.X
	r.Max.Y *= scale.Y
	return r
}

// InitTile - when x, y > 0 its behavior as usual.
// When one of the ordinate < 0 then its modulus means an offset in current tile sheet.
// In this case other ordinate must be zero, otherwise error will be raised.
// In all cases count of elements started from 1 (not from zero).
func (o *TileSheet) InitTile(name string, x, y int, extend *image.Rectangle, flip gfx.FlipMode) *Tile {
	if !Ok() {
		return nil
	}
	if tileExists(name) {
		setError(fmt.Errorf("tile %q already exists", name))
		return nil
	}
	switch {
	default:
		setError(fmt.Errorf("init tile %v, x=%v y=%v: incorrect x, y coordinate schema", name, x, y))
	case x < 0 && y == 0:
		x = -x - 1
		y = x / o.tilesPerW
		x = x % o.tilesPerW
	case x < 0 && y == 0:
		y = -y - 1
		x = y % o.tilesPerW
		y = y / o.tilesPerW
	case x > 0 && y > 0:
		x--
		y--
	} // switch

	if extend == nil {
		extend = &image.Rectangle{}
		*extend = image.Rect(0, 0, 1, 1)
	}

	origin := image.Pt(x, y)
	bounds := extend.Add(origin)
	bearing := origin.Sub(bounds.Min)
	tile := &Tile{
		tex:    o.tex,
		bounds: scaleRect(bounds, o.tileSize),
		pivot:  image.Pt(bearing.X*o.tileSize.X, bearing.Y*o.tileSize.Y),
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
