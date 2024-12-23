package game

import (
	"fmt"
	"time"

	input "github.com/quasilyte/ebitengine-input"
)

type Player struct {
	game *Game
	deck *Deck
	score int
	isDealer bool

	input *input.Handler
}

func NewPlayer(g *Game) *Player {

	p := &Player{
		game: g,
		score: 0,
		deck: NewDeck(make([]*Tile, 0)),
	}

	return p
}

var curFreeX, curFreeY float64 = screenWidth*0.05, screenHeight*0.90

func DrawTile(this *Player) {

	var tileLimit int

	if this.isDealer {
		tileLimit = 14
	} else {
		tileLimit = 13
	}

	fmt.Println(len(this.game.remainingTiles))

	if IsNoneRemainingTiles(this.game)|| len(this.deck.currentTiles) + len(this.deck.revealedTiles) > tileLimit{
		return 
	}

	newTile := this.game.remainingTiles[0]

	newTile.position = Vector{curFreeX, curFreeY}

	newDeck := append(this.deck.currentTiles, this.game.remainingTiles[0])

	curFreeX += 44
	
			if curFreeX > screenWidth*0.95 {
				curFreeX = 0
				curFreeY += 60
		}

	this.deck.currentTiles = newDeck

	this.game.remainingTiles = this.game.remainingTiles[1:]

}


func (p *Player) Update() {
	if p.input.ActionIsPressed(ActionDrawTile) {
		time.Sleep(2000)
		fmt.Println("Drawn a tile")
		DrawTile(p)
	}
}