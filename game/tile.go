package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mitkoooo/GoMahjong/assets"
)

type Suit int

const (
    bamboo Suit = iota + 1
    dot
    thousand
)

type Tile struct {
	name string
	cardinality int
	suit Suit
	sprite *ebiten.Image
}

func NewTile(cardinality int, suit Suit) *Tile {

	sprite :=

	t := &Tile{
		cardinality: cardinality,
		suit: suit,
	}

	return t
}