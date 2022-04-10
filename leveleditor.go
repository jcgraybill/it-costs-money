//go:build !deploy
// +build !deploy

package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/jcgraybill/it-costs-money/level"
	"github.com/jcgraybill/it-costs-money/util"
)

func levelEditor(g *Game) string {
	// Live reload of level
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.level = level.New(0, g.tiles)
	}
	// skip ahead to next spawn point
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		spawnX, spawnY := g.level.LevelImage.Bounds().Dx(), 0
		foundSpawn := false
		for _, actor := range g.level.Actors {
			if actor.Kind == "s" {
				if actor.X > g.player.X && actor.X < spawnX {
					spawnX = actor.X + util.FrameWidth
					spawnY = actor.Y
					foundSpawn = true
				}

			}
		}
		if foundSpawn {
			g.player.X = spawnX
			g.player.Y = spawnY
			g.player.YVelocity = 0
		} else {
			g.player.X, g.player.Y = g.level.StartPosition()
		}
	}
	xCellRune := ' '
	if xCell := (g.player.X-util.FrameWidth/2)/(util.FrameWidth*26) + 64; xCell > 64 {
		xCellRune = rune(xCell)
	}
	pos := fmt.Sprintf("%c%c:%d", xCellRune, rune(((g.player.X-util.FrameWidth/2)/util.FrameWidth)%26+65), g.player.Y/util.FrameWidth+1)

	return fmt.Sprintf("\n[%s](r)eload (s)pawn\ntps %d fps %d", pos, int(ebiten.CurrentTPS()), int(ebiten.CurrentFPS()))

}
