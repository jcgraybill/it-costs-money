package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/jcgraybill/it-costs-money/level"
	"github.com/jcgraybill/it-costs-money/sys"
)

func (g *Game) Update() error {
	g.count++
	g.player.Slides = &g.player.IdleFrames

	touchingGround, touchingLeft, leftAdjacent, touchingRight, rightAdjacent, touchingTop := detectCollisions(g)

	if !touchingGround {
		fall(g)
	}

	if atEndOfLevel(g) {
		levelEnd(g)
		return nil
	}

	go loseCoins(g)

	if touchingGround {
		jump(g)
	}

	run(g, touchingLeft, leftAdjacent, touchingRight, rightAdjacent)

	correctTerrainOverlap(g, touchingGround, touchingLeft, touchingRight, touchingTop)

	pickUpCoins(g)

	fallInHoles(g)

	message = levelEditor(g)

	return nil
}

func loseCoins(g *Game) {
	if g.count%g.level.CoinDecay == 0 && g.player.Coins > 0 {
		g.player.Coins--
		sys.DropCoin(g.count)
	}
}

func levelEnd(g *Game) {

	var message string
	if g.level.LevelNumber == 1 {
		message = fmt.Sprintf("You collected %d coins.\nG: try again\nN: next level", g.player.Coins)
	} else {
		message = fmt.Sprintf("Thanks for playing.\nMore content (including this level)\ncoming soon.")
	}
	sys.WriteMessage(g.level.LevelBackgroundImage.Bounds().Dx()-600, 150, message, g.level.LevelBackgroundImage)

	if inpututil.IsKeyJustPressed(ebiten.KeyG) {
		g.level = level.New(g.level.LevelNumber)
		g.player.X, g.player.Y = g.level.StartPosition()
		g.player.Coins = 0
	}
	if g.level.LevelNumber == 1 && inpututil.IsKeyJustPressed(ebiten.KeyN) {
		g.level.EndAudio()
		g.level = level.New(2)
		g.player.X, g.player.Y = g.level.StartPosition()
		g.player.Coins = 0
	}
}

func fallInHoles(g *Game) {
	if g.player.Y > sys.ScreenHeight*4 {
		g.player.X, g.player.Y = g.level.PreviousSpawn(g.player.X, g.player.Y)
		g.player.YVelocity = 0
		g.player.Coins = g.player.Coins - g.level.CoinHolePenalty
		if g.player.Coins < 0 {
			g.player.Coins = 0
		}
	}
}

func pickUpCoins(g *Game) {
	for _, coin := range g.level.Coin.Coins {
		if coin.Uncollected && coin.X+sys.FrameWidth/2 > g.player.X-sys.FrameWidth && coin.X+sys.FrameWidth/2 < g.player.X && coin.Y+sys.FrameHeight/2 > g.player.Y && coin.Y+sys.FrameHeight/2 < g.player.Y+sys.FrameHeight {
			coin.Uncollected = false
			g.player.Coins++
			go sys.PickupCoin()
		}
	}
}

func correctTerrainOverlap(g *Game, touchingGround, touchingLeft, touchingRight, touchingTop bool) {
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

}

// TODO allow WASD
func run(g *Game, touchingLeft, leftAdjacent, touchingRight, rightAdjacent bool) {
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
}

func fall(g *Game) {
	g.player.Y += int(g.player.YVelocity)
	g.player.YVelocity += g.level.Gravity
	if g.player.WileECoyote == 0 {
		g.player.Slides = &g.player.FallFrames
	} else {
		g.player.WileECoyote -= 1
	}
}

func jump(g *Game) {
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
}

func detectCollisions(g *Game) (bool, bool, bool, bool, bool, bool) {
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

	touchingTop := false
	if _, _, _, a := g.level.LevelImage.At(g.player.X-sys.FrameWidth/2, g.player.Y).RGBA(); a > 0 {
		touchingTop = true
	}
	return touchingGround, touchingLeft, leftAdjacent, touchingRight, rightAdjacent, touchingTop
}

func atEndOfLevel(g *Game) bool {
	if g.player.X > g.level.LevelImage.Bounds().Dx()-sys.ScreenWidth/2+sys.FrameWidth/2-1 {
		return true
	}
	return false
}
