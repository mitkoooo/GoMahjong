package assets

import (
	"image"
	_ "image/png"
	"io/fs"
	"strings"

	"embed"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed *
var assets embed.FS

var TileSprites = mustLoadImages("tiles/*.png")
var TileBackSprite = mustLoadImage("tile_back.png")

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

func mustLoadImages(path string) map[string]*ebiten.Image {
	matches, err := fs.Glob(assets, path)
	if err != nil {
		panic(err)
	}

	// TODO return hashmap from name to *ebiten.Image

	images := make(map[string]*ebiten.Image)
	for _, match := range matches {

		_, k, _ := strings.Cut(match, "/")

		key, _, _ := strings.Cut(k, ".")

		images[key] = mustLoadImage(match)
	}

	return images
}