package main

import (
	"bytes"
	"image"
	_ "image/png"
	"os"
	"strconv"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
)

func generateLevelImage(path string, xSize, ySize int, live bool) *ebiten.Image {
	levelData := make([][]int, screenHeight/ySize)
	levelWidth := 0

	var data []byte
	var err error

	if live {
		data, err = os.ReadFile(path)
	} else {
		data, err = assets.ReadFile(path)
	}

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
	//OpenGL limitation
	if levelWidth *= xSize; levelWidth > 16320 {
		levelWidth = 16320
	}

	levelImage := ebiten.NewImage(levelWidth, screenHeight)
	for row, line := range levelData {
		for col, cell := range line {
			if cell > 0 && cell < len(tiles) {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(col*xSize), float64(row*ySize))
				levelImage.DrawImage(tiles[cell], op)
			}
		}
	}
	return levelImage
}

func loadImage(path string) *ebiten.Image {
	imgBytes, err := assets.ReadFile(path)
	if err == nil {
		img, _, err := image.Decode(bytes.NewReader(imgBytes))
		if err == nil {
			return ebiten.NewImageFromImage(img)
		}
		panic(err)
	}
	panic(err)
}

func loadSpriteSheet(path string) []*ebiten.Image {
	numberofSprites := 0
	spriteSheet := loadImage(path)
	numberofSprites += spriteSheet.Bounds().Dx() / frameWidth * spriteSheet.Bounds().Dy() / frameHeight
	sprites := make([]*ebiten.Image, numberofSprites+2)
	i := 1

	for y := 0; y < spriteSheet.Bounds().Dy()/frameHeight; y++ {
		for x := 0; x < spriteSheet.Bounds().Dx()/frameWidth; x++ {
			sprites[i] = spriteSheet.SubImage(image.Rect(x*frameWidth, y*frameHeight, x*frameWidth+frameWidth, y*frameHeight+frameHeight)).(*ebiten.Image)
			i++
		}
	}

	return sprites
}

func parallax(image *ebiten.Image, offset int, speed int) {
	panelWidth := image.Bounds().Dx()
	position := (offset / speed) % panelWidth
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(position), 0)
	frameBuffer.DrawImage(image, op)

	for i := 0; position+i*panelWidth+panelWidth < screenWidth; i++ {
		op.GeoM.Translate(float64(panelWidth), 0)
		frameBuffer.DrawImage(image, op)
	}
}

func loadActors(path string, live bool) []*Actor {
	actors = make([]*Actor, 0)

	var data []byte
	var err error

	if live {
		data, err = os.ReadFile(path)
	} else {
		data, err = assets.ReadFile(path)
	}

	if err == nil {
		for row, line := range strings.Split(string(data), "\n") {
			for col, cell := range strings.Split(line, ",") {
				if cell != "0" {
					actors = append(actors, &Actor{x: col * frameWidth, y: row * frameHeight, exists: true, kind: cell})
				}
			}
		}
	} else {
		panic(err)
	}
	return actors
}
