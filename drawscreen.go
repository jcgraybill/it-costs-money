package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/jcgraybill/it-costs-money/util"
)

func (g *Game) Draw(screen *ebiten.Image) {

	// TODO objects should draw themselves - add a draw.go to coin, level, player
	// Each returns an appropriate image

	frameBuffer.Clear()
	parallax(g.level.BgImage1, -g.player.X, 4)
	parallax(g.level.BgImage2, -g.player.X, 3)
	parallax(g.level.BgImage3, -g.player.X, 2)

	viewPortOffset := g.player.X - (util.ScreenWidth/2 + util.FrameWidth/2)
	levelViewFinder := image.Rect(viewPortOffset, 0, viewPortOffset+util.ScreenWidth, util.ScreenHeight)
	frameBuffer.DrawImage(g.level.LevelBackgroundImage.SubImage(levelViewFinder).(*ebiten.Image), nil)
	frameBuffer.DrawImage(g.level.LevelImage.SubImage(levelViewFinder).(*ebiten.Image), nil)

	runnerOP := &ebiten.DrawImageOptions{}

	if g.player.FacingLeft {
		runnerOP.GeoM.Scale(-1, 1)
		runnerOP.GeoM.Translate(util.FrameWidth, 0)
	}
	runnerOP.GeoM.Translate(-float64(util.FrameWidth)/2, 0)
	runnerOP.GeoM.Translate(util.ScreenWidth/2, float64(g.player.Y))
	i := (g.count / 5) % len(*g.player.Slides)

	for _, actor := range g.level.Actors {
		if actor.Exists {
			if actor.Kind == "c" {
				if actor.X > viewPortOffset || actor.X < viewPortOffset+util.ScreenWidth/2-util.FrameWidth {
					coinOp := &ebiten.DrawImageOptions{}
					coinOp.GeoM.Translate(-float64(viewPortOffset-actor.X), float64(actor.Y))
					frameBuffer.DrawImage(g.coin.Slides[(g.count/g.coin.AnimationSpeed)%g.coin.NumSlides], coinOp)
				}
			}

		}
	}

	frameBuffer.DrawImage((*g.player.Slides)[i], runnerOP)
	frameBuffer.DrawImage(g.level.LevelForegroundImage.SubImage(levelViewFinder).(*ebiten.Image), nil)
	ebitenutil.DebugPrint(frameBuffer, message)
	screen.DrawImage(frameBuffer, nil)
}

func parallax(image *ebiten.Image, offset int, speed int) {
	panelWidth := image.Bounds().Dx()
	position := (offset / speed) % panelWidth
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(position), 0)
	frameBuffer.DrawImage(image, op)

	for i := 0; position+i*panelWidth+panelWidth < util.ScreenWidth; i++ {
		op.GeoM.Translate(float64(panelWidth), 0)
		frameBuffer.DrawImage(image, op)
	}
}
