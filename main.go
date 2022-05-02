package main

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jcgraybill/it-costs-money/level"
	"github.com/jcgraybill/it-costs-money/player"
	"github.com/jcgraybill/it-costs-money/sys"
)

type Game struct {
	count  int
	level  level.Level
	player player.Player
}

var frameBuffer *ebiten.Image

var message string

func main() {
	frameBuffer = ebiten.NewImage(sys.ScreenWidth, sys.ScreenHeight)

	ebiten.SetWindowSize(sys.ScreenWidth, sys.ScreenHeight)
	ebiten.SetWindowTitle("it costs money to be alive")

	var g Game
	g.count = 0

	g.level = level.New(1)
	g.player = player.New()
	g.player.X, g.player.Y = g.level.StartPosition()
	sys.WriteMessage(540, 150, "run and jump with arrow keys\ncollect coins, but hurry up!\nyou lose coins over time\n(it costs money to be alive)", g.level.LevelBackgroundImage)

	if err := ebiten.RunGame(&g); err != nil {
		panic(err)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return sys.ScreenWidth, sys.ScreenHeight
}
