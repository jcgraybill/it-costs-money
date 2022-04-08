package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *Game) Update() error {
	g.count++

	player.slides = &player.idleFrames

	if player.x > level.levelImage.Bounds().Dx()-screenWidth/2+frameWidth/2-1 {
		message = fmt.Sprintf("Congratulations! You collected %d coins.\nPlease place them in the dumpster.", player.coins)
		return nil
	}

	if g.count%level.coinDecay == 0 && player.coins > 0 {
		player.coins--
	}

	// Detect collisions
	touchingGround := false
	if _, _, _, a := level.levelImage.At(player.x-frameWidth/2, player.y+frameHeight).RGBA(); a > 0 {
		touchingGround = true
	}

	touchingLeft := false
	if _, _, _, a := level.levelImage.At(player.x-frameWidth, player.y+frameHeight/2).RGBA(); a > 0 {
		touchingLeft = true
	}

	leftAdjacent := false
	if _, _, _, a := level.levelImage.At(player.x-frameWidth-moveSpeed, player.y+frameHeight/2).RGBA(); a > 0 {
		leftAdjacent = true
	}

	touchingRight := false
	if _, _, _, a := level.levelImage.At(player.x, player.y+frameHeight/2).RGBA(); a > 0 {
		touchingRight = true
	}

	rightAdjacent := false
	if _, _, _, a := level.levelImage.At(player.x+moveSpeed, player.y+frameHeight/2).RGBA(); a > 0 {
		rightAdjacent = true
	}

	// Jumping

	if touchingGround {
		player.yVelocity = 0
		player.wileECoytoe = wileECoyoteFrames
		if ebiten.IsKeyPressed(ebiten.KeySpace) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
			if player.timeSinceLastJump+jumpRecovery < g.count {
				if _, _, _, a := level.levelImage.At(player.x-frameWidth/2, player.y-jumpHeight).RGBA(); a == 0 {
					player.yVelocity = -jumpHeight
					player.y += int(player.yVelocity)
					player.timeSinceLastJump = g.count
					player.wileECoytoe = 0
				}
			}
		}
	} else {
		player.y += int(player.yVelocity)
		player.yVelocity += gravity
		if player.wileECoytoe == 0 {
			player.slides = &player.fallFrames
		} else {
			player.wileECoytoe -= 1
		}
	}

	touchingTop := false
	if _, _, _, a := level.levelImage.At(player.x-frameWidth/2, player.y).RGBA(); a > 0 {
		touchingTop = true
	}

	// Running
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		player.facingLeft = false
		if player.x < level.levelImage.Bounds().Dx()-screenWidth/2+frameWidth/2-1 && !touchingRight && !rightAdjacent {
			player.slides = &player.runFrames
			player.x += moveSpeed
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		player.facingLeft = true
		if player.x > screenWidth/2+frameWidth/2 && !touchingLeft && !leftAdjacent {
			player.slides = &player.runFrames
			player.x -= moveSpeed
		}
	}

	// Correct if overlapping with terrain

	if touchingTop && !touchingGround {
		player.y += 1
		player.yVelocity = 0
	}

	if _, _, _, a := level.levelImage.At(player.x-frameWidth/2, player.y+frameHeight-1).RGBA(); a != 0 {
		for a != 0 {
			player.y -= 1
			_, _, _, a = level.levelImage.At(player.x-frameWidth/2, player.y+frameHeight-1).RGBA()
		}
	}

	if touchingLeft && !touchingRight {
		player.x += moveSpeed
	}

	if touchingRight && !touchingLeft {
		player.x -= moveSpeed
	}

	// Pick up coins
	for _, actor := range actors {
		if actor.exists && actor.kind == "c" && actor.x+frameWidth/2 > player.x-frameWidth && actor.x+frameWidth/2 < player.x && actor.y+frameHeight/2 > player.y && actor.y+frameHeight/2 < player.y+frameHeight {
			actor.exists = false
			player.coins++
			coin.audioPlayers[g.count%5].Rewind()
			coin.audioPlayers[g.count%5].Play()
		}
	}

	// Fall in holes
	if player.y > screenHeight*4 {
		spawnX, spawnY := 0, 0
		for _, actor := range actors {
			if actor.kind == "s" {
				if actor.x > spawnX && actor.x < player.x {
					spawnX = actor.x + frameWidth
					spawnY = actor.y
				}

			}
		}
		player.x = spawnX
		player.y = spawnY
		player.yVelocity = 0
	}

	message = fmt.Sprintf("Gather coins and bring them to the green chest.\nIt costs money to be alive!\nYour coins: %d", player.coins)

	if editMode {
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
			}
		}

		message = message + fmt.Sprintf("\nEDIT MODE: (r)eload (s)pawn\ntps %d fps %d", int(ebiten.CurrentTPS()), int(ebiten.CurrentFPS()))
	}

	return nil
}
