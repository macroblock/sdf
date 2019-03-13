package types

import "github.com/macroblock/sdf/pkg/misc"

type (
	// Grid -
	Grid struct {
		w, h int
		data [][]interface{}
	}
)

// NewGrid -
func NewGrid(w, h int, initialValue interface{}) *Grid {
	w = misc.MaxInt(0, w)
	h = misc.MaxInt(0, h)
	data := make([][]interface{}, h)
	for j := range data {
		line := make([]interface{}, w)
		for i := range line {
			line[i] = initialValue
		}
		data[j] = line
	}
	grid := &Grid{w: w, h: h, data: data}
	return grid
}

// Data -
func (o *Grid) Data() [][]interface{} {
	return o.data
}

// Size -
func (o *Grid) Size() (int, int) {
	return o.w, o.h
}

// Set -
func (o *Grid) Set(x, y int, val interface{}) {
	if x < 0 || x >= o.w || y < 0 || y >= o.h {
		return
	}
	o.data[y][x] = val
}

// Get -
func (o *Grid) Get(x, y int) interface{} {
	if x < 0 || x >= o.w || y < 0 || y >= o.h {
		return nil
	}
	return o.data[y][x]
}
