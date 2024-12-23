package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mitkoooo/GoMahjong/assets"
)

type Suit int

const (
	bamboo Suit = iota + 1
	dot
	thousand
	dragon
	wind
	flower
	season
)

type Tile struct {
	position Vector
	cardinality int
	suit Suit
	sprite *ebiten.Image
	isRevealed bool
}

func (t *Tile) Draw(screen *ebiten.Image) {

	var spriteToDraw *ebiten.Image
	op := &ebiten.DrawImageOptions{}
	
	op.GeoM.Translate(t.position.X, t.position.Y)

	if t.isRevealed {
		op.GeoM.Scale(1.7, 1.7)
		spriteToDraw = t.sprite
	} else {
		spriteToDraw = assets.TileBackSprite
	}

	screen.DrawImage(spriteToDraw, op)
}


func NewTile(cardinality int, suit Suit, position Vector) *Tile {

	var suit_name string

	switch suit {

	case bamboo:
		suit_name = "bamboo"

	case dot:
		suit_name = "dot"

	case thousand:
		suit_name = "thousand"

	case dragon:
		suit_name = "dragon"

	case wind:
		suit_name = "wind"

	case flower:
		suit_name = "flower"

	case season:
		suit_name = "season"
			
	default:
		suit_name = ""
	}

	key := fmt.Sprintf("%s_%d", suit_name, cardinality)


	sprite := assets.TileSprites[key]

	t := &Tile{
		cardinality: cardinality,
		suit: suit,
		sprite: sprite,
		position: position,
	}

	return t
}