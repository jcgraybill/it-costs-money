//go:build !deploy
// +build !deploy

package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func levelEditor() {
	// Live reload of level
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		go loadLevel(true)
		actors = loadActors("levels/level_0_actors.csv", true)
	}
	// skip ahead to next spawn point
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		spawnX, spawnY := level.levelImage.Bounds().Dx(), 0
		foundSpawn := false
		for _, actor := range actors {
			if actor.kind == "s" {
				if actor.x > player.x && actor.x < spawnX {
					spawnX = actor.x + frameWidth
					spawnY = actor.y
					foundSpawn = true
				}

			}
		}
		if foundSpawn {
			player.x = spawnX
			player.y = spawnY
			player.yVelocity = 0
		} else {
			goToStartPosition()
		}
	}
	xCellRune := ' '
	if xCell := (player.x-frameWidth/2)/(frameWidth*26) + 64; xCell > 64 {
		xCellRune = rune(xCell)
	}
	pos := fmt.Sprintf("%c%c:%d", xCellRune, rune(((player.x-frameWidth/2)/frameWidth)%26+65), player.y/frameWidth+1)

	message = message + fmt.Sprintf("\n[%s](r)eload (s)pawn\ntps %d fps %d", pos, int(ebiten.CurrentTPS()), int(ebiten.CurrentFPS()))

}
