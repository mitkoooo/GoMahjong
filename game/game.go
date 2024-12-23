package game

import (
	_ "image/png"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
	input "github.com/quasilyte/ebitengine-input"
)

const (
	ActionDrawTile input.Action = iota
)

const (
	screenWidth, screenHeight  = 1280, 960
)


type Game struct {
	player *Player
	remainingTiles []*Tile

	inputSystem input.System
	timer *Timer
}

func (g *Game) Update() error {
	// Update the logical state
	g.inputSystem.Update()
	g.player.Update()
	Tick(g.timer)

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

	for _, deckT := range g.player.deck.currentTiles {

		deckT.Draw(screen)
	}

}

func NewGame() *Game {
	g := &Game{
		remainingTiles: make([]*Tile, 0),
	}

	GenerateTiles(g)
	ShuffleTiles(g)

	g.player = NewPlayer(g, true)

	g.inputSystem.Init(input.SystemConfig{
		DevicesEnabled: input.AnyDevice,
	})
	keymap := input.Keymap{
		ActionDrawTile:  {input.KeySpace},
	}

	g.timer = NewTimer()

	g.player.input = g.inputSystem.NewHandler(0, keymap)

	



	
	return g
}

func IsNoneRemainingTiles(g *Game) bool {
	return len(g.remainingTiles) == 0
}

func GenerateTiles(g *Game) {
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
	
				tile := NewTile(i % numberLimit, suit, position)
	
			g.remainingTiles = append(g.remainingTiles, tile)
	
			curFreeX += 44
	
			if curFreeX > screenWidth*0.95 {
				curFreeX = 0
				curFreeY += 60
		}
			}
	
		}
}

func ShuffleTiles(g *Game) {
		// Shuffle tiles

		for i := range g.remainingTiles {
	
			swapI := rand.IntN(144)
						
			g.remainingTiles[swapI].position, g.remainingTiles[i].position = g.remainingTiles[i].position, g.remainingTiles[swapI].position
				
			g.remainingTiles[swapI], g.remainingTiles[i] = g.remainingTiles[i], g.remainingTiles[swapI]
			}
}