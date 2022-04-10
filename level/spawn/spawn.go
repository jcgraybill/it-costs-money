package spawn

type Spawn struct {
	X, Y int
}

func New(x, y int) Spawn {
	var s Spawn
	s.X, s.Y = x, y
	return s
}
