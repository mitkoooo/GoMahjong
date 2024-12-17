package game

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mitkoooo/GoMahjong/assets"
)

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
	screen.DrawImage(assets.TileSprite, nil)
}

func NewGame() *Game {
	g := &Game{}

	return g
}
