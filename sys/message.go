package sys

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

func WriteMessage(x, y int, message string, i *ebiten.Image, ttf *font.Face) {
	text.Draw(i, message, *ttf, x+2, y+2, color.RGBA{0x00, 0x00, 0x00, 0xff})
	text.Draw(i, message, *ttf, x, y, color.RGBA{0xd4, 0xaf, 0x47, 0xff})
}
