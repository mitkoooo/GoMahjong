package game

import (
	"fmt"
	_ "image/png"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
	input "github.com/quasilyte/ebitengine-input"
)

const (
	ActionDrawTile input.Action = iota
)

const (
	screenWidth, screenHeight  = 1500, 1500
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
	InitializeTilesOnScreen(g)

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
	
				position := Vector{0, 0}
				
				tile := NewTile(i % numberLimit, suit, position)
	
			g.remainingTiles = append(g.remainingTiles, tile)

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

func InitializeTilesOnScreen(g *Game) {
	// Form a square
	// Form a wall out of 18 tiles
	// Switch directions
	// Repeat twice
	// 18 * 2 * 4 = 144

	var curFreeX, curFreeY float64 = screenWidth*0.8, screenHeight*0.6

	var buildDirection = &curFreeX

	// Margin by which to place tiles
	deltaPosition := -44

	// At the end of the wall, add some space
	// positionShift := 0

	// Build two level wall
	nextOnTop := false

	// Flip tile sprite
	hasToFlip := false

	for i, t := range g.remainingTiles {
		///////////////////////////////////

		// This part is concerned with initializing a tile, no logic of determining where exactly
		if hasToFlip {
			Flip(t)
		}

		// if next is not on top, move position
		if !nextOnTop {
			*buildDirection += float64(deltaPosition)
		}
		
		t.position = Vector{curFreeX, curFreeY}

		fmt.Println(t.position)

		//////////////////////////////////

		// Determine position for the next tile

		// If end of the wall, switch direction, add margin and flip tiles
		if i != 0 && (i + 1) % 36 == 0 {

			hasToFlip = !hasToFlip

			if deltaPosition < 0 {
			*buildDirection -= 60
			} else {
			*buildDirection += 44
			}

			if buildDirection == &curFreeX {
				buildDirection = &curFreeY
			} else {
				buildDirection = &curFreeX
			}

			// *buildDirection += float64(positionShift)

			if i + 1 == 72 {
				deltaPosition *= -1
				}

			if deltaPosition < 0 {
			*buildDirection += 104
			} else {
				*buildDirection -= 88
			}
		}




		// alternate between on top and not on top
		nextOnTop = !nextOnTop

	}
}