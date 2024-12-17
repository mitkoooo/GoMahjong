package assets

import (
	"image"
	_ "image/png"

	"embed"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed *
var assets embed.FS

var TileSprite = mustLoadImage("deck_mahjong_light_0.png")

func mustLoadImage(name string) *ebiten.Image {
	f, err := assets.Open(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}