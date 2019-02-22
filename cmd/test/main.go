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
	spriteN, spriteS, spriteW, spriteE, sprite0 sdf.IElem
	movement                                    []sdf.IElem
)

const (
	stateN int = iota
	stateS
	stateW
	stateE
	stateIdle
	maxState
)

type game struct {
}

func (o *game) Init() {
	fmt.Printf("initialized\n")

	tex = sdf.LoadTexture("../../assets/testsheet.png")
	fmt.Println("tex ", tex)

	font = sdf.CreatePixelFont(pixfm5x9normal.Font)
	font.SetScale(2)

	sdf.SetScale(3, 3)

	tileSheet = sdf.LoadTileSheet(-8, -6, "../../assets/testsheet.png")
	sprite0 = tileSheet.InitTile("idle", 2, 5, nil, 0)

	tileSheet.InitTile("n move 0", 4, 0, nil, 0)
	tileSheet.InitTile("n move 1", 3, 0, nil, 0)
	tileSheet.InitTile("n move 2", 5, 0, nil, 0)

	tileSheet.InitTile("s move 0", 1, 0, nil, 0)
	tileSheet.InitTile("s move 1", 0, 0, nil, 0)
	tileSheet.InitTile("s move 2", 2, 0, nil, 0)

	tileSheet.InitTile("w move 0", 7, 0, nil, 0)
	tileSheet.InitTile("w move 1", 6, 0, nil, 0)
	tileSheet.InitTile("w move 2", 0, 1, nil, 0)

	tileSheet.InitTile("e move 0", 7, 0, nil, sdf.FlipHorizontal)
	tileSheet.InitTile("e move 1", 6, 0, nil, sdf.FlipHorizontal)
	tileSheet.InitTile("e move 2", 0, 1, nil, sdf.FlipHorizontal)

	spriteN = tileSheet.InitSprite("ff6 north move", 1.0, "n move 0").
		AddKeyframe(0.25, "n move 1").
		AddKeyframe(0.50, "n move 0").
		AddKeyframe(0.75, "n move 2")
	spriteS = tileSheet.InitSprite("ff6 front move", 1.0, "s move 0").
		AddKeyframe(0.25, "s move 1").
		AddKeyframe(0.50, "s move 0").
		AddKeyframe(0.75, "s move 2")
	spriteW = tileSheet.InitSprite("ff6 left move", 1.0, "w move 0").
		AddKeyframe(0.25, "w move 1").
		AddKeyframe(0.50, "w move 0").
		AddKeyframe(0.75, "w move 2").
		SetSpeed(1.5)
	spriteE = tileSheet.InitSprite("ff6 right move", 1.0, "e move 0").
		AddKeyframe(0.25, "e move 1").
		AddKeyframe(0.50, "e move 0").
		AddKeyframe(0.75, "e move 2").
		SetSpeed(2)

	movement = make([]sdf.IElem, maxState)
	movement[stateN] = spriteN
	movement[stateS] = spriteS
	movement[stateW] = spriteW
	movement[stateE] = spriteE
	movement[stateIdle] = sprite0
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
		i := spriteE.(*sdf.Sprite)
		if i.Suspended() {
			i.Run()
		}
	}
	if sdf.Pressed(sdf.InputPause) {
		i := spriteE.(*sdf.Sprite)
		if !i.Suspended() {
			i.Suspend()
		}
	}

	// tile := tileSheet.Tile("front move 0")
	// tile.Copy(150, 100, -1)

	dx := sdf.PressedInt(sdf.InputRight) - sdf.PressedInt(sdf.InputLeft)
	dy := sdf.PressedInt(sdf.InputDown) - sdf.PressedInt(sdf.InputUp)
	state := stateIdle
	if dx != 0 {
		state = stateW
		if dx > 0 {
			state = stateE
		}

	} else if dy != 0 {
		state = stateN
		if dy > 0 {
			state = stateS
		}
	}

	// delta := sdf.DeltaRender()
	// _ = delta
	// time := sdf.Time()
	// _ = time
	// delta = time
	// fmt.Println("delta: ", delta, " fps: ", float64(time.Second)/float64(delta))
	// upd := false
	// upd = movement[state].Update(delta) || upd
	// upd = spriteN.Update(delta) || upd
	// upd = spriteS.Update(delta) || upd
	// upd = spriteE.Update(delta) || upd

	// if !upd {
	// return
	// }
	sdf.Renderer().Clear()
	tex.Copy(5, 5)
	font.Print(0, 100, "Test String")

	movement[state].Copy(150, 10)

	spriteN.Copy(150, 40)
	spriteS.Copy(150, 70)
	spriteE.Copy(150, 100)
}

func main() {
	err := sdf.Run(&game{})

	if err != nil {
		fmt.Println(err)
	}
}
