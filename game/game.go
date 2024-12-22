package game

import (
	"fmt"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

// 640, 480
const (
	screenWidth, screenHeight  = 1280, 960
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

	// Initialize tile pool
	var curFreeX, curFreeY float64 = 0, 0

	

	for _, suit := range []Suit{bamboo, dot, thousand, dragon, wind, flower, season} {

		var numberLimit, numI int

		// determine number of the tiles to be generated

		switch suit {
		case bamboo, dot, thousand:
			numberLimit = 9
			numI = 4
		case dragon:
			numberLimit = 3
			numI = 4
		case wind:
			numberLimit = 4
			numI = 4
		case flower, season:
			numberLimit = 4
			numI = 1
		default:
			numberLimit = 0
		}



		for i := 0; i < numberLimit * numI; i++ {

			position := Vector{curFreeX, curFreeY}

			fmt.Println(position)

			// the problem is that i is tightly connected to the number on the asset name, so when you go above like 3 with dragons, it gets access error

			tile := NewTile(i % numberLimit, suit, position)

		g.remainingTiles = append(g.remainingTiles, tile)

		curFreeX += 44

		if curFreeX > screenWidth*0.95 {
			curFreeX = 0
			curFreeY += 60
	}
		}
	}

	
	return g
}

// Shuffle tiles

// for i := range g.remainingTiles {
	
// 	swapI := rand.IntN(144)

// 	g.remainingTiles[swapI], g.remainingTiles[i] = g.remainingTiles[i], g.remainingTiles[swapI]
// }

// switch suit {
// case bamboo, dot, thousand:
// 	numberLimit = 9 * 4
// 	numberLimit = 3 * 4
// case wind:
// 	numberLimit = 4 * 4
// case flower, season:
// 	numberLimit = 4
// }