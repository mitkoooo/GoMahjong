package main

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/mitkoooo/GoMahjong/game"
)

func main() {
	g := game.NewGame()
	err := ebiten.RunGame(g)
	ebiten.SetWindowSize(2800, 2800)
	if err != nil {
		panic(err)
	}
}