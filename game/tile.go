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

	// 0 to 8 bamboo
	// 9 to 17 dot
	// 18 to 27 thousand

	var index int

	if suit == dot {

	}

	sprite := assets.TileSprites[0]

	t := &Tile{
		cardinality: cardinality,
		suit: suit,
		sprite: sprite,
	}

	return t
}