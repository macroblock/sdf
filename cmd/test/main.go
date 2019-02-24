package main

import (
	"fmt"

	"github.com/macroblock/sdf/pkg/fonts/pixfm5x9normal"
	"github.com/macroblock/sdf/pkg/sdf"
)

var (
	tex       *sdf.Texture
	font      *sdf.PixelFont
	tileSheet *sdf.TileSheet
	// sprite0, sprite1, sprite2                *sdf.Sprite
	anim0 *sdf.Animation
	hero  *sdf.GameObject
)

// const (
// 	stateN int = iota
// 	stateS
// 	stateW
// 	stateE
// 	stateIdle
// 	maxState
// )

type game struct {
}

func (o *game) Init() {
	fmt.Printf("initialized\n")

	tex = sdf.LoadTexture("../../assets/testsheet.png")

	font = sdf.CreatePixelFont(pixfm5x9normal.Font)
	font.SetScale(2)

	sdf.SetScale(3, 3)

	tileSheet = sdf.LoadTileSheet("test tile sheet", -8, -6, "../../assets/testsheet.png")

	sdf.CreateTile("idle 0", 1, 0, nil, 0)
	sdf.CreateTile("idle 1", 3, 2, nil, 0)

	sdf.CreateTile("n move 0", 4, 0, nil, 0)
	sdf.CreateTile("n move 1", 3, 0, nil, 0)
	sdf.CreateTile("n move 2", 5, 0, nil, 0)

	sdf.CreateTile("s move 0", 1, 0, nil, 0)
	sdf.CreateTile("s move 1", 0, 0, nil, 0)
	sdf.CreateTile("s move 2", 2, 0, nil, 0)

	sdf.CreateTile("w move 0", 7, 0, nil, 0)
	sdf.CreateTile("w move 1", 6, 0, nil, 0)
	sdf.CreateTile("w move 2", 0, 1, nil, 0)

	sdf.CreateTile("e move 0", 7, 0, nil, sdf.FlipHorizontal)
	sdf.CreateTile("e move 1", 6, 0, nil, sdf.FlipHorizontal)
	sdf.CreateTile("e move 2", 0, 1, nil, sdf.FlipHorizontal)

	anim0 = sdf.CreateAnimation("test move").Sequence("n move 0", "n move 1", "n move 0", "n move 2").StretchTo(1.0)
	sdf.CreateAnimation("hero idle").Sequence("idle 0", "idle 1").StretchTo(1.0)
	sdf.CreateAnimation("hero move N").Sequence("n move 0", "n move 1", "n move 0", "n move 2").StretchTo(1.0)
	sdf.CreateAnimation("hero move S").Sequence("s move 0", "s move 1", "s move 0", "s move 2").StretchTo(1.0)
	sdf.CreateAnimation("hero move W").Sequence("w move 0", "w move 1", "w move 0", "w move 2").StretchTo(1.0)
	sdf.CreateAnimation("hero move E").Sequence("e move 0", "e move 1", "e move 0", "e move 2").StretchTo(1.0)
	hero = sdf.NewGameObject("hero").
		AddAnimation("idle", "hero idle").
		AddAnimation("move N", "hero move N").
		AddAnimation("move S", "hero move S").
		AddAnimation("move W", "hero move W").
		AddAnimation("move E", "hero move E")
	hero.Play("move E")
}

func (o *game) HandleEvents() {
	// int start = SDL_GetTicks();
	//     gameLoop->update();
	//     int time = SDL_GetTicks() - start;
	//     if (time < 0) continue; // if time is negative, the time probably overflew, so continue asap

	//     int sleepTime = gameLoop->millisecondsForFrame - time;
	//     if (sleepTime > 0)
	//     {
	//         SDL_Delay(sleepTime);
	// 	}
}

func (o *game) CleanUp() {
	fmt.Printf("cleaned up\n")
}

func (o *game) Render() {
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
	hero.Play(state)

	sdf.Renderer().Clear()
	tex.Copy(5, 5)
	font.Print(0, 100, "Test String")

	hero.Copy(150, 10)

	// spriteN.Copy(150, 40)
	// spriteS.Copy(150, 70)
	// spriteE.Copy(150, 100)
	anim0.Tile(sdf.FixedTime()).Copy(150, 100)
}

func main() {
	err := sdf.Run(&game{})

	if err != nil {
		fmt.Println(err)
	}
}
