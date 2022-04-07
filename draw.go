package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func (g *Game) Draw(screen *ebiten.Image) {

	frameBuffer.Clear()
	parallax(level.bgImage1, -player.x, 4)
	parallax(level.bgImage2, -player.x, 3)
	parallax(level.bgImage3, -player.x, 2)

	viewPortOffset := player.x - (screenWidth/2 + frameWidth/2)
	levelViewFinder := image.Rect(viewPortOffset, 0, viewPortOffset+screenWidth, screenHeight)
	frameBuffer.DrawImage(level.levelBackgroundImage.SubImage(levelViewFinder).(*ebiten.Image), nil)
	frameBuffer.DrawImage(level.levelImage.SubImage(levelViewFinder).(*ebiten.Image), nil)

	runnerOP := &ebiten.DrawImageOptions{}

	if player.facingLeft {
		runnerOP.GeoM.Scale(-1, 1)
		runnerOP.GeoM.Translate(frameWidth, 0)
	}
	runnerOP.GeoM.Translate(-float64(frameWidth)/2, 0)
	runnerOP.GeoM.Translate(screenWidth/2, float64(player.y))
	i := (g.count / 5) % len(*player.slides)

	for _, actor := range actors {
		if actor.exists {
			if actor.x > viewPortOffset || actor.x < viewPortOffset+screenWidth/2-frameWidth {
				coinOp := &ebiten.DrawImageOptions{}
				coinOp.GeoM.Translate(-float64(viewPortOffset-actor.x), float64(actor.y))
				frameBuffer.DrawImage(coin.slides[(g.count/coin.animationSpeed)%coin.numSlides], coinOp)
			}
		}
	}

	frameBuffer.DrawImage((*player.slides)[i], runnerOP)
	frameBuffer.DrawImage(level.levelForegroundImage.SubImage(levelViewFinder).(*ebiten.Image), nil)
	frameBuffer.DrawImage(level.levelJunkImage.SubImage(levelViewFinder).(*ebiten.Image), nil)
	ebitenutil.DebugPrint(frameBuffer, message)
	screen.DrawImage(frameBuffer, nil)
}
