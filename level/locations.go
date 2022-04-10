package level

import "github.com/jcgraybill/it-costs-money/sys"

// FIXME player starts slightly to the right of start position

func (l Level) StartPosition() (int, int) {
	return l.Spawns[0].X + sys.FrameWidth, l.Spawns[0].Y
}

func (l Level) PreviousSpawn(x, y int) (int, int) {
	spawnX, spawnY := 0, 0
	for _, spawn := range l.Spawns {
		if spawn.X > spawnX && spawn.X < x {
			spawnX = spawn.X + sys.FrameWidth
			spawnY = spawn.Y
		}
	}
	return spawnX, spawnY
}

func (l Level) NextSpawn(x, y int) (int, int) {
	spawnX, spawnY := l.LevelImage.Bounds().Dx(), 0
	foundSpawn := false
	for _, spawn := range l.Spawns {
		if spawn.X > x && spawn.X < spawnX {
			spawnX = spawn.X + sys.FrameWidth
			spawnY = spawn.Y
			foundSpawn = true
		}
	}
	if foundSpawn {
		return spawnX, spawnY
	} else {
		return l.StartPosition()
	}

}
