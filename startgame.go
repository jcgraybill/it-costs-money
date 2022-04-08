package main

import (
	"bytes"
	"fmt"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

func init() {
	frameBuffer = ebiten.NewImage(screenWidth, screenHeight)

	tiles = loadSpriteSheet("assets/1-tiles-city.png")
	tiles = append(tiles, loadSpriteSheet("assets/2-tiles-country.png")...)
	tiles = append(tiles, loadSpriteSheet("assets/3-objects-city.png")...)
	tiles = append(tiles, loadSpriteSheet("assets/4-objects-country.png")...)

	level.coinDecay = 120
	go loadLevel(false)

	actors = loadActors("levels/level_0_actors.csv", false)
	goToStartPosition()

	player.yVelocity = 0
	player.timeSinceLastJump = -jumpRecovery
	player.facingLeft = false
	player.wileECoytoe = wileECoyoteFrames

	runner := loadSpriteSheet("assets/runner.png")
	player.idleFrames = runner[1:6]
	player.runFrames = runner[9:17]
	player.fallFrames = runner[17:21]
	player.coins = 0
	player.slides = &player.idleFrames

	coinSprites := loadSpriteSheet("assets/coin.png")
	coin.slides = coinSprites[1:7]
	coin.numSlides = 6
	coin.animationSpeed = 10
	coin.audioContext = audio.NewContext(sampleRate)

	for i := 0; i < 5; i++ {
		audioBytes, err := assets.ReadFile(fmt.Sprintf("assets/Coins_Grab_0%d.wav", i))
		if err != nil {
			panic(err)
		}
		d, err := wav.Decode(coin.audioContext, bytes.NewReader(audioBytes))
		if err != nil {
			panic(err)
		}
		coin.audioPlayers[i], err = coin.audioContext.NewPlayer(d)
		if err != nil {
			panic(err)
		}

	}

}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("it costs money to be alive")
	if err := ebiten.RunGame(start(&Game{})); err != nil {
		panic(err)
	}
}

func start(g *Game) *Game {
	g.count = 0
	return g
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
