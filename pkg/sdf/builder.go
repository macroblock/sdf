package sdf

import (
	"fmt"
	"image"

	"github.com/macroblock/sdf/pkg/gfx"
)

type (
	// TileBuilder -
	TileBuilder struct {
		prefix  string
		counter uint
	}

	// TileTemplateBuilder -
	TileTemplateBuilder struct {
		params []tileTemplateType
	}

	tileTemplateType struct {
		offs   int
		extend *image.Rectangle
		flip   gfx.FlipMode
	}
)

// BuildTileSet -
func BuildTileSet(prefix string, counter uint) *TileBuilder {
	if prefix == "" {
		prefix = "/"
	}
	return &TileBuilder{prefix, counter}
}

// Tile -
func (o *TileBuilder) Tile(offs int, extend *image.Rectangle, flip gfx.FlipMode) *TileBuilder {
	if offs < 0 {
		setError(fmt.Errorf("negative tile offset"))
		return o
	}
	name := o.genName()
	CreateTile(name, -(offs + 1), 0, extend, flip)
	return o
}

// BuildTileTemplate -
func BuildTileTemplate() *TileTemplateBuilder {
	return &TileTemplateBuilder{}
}

// Tile -
func (o *TileTemplateBuilder) Tile(offs int, extend *image.Rectangle, flip gfx.FlipMode) *TileTemplateBuilder {
	if extend != nil {
		ext := *extend
		extend = &ext
	}
	o.params = append(o.params, tileTemplateType{offs: offs, extend: extend, flip: flip})
	return o
}

// Build -
func (o *TileTemplateBuilder) Build(prefix string, baseOffs int, flip gfx.FlipMode) *TileBuilder {
	builder := BuildTileSet(prefix, 0)
	for i := range o.params {
		params := &o.params[i]
		builder.Tile(params.offs+baseOffs, params.extend, params.flip^flip)
	}
	return builder
}

func (o *TileBuilder) genName() string {
	ret := joinPaths(o.prefix, fmt.Sprintf("%03d", o.counter))
	ret = AbsTilePath(ret)
	o.counter++
	return ret
}
