package main

import (
	"bytes"
	"fmt"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/jcgraybill/it-costs-money/level"
	"github.com/jcgraybill/it-costs-money/player"
	"github.com/jcgraybill/it-costs-money/sys"
	"golang.org/x/image/font"
)

type Game struct {
	count        int
	level        level.Level
	player       player.Player
	tiles        []*ebiten.Image
	ttf          *font.Face
	AudioPlayers [5]*audio.Player
}

var frameBuffer *ebiten.Image
var audioContext *audio.Context
var message string

func main() {
	frameBuffer = ebiten.NewImage(sys.ScreenWidth, sys.ScreenHeight)
	audioContext = audio.NewContext(24000)

	ebiten.SetWindowSize(sys.ScreenWidth, sys.ScreenHeight)
	ebiten.SetWindowTitle("it costs money to be alive")

	var g Game
	g.count = 0
	g.ttf = sys.Font()
	g.tiles = sys.LoadSpriteSheet("assets/1-tiles-city.png")
	g.tiles = append(g.tiles, sys.LoadSpriteSheet("assets/2-tiles-country.png")...)
	g.tiles = append(g.tiles, sys.LoadSpriteSheet("assets/3-objects-city.png")...)
	g.tiles = append(g.tiles, sys.LoadSpriteSheet("assets/4-objects-country.png")...)
	g.level = level.New(1, g.tiles, audioContext)
	g.player = player.New()
	g.player.X, g.player.Y = g.level.StartPosition()
	sys.WriteMessage(540, 150, "run and jump with arrow keys\ncollect coins, but hurry up!\nyou lose coins over time\n(it costs money to be alive)", g.level.LevelBackgroundImage, g.ttf)

	for i := 0; i < 5; i++ {
		audioBytes, err := sys.GameData(fmt.Sprintf("assets/Coins_Grab_0%d.ogg", i))
		if err == nil {
			d, err := vorbis.Decode(audioContext, bytes.NewReader(audioBytes))
			if err == nil {
				g.AudioPlayers[i], err = audioContext.NewPlayer(d)
				if err != nil {
					panic(err)
				}
			} else {
				panic(err)
			}
		} else {
			panic(err)
		}
	}

	if err := ebiten.RunGame(&g); err != nil {
		panic(err)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return sys.ScreenWidth, sys.ScreenHeight
}
