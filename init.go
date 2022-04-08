package main

import (
	"bytes"
	"fmt"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

func loadLevel(live bool) {
	level.bgImage1 = loadImage("assets/Background_Layer_1.png")
	level.bgImage2 = loadImage("assets/Background_Layer_2.png")
	level.bgImage3 = loadImage("assets/Background_Layer_3.png")
	level.levelImage = generateLevelImage("levels/level_0_main.csv", frameWidth, frameHeight, live)
	level.levelBackgroundImage = generateLevelImage("levels/level_0_background.csv", frameWidth, frameHeight, live)
	level.levelForegroundImage = generateLevelImage("levels/level_0_foreground.csv", frameWidth, frameHeight, live)
}

func init() {
	frameBuffer = ebiten.NewImage(screenWidth, screenHeight)

	tiles = loadSpriteSheet("assets/1-tiles-city.png")
	tiles = append(tiles, loadSpriteSheet("assets/2-tiles-country.png")...)
	tiles = append(tiles, loadSpriteSheet("assets/3-objects-city.png")...)
	tiles = append(tiles, loadSpriteSheet("assets/4-objects-country.png")...)

	level.coinDecay = 120
	go loadLevel(false)

	actors = loadActors("levels/level_0_actors.csv", false)

	player.x, player.y = 0, 0
	for _, actor := range actors {
		if actor.kind == "s" {
			if player.x == 0 {
				player.x = actor.x + frameWidth
				player.y = actor.y
			} else if player.x > actor.x {
				player.x = actor.x + frameWidth
				player.y = actor.y
			}
		}
	}

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
