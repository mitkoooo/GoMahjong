package game

import (
	_ "image/png"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 1200
	screenHeight = 1000
)


type Game struct {
	player *Player
	remainingTiles []*Tile
}

func (g *Game) Update() error {
	// Update the logical state
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	// Return the game logical screen size
	// The screen is automatically scaled
	return screenWidth, screenHeight
}

func (g *Game) Draw(screen *ebiten.Image) {

	for _, t := range g.remainingTiles {

		t.Draw(screen)
	}

}

func NewGame() *Game {
	g := &Game{
		remainingTiles: make([]*Tile, 0),
	}

	var curFreeX, curFreeY float64 = 0, 0

	for _, suit := range []Suit{bamboo, dot, thousand, dragon, wind, flower, season} {

		var numberLimit int

		if suit == bamboo || suit == dot || suit == thousand {
			numberLimit = 10
		} else if suit == dragon {
			numberLimit = 4
		} else {
			numberLimit = 5
		}

		for i := 1; i < numberLimit; i++ {

			position := Vector{curFreeX, curFreeY}

			tile := NewTile(i, suit, position)

		g.remainingTiles = append(g.remainingTiles, tile)
		
		curFreeX += 60

		if curFreeX > screenWidth - 80 {
				curFreeX = math.Mod(curFreeX, (screenWidth - 80))
				curFreeY += 70
		}

		}

	}
	return g
}
