package player

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jcgraybill/it-costs-money/sys"
)

const (
	wileECoyoteFrames = 16
)

type Player struct {
	X, Y                              int
	FacingLeft                        bool
	YVelocity                         float64
	WileECoyote                       int
	TimeSinceLastJump                 int
	Slides                            *[]*ebiten.Image
	IdleFrames, RunFrames, FallFrames []*ebiten.Image
	Coins                             int
	JumpRecovery                      int
}

func (p Player) ResetWileECoyote() {
	p.WileECoyote = wileECoyoteFrames
}

func New() Player {
	var p Player
	p.YVelocity = 0
	p.JumpRecovery = 40
	p.TimeSinceLastJump = -p.JumpRecovery
	p.FacingLeft = false
	p.WileECoyote = wileECoyoteFrames

	p.IdleFrames = sys.RunnerTiles[1:6]
	p.RunFrames = sys.RunnerTiles[9:17]
	p.FallFrames = sys.RunnerTiles[17:21]
	p.Coins = 0
	p.Slides = &p.IdleFrames
	return p
}
