package game

type Deck struct {
	currentTiles []*Tile
	revealedTiles []*Tile
}

func NewDeck(newTiles []*Tile) *Deck {
	d := &Deck{
		currentTiles: newTiles,
		revealedTiles: []*Tile{},
	}

	return d
}



