package coin

import (
	"bytes"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/jcgraybill/it-costs-money/sys"
)

type Coin struct {
	Slides         []*ebiten.Image
	NumSlides      int
	AnimationSpeed int
	AudioPlayers   [5]*audio.Player
	audioPlayer    *audio.Player
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

	audioBytes, err := sys.GameData("assets/smb_coin.wav")
	if err != nil {
		panic(err)
	}
	d, err := wav.Decode(audioContext, bytes.NewReader(audioBytes))
	if err != nil {
		panic(err)
	}
	c.audioPlayer, err = audioContext.NewPlayer(d)
	if err != nil {
		panic(err)
	}
	return c
}

func (c Coin) PlaySound(count int) {
	c.audioPlayer.Rewind()
	c.audioPlayer.Play()
}
