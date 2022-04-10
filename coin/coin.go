// TODO coin should be contained within level, since it's an actor

package coin

import (
	"bytes"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/jcgraybill/it-costs-money/util"
)

type Coin struct {
	Slides         []*ebiten.Image
	NumSlides      int
	AnimationSpeed int
	audioContext   *audio.Context
	AudioPlayers   [5]*audio.Player
	sampleRate     int
}

func New() Coin {
	var c Coin
	coinSprites := util.LoadSpriteSheet("assets/coin.png")
	c.Slides = coinSprites[1:7]
	c.NumSlides = 6
	c.AnimationSpeed = 10
	c.sampleRate = 48000
	c.audioContext = audio.NewContext(c.sampleRate)

	for i := 0; i < 5; i++ {
		audioBytes, err := util.GameData(fmt.Sprintf("assets/Coins_Grab_0%d.wav", i))
		if err != nil {
			panic(err)
		}
		d, err := wav.Decode(c.audioContext, bytes.NewReader(audioBytes))
		if err != nil {
			panic(err)
		}
		c.AudioPlayers[i], err = c.audioContext.NewPlayer(d)
		if err != nil {
			panic(err)
		}

	}
	return c
}
