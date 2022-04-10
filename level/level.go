package level

import (
	"fmt"
	_ "image/png"
	"strconv"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/jcgraybill/it-costs-money/level/coin"
	"github.com/jcgraybill/it-costs-money/level/spawn"
	"github.com/jcgraybill/it-costs-money/sys"
)

type Level struct {
	BgImage1, BgImage2, BgImage3, LevelImage, LevelBackgroundImage, LevelForegroundImage *ebiten.Image
	CoinDecay, CoinHolePenalty, MoveSpeed, JumpHeight                                    int
	Gravity                                                                              float64
	Coin                                                                                 coin.Coin
	Spawns                                                                               []*spawn.Spawn
}

func New(levelNumber int, tiles []*ebiten.Image, audioContext *audio.Context) Level {
	var l Level
	l.CoinDecay = 90
	l.CoinHolePenalty = 5
	l.MoveSpeed = 4
	l.JumpHeight = 8
	l.Gravity = 0.5
	l.BgImage1 = sys.LoadImage("assets/Background_Layer_1.png")
	l.BgImage2 = sys.LoadImage("assets/Background_Layer_2.png")
	l.BgImage3 = sys.LoadImage("assets/Background_Layer_3.png")
	l.LevelImage = generateLevelImage(fmt.Sprintf("leveldata/level_%d_main.csv", levelNumber), tiles)
	l.LevelBackgroundImage = generateLevelImage(fmt.Sprintf("leveldata/level_%d_background.csv", levelNumber), tiles)
	l.LevelForegroundImage = generateLevelImage(fmt.Sprintf("leveldata/level_%d_foreground.csv", levelNumber), tiles)
	l.Coin = coin.New(audioContext)
	l.Coin.Coins, l.Spawns = loadActors(fmt.Sprintf("leveldata/level_%d_actors.csv", levelNumber))
	return l
}

func generateLevelImage(path string, tiles []*ebiten.Image) *ebiten.Image {
	levelData := make([][]int, sys.ScreenHeight/sys.FrameHeight)
	levelWidth := 0

	data, err := sys.GameData(path)

	if err == nil {
		for row, line := range strings.Split(string(data), "\n") {
			line = strings.TrimSuffix(line, "\r")
			if len(line) > 0 {
				for col, cell := range strings.Split(line, ",") {
					cellValue, err := strconv.Atoi(cell)
					if err == nil {
						levelData[row] = append(levelData[row], cellValue)
					} else {
						levelData[row] = append(levelData[row], 0)
					}
					if col > levelWidth {
						levelWidth = col
					}
				}
			}
		}
	} else {
		panic(err)
	}

	levelWidth += 1
	// This is a hardware limitation, will vary machine by machine.
	// FIXME slice levels into smaller images so the level can be arbitrarily long
	if levelWidth *= sys.FrameWidth; levelWidth > 16320 {
		levelWidth = 16320
	}

	levelImage := ebiten.NewImage(levelWidth, sys.ScreenHeight)
	for row, line := range levelData {
		for col, cell := range line {
			if cell > 0 && cell < len(tiles) {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(col*sys.FrameWidth), float64(row*sys.FrameHeight))
				levelImage.DrawImage(tiles[cell], op)
			}
		}
	}
	return levelImage
}

func loadActors(path string) ([]*coin.Coins, []*spawn.Spawn) {
	coins := make([]*coin.Coins, 0)
	spawns := make([]*spawn.Spawn, 0)
	data, err := sys.GameData(path)

	if err == nil {
		for row, line := range strings.Split(string(data), "\n") {
			for col, cell := range strings.Split(line, ",") {
				if cell != "0" && cell != "" {
					if cell == "s" {
						spawns = append(spawns, &spawn.Spawn{X: col * sys.FrameWidth, Y: row * sys.FrameHeight})
					} else if cell == "c" {
						coins = append(coins, &coin.Coins{X: col * sys.FrameWidth, Y: row * sys.FrameHeight, Uncollected: true})
					}

				}
			}
		}
	} else {
		panic(err)
	}
	return coins, spawns
}
