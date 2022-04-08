package main

import (
	"embed"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

const (
	moveSpeed         = 4
	jumpHeight        = 8
	gravity           = 0.5
	jumpRecovery      = 40
	wileECoyoteFrames = 16
	enableLevelReload = true

	screenWidth  = 1024
	screenHeight = 512
	frameWidth   = 64
	frameHeight  = 64

	sampleRate = 48000
)

var level Level
var player Player
var frameBuffer *ebiten.Image
var coin Coin
var actors []*Actor
var message string
var tiles []*ebiten.Image

//go:embed assets levels
var assets embed.FS

type Game struct {
	count int
}

type Actor struct {
	x, y   int
	exists bool
	kind   string
}

type Coin struct {
	slides         []*ebiten.Image
	numSlides      int
	animationSpeed int
	audioContext   *audio.Context
	audioPlayers   [5]*audio.Player
}

type Level struct {
	bgImage1, bgImage2, bgImage3, levelImage, levelBackgroundImage, levelForegroundImage *ebiten.Image
	coinDecay                                                                            int
	startingYPosition                                                                    int
}

type Player struct {
	x, y                              int
	facingLeft                        bool
	yVelocity                         float64
	wileECoytoe                       int
	timeSinceLastJump                 int
	slides                            *[]*ebiten.Image
	idleFrames, runFrames, fallFrames []*ebiten.Image
	coins                             int
}

func start(g *Game) *Game {
	g.count = 0
	return g
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("it costs money to be alive")
	if err := ebiten.RunGame(start(&Game{})); err != nil {
		panic(err)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
