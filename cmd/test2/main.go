package main

import (
	"fmt"
	"image/color"

	"github.com/macroblock/sdf/pkg/sdf"
	"github.com/macroblock/sdf/pkg/types"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	cellSize = 16
	gridSize = 7
)

type game struct {
	gridX, gridY int
	grid         *types.Grid
}

// Init -
func (o *game) Init() {
	o.gridX = 50
	o.gridY = 50
	o.grid = types.NewGrid(gridSize, gridSize, 0)
	o.grid.Set(1, 1, -1)
}

// CleanUp -
func (o *game) CleanUp() {
	fmt.Printf("cleaned up\n")
}

// HandleEvent -
func (o *game) HandleEvent(ev sdf.IEvent) {
	// fmt.Printf("%v\n", ev)
	x, y := -1, -1
	val := 0
	switch ev := ev.(type) {
	case *sdf.MouseClickEvent:
		if ev.Pressed {
			x = ev.X
			y = ev.Y
		}
		if ev.Button == 1 {
			val = 1
		}
	case *sdf.MouseMotionEvent:
		if ev.Buttons == 0 {
			break
		}
		x = ev.X
		y = ev.Y
		if ev.Buttons&1 != 0 {
			val = 1
		}
	}
	x = (x - o.gridX) / cellSize
	y = (y - o.gridY) / cellSize
	o.grid.Set(x, y, val)
}

// Render
func (o *game) Render() {
	o.drawGrid()
}

func (o *game) drawGrid() {
	c1 := color.RGBA{255, 255, 0, 127}
	c2 := color.RGBA{0, 255, 255, 127}
	r := sdf.Renderer()
	r.SetDrawBlendMode(sdl.BLENDMODE_BLEND)
	// w, h := sdf.Size()
	pos := 0
	for pos < gridSize {
		x := o.gridX + pos*cellSize
		y := o.gridY
		h := gridSize * cellSize
		drawLine(r, x, y, x, y+h-1, c1)
		x += cellSize
		drawLine(r, x-1, y, x-1, y+h-1, c2)
		pos++
	}
	pos = 0
	for pos < gridSize {
		x := o.gridX
		y := o.gridY + pos*cellSize
		w := gridSize * cellSize
		drawLine(r, x, y, x+w-1, y, c1)
		y += cellSize
		drawLine(r, x, y-1, x+w-1, y-1, c2)
		pos++
	}
	w, h := o.grid.Size()
	for j := 0; j < h; j++ {
		for i := 0; i < w; i++ {
			c := color.RGBA{0, 0, 0, 255}
			if o.grid.Get(i, j).(int) != 0 {
				c = color.RGBA{125, 0, 0, 255}
			}
			fillRect(r, o.gridX+i*cellSize+1, o.gridY+j*cellSize+1, cellSize-2, cellSize-2, c)
		}
	}
	drawLine(r, -1, -1, -1, -1, color.RGBA{0, 0, 0, 255})
}

func fillRect(rend *sdl.Renderer, x, y, w, h int, c color.Color) {
	r, g, b, a := c.RGBA()
	rend.SetDrawColor(uint8(r), uint8(g), uint8(b), uint8(a))
	rect := sdl.Rect{X: int32(x), Y: int32(y), W: int32(w), H: int32(h)}
	rend.FillRect(&rect)
}

func drawLine(rend *sdl.Renderer, x1, y1, x2, y2 int, c color.Color) {
	r, g, b, a := c.RGBA()
	rend.SetDrawColor(uint8(r), uint8(g), uint8(b), uint8(a))
	rend.DrawLine(int32(x1), int32(y1), int32(x2), int32(y2))
}

func main() {
	err := sdf.Run(&game{})

	if err != nil {
		fmt.Println(err)
	}
}
