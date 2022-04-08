package main

import (
	_ "image/png"
)

func loadLevel(live bool) {
	level.bgImage1 = loadImage("assets/Background_Layer_1.png")
	level.bgImage2 = loadImage("assets/Background_Layer_2.png")
	level.bgImage3 = loadImage("assets/Background_Layer_3.png")
	level.levelImage = generateLevelImage("levels/level_0_main.csv", frameWidth, frameHeight, live)
	level.levelBackgroundImage = generateLevelImage("levels/level_0_background.csv", frameWidth, frameHeight, live)
	level.levelForegroundImage = generateLevelImage("levels/level_0_foreground.csv", frameWidth, frameHeight, live)
}
