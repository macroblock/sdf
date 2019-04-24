package main

import (
	"fmt"
	"image/color"

	"github.com/macroblock/sdf/pkg/event"
	"github.com/macroblock/sdf/pkg/geom"
	"github.com/macroblock/sdf/pkg/sdf"
	"github.com/macroblock/sdf/pkg/types"
	"github.com/macroblock/sdf/pkg/ui"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	cellSize = 16
	gridSize = 7
)

type game struct {
	gridX, gridY int
	grid         *types.Grid
	ui           *ui.UI
}

var (
	xxx, yyy int
	panel    *ui.Panel
)

// Init -
func (o *game) Init() {
	o.gridX = 50
	o.gridY = 50
	o.grid = types.NewGrid(gridSize, gridSize, 0)
	o.grid.Set(1, 1, -1)

	xxx, yyy = 50, 50
	panel = ui.NewButton().SetBounds(geom.InitRect2i(xxx, yyy, 100, 100))
	o.ui = ui.NewUI(sdf.Renderer())
	o.ui.SetBounds(geom.InitRect2i(100, 100, 500, 250))
	o.ui.AddChildren(
		panel,
		ui.NewPanel().SetBounds(geom.InitRect2i(200, 50, 100, 25)),
	)
	fmt.Println("ui:\n", o.ui)
}

// CleanUp -
func (o *game) CleanUp() {
	fmt.Printf("cleaned up\n")
}

// HandleEvent -
func (o *game) HandleEvent(ev event.IEvent) {
	// fmt.Printf("%v\n", ev)
	x, y := -1, -1
	val := 0
	switch ev := ev.(type) {
	case *event.MouseClick:
		if ev.Pressed {
			x = ev.X
			y = ev.Y
		}
		if ev.Button == 1 {
			val = 1
		}
	case *event.MouseMotion:
		if ev.Buttons == 0 {
			break
		}
		x = ev.X
		y = ev.Y
		if ev.Buttons&1 != 0 {
			val = 1
		}
	case *event.Keyboard:
		switch ev.Key {
		case sdl.SCANCODE_ESCAPE:
			sdf.Quit()
		case sdl.SCANCODE_LEFT:
			xxx--
			once = true
			panel.SetPos(geom.InitPoint2i(xxx, yyy))
		case sdl.SCANCODE_RIGHT:
			xxx++
			once = true
			panel.SetPos(geom.InitPoint2i(xxx, yyy))
		}
	}
	x = (x - o.gridX) / cellSize
	y = (y - o.gridY) / cellSize
	o.grid.Set(x, y, val)
}

var once = true

// Render
func (o *game) Render() {
	o.drawGrid()
	// o.ui.SetBounds(geom.InitRect2i(100, 100, 50, 79))
	w, h := sdf.Renderer().Size()
	if once {
		o.ui.DrawScheme(geom.InitPoint2i(0, 0), geom.InitRect2i(0, 0, w, h))
		once = false
	}
}

func (o *game) drawGrid() {
	c1 := color.RGBA{255, 255, 0, 127}
	c2 := color.RGBA{0, 255, 255, 127}
	r := sdf.Renderer()
	// w, h := sdf.Size()
	pos := 0
	for pos < gridSize {
		x := o.gridX + pos*cellSize
		y := o.gridY
		h := gridSize * cellSize
		r.SetColor(c1)
		r.DrawLine(x, y, x, y+h-1)
		x += cellSize
		r.SetColor(c2)
		r.DrawLine(x-1, y, x-1, y+h-1)
		pos++
	}
	pos = 0
	for pos < gridSize {
		x := o.gridX
		y := o.gridY + pos*cellSize
		w := gridSize * cellSize
		r.SetColor(c1)
		r.DrawLine(x, y, x+w-1, y)
		y += cellSize
		r.SetColor(c2)
		r.DrawLine(x, y-1, x+w-1, y-1)
		pos++
	}
	w, h := o.grid.Size()
	for j := 0; j < h; j++ {
		for i := 0; i < w; i++ {
			c := color.RGBA{0, 0, 0, 255}
			if o.grid.Get(i, j).(int) != 0 {
				c = color.RGBA{125, 0, 0, 255}
			}
			r.SetColor(c)
			r.FillRect(o.gridX+i*cellSize+1, o.gridY+j*cellSize+1, cellSize-2, cellSize-2)
		}
	}
}

func main() {
	err := sdf.Run(&game{})

	if err != nil {
		fmt.Println(err)
	}
}
