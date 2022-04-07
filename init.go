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
	level.coinDecay = 240
	level.levelImage = generateLevelImage("levels/level_0_main.csv", frameWidth, frameHeight, live)
	level.levelBackgroundImage = generateLevelImage("levels/level_0_background.csv", frameWidth, frameHeight, live)
	level.levelForegroundImage = generateLevelImage("levels/level_0_foreground.csv", frameWidth, frameHeight, live)
	level.levelJunkImage = generateLevelImage("levels/level_0_junk.csv", frameWidth/2, frameHeight/2, live)
}

func init() {
	frameBuffer = ebiten.NewImage(screenWidth, screenHeight)
	level.bgImage1 = loadImage("assets/Background_Layer_1.png")
	level.bgImage2 = loadImage("assets/Background_Layer_2.png")
	level.bgImage3 = loadImage("assets/Background_Layer_3.png")

	tiles = loadSpriteSheet("assets/tileset.png")
	tiles = append(tiles, loadSpriteSheet("assets/tileset2.png")...)
	tiles = append(tiles, loadSpriteSheet("assets/objects.png")...)
	tiles = append(tiles, loadSpriteSheet("assets/objects2.png")...)

	go loadLevel(false)

	player.x = screenWidth/2 + frameWidth/2
	player.y = screenHeight - groundHeight - frameHeight
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

	actors = loadActors("levels/level_0_actors.csv", false)
}
