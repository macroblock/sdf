package main

import (
	"fmt"
	"image/color"
	"sort"
	"time"

	"github.com/macroblock/sdf/pkg/fonts/pixfm5x9normal"
	"github.com/macroblock/sdf/pkg/sdf"
	"github.com/veandco/go-sdl2/sdl"
)

var (
	tex       *sdf.Texture
	font      *sdf.PixelFont
	tileSheet *sdf.TileSheet
	// sprite0, sprite1, sprite2                *sdf.Sprite
	anim0 *sdf.Animation
	hero  *Hero
)

type game struct {
}

func (o *game) Init() {
	fmt.Printf("initialized\n")

	tex = sdf.LoadTexture("../../assets/testsheet.png")

	font = sdf.CreatePixelFont(pixfm5x9normal.Font)
	// font.SetScale(2)

	sdf.SetScale(3, 3)

	tileSheet = sdf.LoadTileSheet("ff", -8, -6, "../../assets/testsheet.png")

	sdf.SetTilePath("/ff/move")
	motion := sdf.BuildTileTemplate().
		Tile(0, nil, 0).
		Tile(1, nil, 0).
		Tile(2, nil, 0).
		Tile(1, nil, 0)

	motion.Build("S", 0, 0)
	motion.Build("N", 3, 0)
	motion.Build("W", 6, 0)
	motion.Build("E", 6, sdf.FlipHorizontal)

	sdf.SetTilePath("/ff")
	sdf.BuildTileSet("idle", 0).
		Tile(12, nil, 0).
		Tile(15, nil, 0)

	anim0 = sdf.CreateAnimation("test move").Plain("/ff/move/S*").StretchTo(1.0)
	sdf.CreateAnimation("hero idle").Plain("/ff/idle*").StretchTo(1.0)
	sdf.CreateAnimation("hero move N").Plain("/ff/move/N*").StretchTo(1.0)
	sdf.CreateAnimation("hero move S").Plain("/ff/move/S*").StretchTo(1.0)
	sdf.CreateAnimation("hero move W").Plain("/ff/move/W*").StretchTo(1.0)
	sdf.CreateAnimation("hero move E").Plain("/ff/move/E*").StretchTo(1.0)

	// tileSheet = sdf.LoadTileSheet("girl", -4, -5, "../../assets/girl-03.png")

	// sdf.SetTilePrefix("s move")
	// sdf.CreateTile("0", 0, 0, nil, 0)
	// sdf.CreateTile("1", 1, 0, nil, 0)
	// sdf.CreateTile("2", 3, 0, nil, 0)

	// sdf.SetTilePrefix("w move")
	// sdf.CreateTile("0", 0, 1, nil, 0)
	// sdf.CreateTile("1", 1, 1, nil, 0)
	// sdf.CreateTile("2", 3, 1, nil, 0)

	// sdf.SetTilePrefix("e move")
	// sdf.CreateTile("0", 0, 1, nil, sdf.FlipHorizontal)
	// sdf.CreateTile("1", 1, 1, nil, sdf.FlipHorizontal)
	// sdf.CreateTile("2", 3, 1, nil, sdf.FlipHorizontal)

	// sdf.SetTilePrefix("n move")
	// sdf.CreateTile("0", 0, 2, nil, 0)
	// sdf.CreateTile("1", 1, 2, nil, 0)
	// sdf.CreateTile("2", 3, 2, nil, 0)

	// sdf.SetTilePrefix("idle")
	// sdf.CreateTile("0", 0, 3, nil, 0)
	// sdf.CreateTile("1", 1, 3, nil, 0)
	// sdf.CreateTile("2", 2, 3, nil, 0)
	// sdf.CreateTile("3", 3, 3, nil, 0)

	// sdf.CreateAnimation("girl/idle").Sequence("/0", "/1", "/2", "/3").StretchTo(0.5)
	// sdf.CreateAnimation("girl move N").Sequence("girl/n move/0", "girl/n move/1", "girl/n move/0", "girl/n move/2").StretchTo(0.5)
	// sdf.CreateAnimation("girl move S").Sequence("girl/s move/0", "girl/s move/1", "girl/s move/0", "girl/s move/2").StretchTo(0.5)
	// sdf.CreateAnimation("girl move W").Sequence("girl/w move/0", "girl/w move/1", "girl/w move/0", "girl/w move/2").StretchTo(0.5)
	// sdf.CreateAnimation("girl move E").Sequence("girl/e move/0", "girl/e move/1", "girl/e move/0", "girl/e move/2").StretchTo(0.5)

	hero = NewHero(32, 32)
	// hero = NewGirl(32, 32)
	// sdf.StopTextInput()

	list := sdf.ListAssets()
	sort.Strings(list)
	for _, s := range list {
		fmt.Println(s)
	}
}

// Hero -
type Hero struct {
	*sdf.GameObject
	tween sdf.Tween
	x, y  int
}

// NewHero -
func NewHero(x, y int) *Hero {
	hero := &Hero{x: x, y: y}
	hero.GameObject = sdf.NewGameObject("hero").
		AddAnimation("idle", "hero idle").
		AddAnimation("move N", "hero move N").
		AddAnimation("move S", "hero move S").
		AddAnimation("move W", "hero move W").
		AddAnimation("move E", "hero move E")
	hero.Play("move E")
	return hero
}

