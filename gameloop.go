package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	_ "github.com/jcgraybill/it-costs-money/player"
	"github.com/jcgraybill/it-costs-money/sys"
)

func (g *Game) Update() error {
	g.count++

	g.player.Slides = &g.player.IdleFrames

	//TODO if player is mid-jump, will get stuck hovering in the air
	//Will require rearranging what order these checks happen in. What depends on what, or should pre-empt what?

	if g.player.X > g.level.LevelImage.Bounds().Dx()-sys.ScreenWidth/2+sys.FrameWidth/2-1 {
		message = fmt.Sprintf("Congratulations! You collected %d coins.\nPlease place them in the dumpster.", g.player.Coins)
		return nil
	}

	if g.count%g.level.CoinDecay == 0 && g.player.Coins > 0 {
		g.player.Coins--
	}

	// Detect collisions
	touchingGround := false
	if _, _, _, a := g.level.LevelImage.At(g.player.X-sys.FrameWidth/2, g.player.Y+sys.FrameHeight).RGBA(); a > 0 {
		touchingGround = true
	}

	touchingLeft := false
	if _, _, _, a := g.level.LevelImage.At(g.player.X-sys.FrameWidth, g.player.Y+sys.FrameHeight/2).RGBA(); a > 0 {
		touchingLeft = true
	}

	leftAdjacent := false
	if _, _, _, a := g.level.LevelImage.At(g.player.X-sys.FrameWidth-g.level.MoveSpeed, g.player.Y+sys.FrameHeight/2).RGBA(); a > 0 {
		leftAdjacent = true
	}

	touchingRight := false
	if _, _, _, a := g.level.LevelImage.At(g.player.X, g.player.Y+sys.FrameHeight/2).RGBA(); a > 0 {
		touchingRight = true
	}

	rightAdjacent := false
	if _, _, _, a := g.level.LevelImage.At(g.player.X+g.level.MoveSpeed, g.player.Y+sys.FrameHeight/2).RGBA(); a > 0 {
		rightAdjacent = true
	}

	// Jumping

	if touchingGround {
		g.player.YVelocity = 0

		g.player.ResetWileECoyote()

		if ebiten.IsKeyPressed(ebiten.KeySpace) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
			if g.player.TimeSinceLastJump+g.player.JumpRecovery < g.count {
				if _, _, _, a := g.level.LevelImage.At(g.player.X-sys.FrameWidth/2, g.player.Y-g.level.JumpHeight).RGBA(); a == 0 {
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
	if _, _, _, a := g.level.LevelImage.At(g.player.X-sys.FrameWidth/2, g.player.Y).RGBA(); a > 0 {
		touchingTop = true
	}

	// Running
	// TODO allow WASD
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.player.FacingLeft = false
		if g.player.X < g.level.LevelImage.Bounds().Dx()-sys.ScreenWidth/2+sys.FrameWidth/2-1 && !touchingRight && !rightAdjacent {
			g.player.Slides = &g.player.RunFrames
			g.player.X += g.level.MoveSpeed
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.player.FacingLeft = true
		if g.player.X > sys.ScreenWidth/2+sys.FrameWidth/2 && !touchingLeft && !leftAdjacent {
			g.player.Slides = &g.player.RunFrames
			g.player.X -= g.level.MoveSpeed
		}
	}

	// Correct if overlapping with terrain

	if touchingTop && !touchingGround {
		g.player.Y += 1
		g.player.YVelocity = 0
	}

	if _, _, _, a := g.level.LevelImage.At(g.player.X-sys.FrameWidth/2, g.player.Y+sys.FrameHeight-1).RGBA(); a != 0 {
		for a != 0 {
			g.player.Y -= 1
			_, _, _, a = g.level.LevelImage.At(g.player.X-sys.FrameWidth/2, g.player.Y+sys.FrameHeight-1).RGBA()
		}
	}

	if touchingLeft && !touchingRight {
		g.player.X += g.level.MoveSpeed
	}

	if touchingRight && !touchingLeft {
		g.player.X -= g.level.MoveSpeed
	}

	// Pick up coins
	for _, coin := range g.level.Coin.Coins {
		if coin.Uncollected && coin.X+sys.FrameWidth/2 > g.player.X-sys.FrameWidth && coin.X+sys.FrameWidth/2 < g.player.X && coin.Y+sys.FrameHeight/2 > g.player.Y && coin.Y+sys.FrameHeight/2 < g.player.Y+sys.FrameHeight {
			coin.Uncollected = false
			g.player.Coins++

			g.level.Coin.PlaySound(g.count)
		}
	}

	// Fall in holes

	if g.player.Y > sys.ScreenHeight*4 {
		g.player.X, g.player.Y = g.level.PreviousSpawn(g.player.X, g.player.Y)
		g.player.YVelocity = 0
		g.player.Coins = g.player.Coins - g.level.CoinHolePenalty
		if g.player.Coins < 0 {
			g.player.Coins = 0
		}
	}

	message = fmt.Sprintf("Gather coins and bring them to the green chest.\nIt costs money to be alive!\nYour coins: %d", g.player.Coins)

	message = message + levelEditor(g)

	return nil
}
