package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	_ "github.com/jcgraybill/it-costs-money/player"
	"github.com/jcgraybill/it-costs-money/util"
)

func (g *Game) Update() error {
	g.count++

	g.player.Slides = &g.player.IdleFrames

	//TODO if player is mid-jump, will get stuck hovering in the air
	if g.player.X > g.level.LevelImage.Bounds().Dx()-util.ScreenWidth/2+util.FrameWidth/2-1 {
		message = fmt.Sprintf("Congratulations! You collected %d coins.\nPlease place them in the dumpster.", g.player.Coins)
		return nil
	}

	if g.count%g.level.CoinDecay == 0 && g.player.Coins > 0 {
		g.player.Coins--
	}

	// Detect collisions
	touchingGround := false
	if _, _, _, a := g.level.LevelImage.At(g.player.X-util.FrameWidth/2, g.player.Y+util.FrameHeight).RGBA(); a > 0 {
		touchingGround = true
	}

	touchingLeft := false
	if _, _, _, a := g.level.LevelImage.At(g.player.X-util.FrameWidth, g.player.Y+util.FrameHeight/2).RGBA(); a > 0 {
		touchingLeft = true
	}

	leftAdjacent := false
	if _, _, _, a := g.level.LevelImage.At(g.player.X-util.FrameWidth-g.level.MoveSpeed, g.player.Y+util.FrameHeight/2).RGBA(); a > 0 {
		leftAdjacent = true
	}

	touchingRight := false
	if _, _, _, a := g.level.LevelImage.At(g.player.X, g.player.Y+util.FrameHeight/2).RGBA(); a > 0 {
		touchingRight = true
	}

	rightAdjacent := false
	if _, _, _, a := g.level.LevelImage.At(g.player.X+g.level.MoveSpeed, g.player.Y+util.FrameHeight/2).RGBA(); a > 0 {
		rightAdjacent = true
	}

	// Jumping

	if touchingGround {
		g.player.YVelocity = 0

		g.player.ResetWileECoyote()

		if ebiten.IsKeyPressed(ebiten.KeySpace) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
			if g.player.TimeSinceLastJump+g.player.JumpRecovery < g.count {
				if _, _, _, a := g.level.LevelImage.At(g.player.X-util.FrameWidth/2, g.player.Y-g.level.JumpHeight).RGBA(); a == 0 {
					g.player.YVelocity = -float64(g.level.JumpHeight)
					g.player.Y += int(g.player.YVelocity)
					g.player.TimeSinceLastJump = g.count
					g.player.WileECoyote = 0
				}
			}
		}
	} else {
		g.player.Y += int(g.player.YVelocity)
		g.player.YVelocity += g.level.Gravity
		if g.player.WileECoyote == 0 {
			g.player.Slides = &g.player.FallFrames
		} else {
			g.player.WileECoyote -= 1
		}
	}

	touchingTop := false
	if _, _, _, a := g.level.LevelImage.At(g.player.X-util.FrameWidth/2, g.player.Y).RGBA(); a > 0 {
		touchingTop = true
	}

	// Running
	// TODO allow WASD
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.player.FacingLeft = false
		if g.player.X < g.level.LevelImage.Bounds().Dx()-util.ScreenWidth/2+util.FrameWidth/2-1 && !touchingRight && !rightAdjacent {
			g.player.Slides = &g.player.RunFrames
			g.player.X += g.level.MoveSpeed
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.player.FacingLeft = true
		if g.player.X > util.ScreenWidth/2+util.FrameWidth/2 && !touchingLeft && !leftAdjacent {
			g.player.Slides = &g.player.RunFrames
			g.player.X -= g.level.MoveSpeed
		}
	}

	// Correct if overlapping with terrain

	if touchingTop && !touchingGround {
		g.player.Y += 1
		g.player.YVelocity = 0
	}

	if _, _, _, a := g.level.LevelImage.At(g.player.X-util.FrameWidth/2, g.player.Y+util.FrameHeight-1).RGBA(); a != 0 {
		for a != 0 {
			g.player.Y -= 1
			_, _, _, a = g.level.LevelImage.At(g.player.X-util.FrameWidth/2, g.player.Y+util.FrameHeight-1).RGBA()
		}
	}

	if touchingLeft && !touchingRight {
		g.player.X += g.level.MoveSpeed
	}

	if touchingRight && !touchingLeft {
		g.player.X -= g.level.MoveSpeed
	}

	// Pick up coins
	for _, actor := range g.level.Actors {
		if actor.Exists && actor.Kind == "c" && actor.X+util.FrameWidth/2 > g.player.X-util.FrameWidth && actor.X+util.FrameWidth/2 < g.player.X && actor.Y+util.FrameHeight/2 > g.player.Y && actor.Y+util.FrameHeight/2 < g.player.Y+util.FrameHeight {
			actor.Exists = false
			g.player.Coins++
			// TODO noticeable framerate drop when sounds play
			g.coin.AudioPlayers[g.count%5].Rewind()
			g.coin.AudioPlayers[g.count%5].Play()
		}
	}

	// Fall in holes
	// TODO make player lose money when falling in a hole
	if g.player.Y > util.ScreenHeight*4 {
		spawnX, spawnY := 0, 0
		for _, actor := range g.level.Actors {
			if actor.Kind == "s" {
				if actor.X > spawnX && actor.X < g.player.X {
					spawnX = actor.X + util.FrameWidth
					spawnY = actor.Y
				}

			}
		}
		g.player.X = spawnX
		g.player.Y = spawnY
		g.player.YVelocity = 0
	}

	message = fmt.Sprintf("Gather coins and bring them to the green chest.\nIt costs money to be alive!\nYour coins: %d", g.player.Coins)

	message = message + levelEditor(g)

	return nil
}
