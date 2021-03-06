package sys

import (
	"bytes"
	"image"
	"image/color"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	CoinSlides         = 6
	CoinAnimationSpeed = 10
)

var Tiles []*ebiten.Image
var CoinTiles []*ebiten.Image
var RunnerTiles []*ebiten.Image

func init() {
	Tiles = loadSpriteSheet("assets/1-tiles-city.png")
	Tiles = append(Tiles, loadSpriteSheet("assets/2-tiles-country.png")...)
	Tiles = append(Tiles, loadSpriteSheet("assets/3-objects-city.png")...)
	Tiles = append(Tiles, loadSpriteSheet("assets/4-objects-country.png")...)

	coinSprites := loadSpriteSheet("assets/coin.png")
	CoinTiles = coinSprites[1:7]

	RunnerTiles = loadSpriteSheet("assets/runner.png")

}

func LoadImage(path string) *ebiten.Image {
	imgBytes, err := GameData(path)
	if err == nil {
		img, _, err := image.Decode(bytes.NewReader(imgBytes))
		if err == nil {
			return ebiten.NewImageFromImage(img)
		}
		panic(err)
	}
	panic(err)
}

func loadSpriteSheet(path string) []*ebiten.Image {
	numberofSprites := 0
	spriteSheet := LoadImage(path)
	numberofSprites += spriteSheet.Bounds().Dx() / FrameWidth * spriteSheet.Bounds().Dy() / FrameHeight
	sprites := make([]*ebiten.Image, numberofSprites+2)
	i := 1
	sprites[0] = ebiten.NewImage(FrameWidth, FrameHeight)
	sprites[0].Fill(color.Black)
	for y := 0; y < spriteSheet.Bounds().Dy()/FrameHeight; y++ {
		for x := 0; x < spriteSheet.Bounds().Dx()/FrameWidth; x++ {
			sprites[i] = spriteSheet.SubImage(image.Rect(x*FrameWidth, y*FrameHeight, x*FrameWidth+FrameWidth, y*FrameHeight+FrameHeight)).(*ebiten.Image)
			i++
		}
	}
	sprites[i] = sprites[0]
	return sprites
}
