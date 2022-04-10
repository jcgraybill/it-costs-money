package main

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jcgraybill/it-costs-money/coin"
	"github.com/jcgraybill/it-costs-money/level"
	"github.com/jcgraybill/it-costs-money/player"
	"github.com/jcgraybill/it-costs-money/sys"
)

type Game struct {
	count  int
	level  level.Level
	player player.Player
	tiles  []*ebiten.Image
	coin   coin.Coin
}

var frameBuffer *ebiten.Image
var message string

func main() {
	frameBuffer = ebiten.NewImage(sys.ScreenWidth, sys.ScreenHeight)
	ebiten.SetWindowSize(sys.ScreenWidth, sys.ScreenHeight)
	ebiten.SetWindowTitle("it costs money to be alive")

	var g Game
	g.count = 0
	g.tiles = sys.LoadSpriteSheet("assets/1-tiles-city.png")
	g.tiles = append(g.tiles, sys.LoadSpriteSheet("assets/2-tiles-country.png")...)
	g.tiles = append(g.tiles, sys.LoadSpriteSheet("assets/3-objects-city.png")...)
	g.tiles = append(g.tiles, sys.LoadSpriteSheet("assets/4-objects-country.png")...)
	g.level = level.New(1, g.tiles)
	g.player = player.New()
	g.player.X, g.player.Y = g.level.StartPosition()
	g.coin = coin.New()

	if err := ebiten.RunGame(&g); err != nil {
		panic(err)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return sys.ScreenWidth, sys.ScreenHeight
}
