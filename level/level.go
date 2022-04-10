package level

import (
	"fmt"
	_ "image/png"
	"strconv"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jcgraybill/it-costs-money/sys"
)

type Level struct {
	BgImage1, BgImage2, BgImage3, LevelImage, LevelBackgroundImage, LevelForegroundImage *ebiten.Image
	CoinDecay, MoveSpeed, JumpHeight                                                     int
	Gravity                                                                              float64
	Actors                                                                               []*Actor
}

type Actor struct {
	X, Y   int
	Exists bool
	Kind   string
}

func New(levelNumber int, tiles []*ebiten.Image) Level {
	var l Level
	l.CoinDecay = 90
	l.MoveSpeed = 4
	l.JumpHeight = 8
	l.Gravity = 0.5
	l.BgImage1 = sys.LoadImage("assets/Background_Layer_1.png")
	l.BgImage2 = sys.LoadImage("assets/Background_Layer_2.png")
	l.BgImage3 = sys.LoadImage("assets/Background_Layer_3.png")
	l.LevelImage = generateLevelImage(fmt.Sprintf("leveldata/level_%d_main.csv", levelNumber), tiles)
	l.LevelBackgroundImage = generateLevelImage(fmt.Sprintf("leveldata/level_%d_background.csv", levelNumber), tiles)
	l.LevelForegroundImage = generateLevelImage(fmt.Sprintf("leveldata/level_%d_foreground.csv", levelNumber), tiles)

	l.Actors = loadActors(fmt.Sprintf("leveldata/level_%d_actors.csv", levelNumber))
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

func loadActors(path string) []*Actor {
	actors := make([]*Actor, 0)

	data, err := sys.GameData(path)

	if err == nil {
		for row, line := range strings.Split(string(data), "\n") {
			for col, cell := range strings.Split(line, ",") {
				if cell != "0" && cell != "" {
					actors = append(actors, &Actor{X: col * sys.FrameWidth, Y: row * sys.FrameHeight, Exists: true, Kind: cell})
				}
			}
		}
	} else {
		panic(err)
	}
	return actors
}

func (l Level) StartPosition() (x, y int) {
	x, y = 0, 0
	for _, actor := range l.Actors {
		if actor.Kind == "s" {
			if x == 0 {
				x = actor.X + sys.FrameWidth
				y = actor.Y
			} else if x > actor.X {
				x = actor.X + sys.FrameWidth
				y = actor.Y
			}
		}
	}
	return
}
