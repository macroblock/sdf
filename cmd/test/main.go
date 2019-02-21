package main

import (
	"fmt"

	"github.com/macroblock/sdf/pkg/fonts/pixfm5x9normal"
	"github.com/macroblock/sdf/pkg/sdf"
)

var (
	tex                       *sdf.Texture
	font                      *sdf.PixelFont
	tileSheet                 *sdf.TileSheet
	sprite0, sprite1, sprite2 *sdf.Sprite
)

type game struct {
}

func (o *game) Init() {
	fmt.Printf("initialized\n")

	tex = sdf.LoadTexture("../../assets/ff6sheet.png")
	fmt.Println("tex ", tex)

	font = sdf.CreatePixelFont(pixfm5x9normal.Font)
	font.SetScale(2)

	sdf.SetScale(3, 3)

	tileSheet = sdf.LoadTileSheet(-8, -6, "../../assets/ff6sheet.png")
	tileSheet.InitTile("front move 0", 1, 0, nil, 0)
	tileSheet.InitTile("front move 1", 0, 0, nil, 0)
	tileSheet.InitTile("front move 2", 2, 0, nil, 0)

	tileSheet.InitTile("left move 0", 7, 0, nil, 0)
	tileSheet.InitTile("left move 1", 6, 0, nil, 0)
	tileSheet.InitTile("left move 2", 0, 1, nil, 0)

	tileSheet.InitTile("right move 0", 7, 0, nil, sdf.FlipHorizontal)
	tileSheet.InitTile("right move 1", 6, 0, nil, sdf.FlipHorizontal)
	tileSheet.InitTile("right move 2", 0, 1, nil, sdf.FlipHorizontal)

	sprite0 = tileSheet.InitSprite("ff6 front move", 1.0, "front move 0").
		AddKeyframe(0.25, "front move 1").
		AddKeyframe(0.50, "front move 0").
		AddKeyframe(0.75, "front move 2")
	sprite1 = tileSheet.InitSprite("ff6 left move", 1.0, "left move 0").
		AddKeyframe(0.25, "left move 1").
		AddKeyframe(0.50, "left move 0").
		AddKeyframe(0.75, "left move 2").
		SetSpeed(1.5)
	sprite2 = tileSheet.InitSprite("ff6 right move", 1.0, "right move 0").
		AddKeyframe(0.25, "right move 1").
		AddKeyframe(0.50, "right move 0").
		AddKeyframe(0.75, "right move 2").
		SetSpeed(2)
}

func (o *game) CleanUp() {
	fmt.Printf("cleaned up\n")
}

func (o *game) Render() {
	delta := sdf.DeltaRender()

	tex.Copy(5, 5)
	font.Print(0, 100, "Test String")
	tile := tileSheet.Tile("front move 0")
	tile.Copy(150, 100, -1)

	sprite0.Copy(150, 10, delta)
	sprite1.Copy(150, 40, delta)
	sprite2.Copy(150, 70, delta)
}

func main() {
	err := sdf.Run(&game{})

	if err != nil {
		fmt.Println(err)
	}
}