// NewGirl -
func NewGirl(x, y int) *Hero {
	hero := &Hero{x: x, y: y}
	hero.GameObject = sdf.NewGameObject("girl").
		AddAnimation("idle", "girl idle").
		AddAnimation("move N", "girl move N").
		AddAnimation("move S", "girl move S").
		AddAnimation("move W", "girl move W").
		AddAnimation("move E", "girl move E")
	hero.Play("move E")
	return hero
}

// HandleEvents -
func (o *Hero) HandleEvents() {
	dx := sdf.PressedInt(sdf.InputRight) - sdf.PressedInt(sdf.InputLeft)
	dy := sdf.PressedInt(sdf.InputDown) - sdf.PressedInt(sdf.InputUp)
	state := "idle"
	if dx != 0 {
		state = "move W"
		if dx > 0 {
			state = "move E"
		}

	} else if dy != 0 {
		state = "move N"
		if dy > 0 {
			state = "move S"
		}
	}
	rest, ok := o.tween.Process(sdf.FixedTime())
	_ = rest
	if ok {
		const len = 250
		hero.Play(state)
		xptr, yptr := hero.GetOffsetPtr()
		switch state {
		case "move N":
			o.tween.Reset(yptr, sdf.FixedTime(), len*time.Millisecond, *yptr-1, *yptr-16)
		case "move S":
			o.tween.Reset(yptr, sdf.FixedTime(), len*time.Millisecond, *yptr+1, *yptr+16)
		case "move W":
			o.tween.Reset(xptr, sdf.FixedTime(), len*time.Millisecond, *xptr-1, *xptr-16)
		case "move E":
			o.tween.Reset(xptr, sdf.FixedTime(), len*time.Millisecond, *xptr+1, *xptr+16)
		}
	}
}

func (o *game) HandleEvent(ev sdf.KeyboardEvent) {
	// switch ev := ev.(type) {
	// case *sdl.KeyboardEvent:
	// 	if ev.Type != sdl.KEYDOWN {
	// 		return
	// 	}
	// 	fmt.Printf("keysym %q\n", ev.Keysym.Sym)
	// case *sdl.TextInputEvent:
	// 	textInput := ""
	// 	slice := ev.Text[:]
	// 	for len(slice) > 0 {
	// 		r, size := utf8.DecodeRune(slice)
	// 		// fmt.Printf("%c %v\n", r, size)
	// 		if r == '\x00' {
	// 			break
	// 		}
	// 		textInput += string(r)
	// 		slice = slice[size:]
	// 	}
	// 	fmt.Printf("text input %q\n", textInput)
	// 	// textInput += string(slice)
	// }
	// e := sdf.Event{}
	// e.Align = 1
	// e.Type = 2
	// e.Mod = 3
	// e.Rune = rune(4)
	// fmt.Println(e)
	// x := int64(0)
	// x = e.BinaryKey()
	// fmt.Printf("%x\n", x)
	// fmt.Printf("%v, %v, %q\n", ev.Timestamp, ev.Mod, ev.Rune)
}

func (o *game) CleanUp() {
	fmt.Printf("cleaned up\n")
}

func (o *game) Render() {
	// if str, ok := sdf.Scanbuf(); ok {
	// 	fmt.Println(str)
	// }

	if sdf.Pressed(sdf.InputCancel) {
		sdf.Quit()
	}

	if sdf.Pressed(sdf.InputAccept) {
		if hero.Suspended() {
			hero.Continue()
		}
	}
	if sdf.Pressed(sdf.InputPause) {
		if !hero.Suspended() {
			hero.Suspend()
		}
	}
	hero.HandleEvents()

	sdf.Renderer().Clear()
	drawGrid()
	// tex.Copy(5, 5)
	// font.Print(0, 100, "Test String")

	hero.Copy(hero.x, hero.y)

	// spriteN.Copy(150, 40)
	// spriteS.Copy(150, 70)
	// spriteE.Copy(150, 100)
	anim0.Tile(sdf.FixedTime()).Copy(150, 100)

	a := sdf.JustPressedInt(sdf.InputCopy)
	b := sdf.PressedInt(sdf.InputCopy)
	msg := fmt.Sprintf("copy pressed once %v; down %v", a, b)
	font.Print(0, 10, msg)

	a = sdf.JustPressedInt(sdf.InputPaste)
	b = sdf.PressedInt(sdf.InputPaste)
	msg = fmt.Sprintf("paste pressed once %v; down %v", a, b)
	font.Print(0, 20, msg)

	a = sdf.JustPressedInt(sdf.InputDelete)
	b = sdf.PressedInt(sdf.InputDelete)
	msg = fmt.Sprintf("delete pressed once %v; down %v", a, b)
	font.Print(0, 30, msg)

	// fmt.Printf("%q\n", sdf.TextInput())
	// font.Print(0, 50, sdf.TextInput())
}

func drawGrid() {
	const (
		cellW = 16
		cellH = 16
	)
	c1 := color.RGBA{255, 255, 0, 127}
	c2 := color.RGBA{0, 255, 255, 127}
	r := sdf.Renderer()
	r.SetDrawBlendMode(sdl.BLENDMODE_BLEND)
	w, h := sdf.Size()
	x := 0
	for x < w {
		drawLine(r, x, 0, x, h-1, c1)
		x += cellW
		drawLine(r, x-1, 0, x-1, h-1, c2)
	}
	y := 0
	for y < h {
		drawLine(r, 0, y, w-1, y, c1)
		y += cellH
		drawLine(r, 0, y-1, w-1, y-1, c2)
	}
	drawLine(r, -1, -1, -1, -1, color.RGBA{0, 0, 0, 255})
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
