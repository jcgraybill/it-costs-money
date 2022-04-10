package coin

import (
	"bytes"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/jcgraybill/it-costs-money/sys"
)

type Coin struct {
	Slides         []*ebiten.Image
	NumSlides      int
	AnimationSpeed int
	AudioPlayers   [5]*audio.Player
	Coins          []*Coins
}

type Coins struct {
	X, Y        int
	Uncollected bool
}

func New(audioContext *audio.Context) Coin {
	var c Coin
	coinSprites := sys.LoadSpriteSheet("assets/coin.png")
	c.Slides = coinSprites[1:7]
	c.NumSlides = 6
	c.AnimationSpeed = 10

	c.Coins = make([]*Coins, 0)

	for i := 0; i < 5; i++ {
		audioBytes, err := sys.GameData(fmt.Sprintf("assets/Coins_Grab_0%d.ogg", i))
		if err != nil {
			panic(err)
		}
		d, err := vorbis.Decode(audioContext, bytes.NewReader(audioBytes))
		if err != nil {
			panic(err)
		}
		c.AudioPlayers[i], err = audioContext.NewPlayer(d)
		if err != nil {
			panic(err)
		}

	}
	return c
}

func (c Coin) PlaySound(count int) {
	c.AudioPlayers[count%5].Rewind()
	c.AudioPlayers[count%5].Play()
}
