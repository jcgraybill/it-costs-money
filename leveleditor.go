//go:build !deploy
// +build !deploy

package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/jcgraybill/it-costs-money/level"
	"github.com/jcgraybill/it-costs-money/sys"
)

func levelEditor(g *Game) string {
	// Live reload of level
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.level = level.New(g.level.LevelNumber)
	}

	// skip ahead to next spawn point
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		g.player.X, g.player.Y = g.level.NextSpawn(g.player.X, g.player.Y)
		g.player.YVelocity = 0
	}

	// display current coordinates of player
	xCellRune := ' '
	if xCell := (g.player.X-sys.FrameWidth/2)/(sys.FrameWidth*26) + 64; xCell > 64 {
		xCellRune = rune(xCell)
	}
	pos := fmt.Sprintf("%c%c:%d", xCellRune, rune(((g.player.X-sys.FrameWidth/2)/sys.FrameWidth)%26+65), g.player.Y/sys.FrameWidth+1)

	return fmt.Sprintf("level edit mode\n[%s](r)eload (s)pawn\ntps %d fps %d", pos, int(ebiten.CurrentTPS()), int(ebiten.CurrentFPS()))

}
