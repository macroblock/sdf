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
		elems     map[string]IElem
	}
)

// LoadTileSheet - if tileW or/and tileH is negative values,
// it will be interpreted as a number of tiles per width
// or/and per height respectively.
func LoadTileSheet(tileW, tileH int, path string) *TileSheet {
	if !Ok() {
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
	ret.elems = map[string]IElem{}
	return ret
}

// InitTile -
func (o *TileSheet) InitTile(name string, x0, y0 int, extend *geom.Rect2i, flip FlipMode) *Tile {
	if !Ok() {
		return nil
	}
	if _, ok := o.elems[name]; ok {
		setError(fmt.Errorf("element %q already exists", name))
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
		tex:     o.tex,
		bounds:  bounds.Mul(o.tileSize),
		bearing: bearing.Mul(o.tileSize),
		flip:    flip,
	}

	o.elems[name] = tile
	return tile
}

// InitSprite -
func (o *TileSheet) InitSprite(name string, duration float64, tileName string) *Sprite {
	if !Ok() {
		return nil
	}
	if _, ok := o.elems[name]; ok {
		setError(fmt.Errorf("element %q already exists", name))
		return nil
	}
	tile := o.Tile(tileName)
	if !Ok() {
		return nil
	}
	sprite := newSprite(o, duration, tile, nil)
	o.elems[name] = sprite
	return sprite
}

// Elem -
func (o *TileSheet) Elem(name string) IElem {
	if !Ok() {
		return nil
	}
	elem, ok := o.elems[name]
	if !ok {
		setError(fmt.Errorf("elem %q was not found", name))
		return nil
	}
	return elem
}

// Tile -
func (o *TileSheet) Tile(name string) *Tile {
	i := o.Elem(name)
	tile, ok := i.(*Tile)
	if !ok {
		setError(fmt.Errorf("%q is not a tile", name))
		return nil
	}
	return tile
}

// Sprite -
func (o *TileSheet) Sprite(name string) *Sprite {
	i := o.Elem(name)
	sprite, ok := i.(*Sprite)
	if !ok {
		setError(fmt.Errorf("%q is not a sprite", name))
		return nil
	}
	return sprite
}
