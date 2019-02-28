package sdf

import "github.com/macroblock/sdf/pkg/geom"

type (
	// TileBuilder -
	TileBuilder struct {
		prefix string
	}
)

// BuildTile -
func BuildTile(prefix string) *TileBuilder {
	return &TileBuilder{prefix}
}

// XY -
func (o *TileBuilder) XY(name string, x0, y0 int, extend *geom.Rect2i, flip FlipMode) *TileBuilder {

}
