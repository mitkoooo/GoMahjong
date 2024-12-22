package game

type Player struct {
	game *Game
	deck *Deck
	score int
}

func NewPlayer(g *Game) *Player {




	p := &Player{
		game: g,
		score: 0,
	}

	return p
}