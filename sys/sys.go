package sys

import (
	"bytes"
	"image"
	"image/color"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

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

func LoadSpriteSheet(path string) []*ebiten.Image {
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

func Font() *font.Face {
	ttbytes, err := GameData("assets/Modak-Regular.ttf")
	if err == nil {
		tt, err := opentype.Parse(ttbytes)
		if err == nil {
			fontface, err := opentype.NewFace(tt, &opentype.FaceOptions{
				Size:    36,
				DPI:     72,
				Hinting: font.HintingFull,
			})
			if err == nil {
				return &fontface
			}
			panic(err)
		}
		panic(err)
	}
	panic(err)
}
