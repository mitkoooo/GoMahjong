package main

import (
	"image"
	_ "image/png"

	"github.com/mitkooo"

	"github.com/hajimehoshi/ebiten/v2"
)

var TileSprite = mustLoadImage("deck_mahjong_light_0.png")

func mustLoadImage(name string) *ebiten.Image {
	f, err := assets.Open(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}

type Game struct {}

func (g *Game) Update() error {
	// Update the logical state
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	// Return the game logical screen size
	// The screen is automatically scaled
	return 320, 240
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("GoMahjong")
	game := &Game{}
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}