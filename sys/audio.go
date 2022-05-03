package sys

import (
	"bytes"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

var AudioContext *audio.Context
var dropCoinSound [5]*audio.Player
var pickupCoinSound *audio.Player
var ambientSound *audio.Player

func init() {
	AudioContext = audio.NewContext(44100)

	for i := 0; i < 5; i++ {
		audioBytes, err := GameData(fmt.Sprintf("assets/Coins_Grab_0%d.ogg", i))
		if err != nil {
			panic(err)
		}

		d, err := vorbis.Decode(AudioContext, bytes.NewReader(audioBytes))
		if err != nil {
			panic(err)
		}
		dropCoinSound[i], err = AudioContext.NewPlayer(d)
		if err != nil {
			panic(err)
		}
	}

	audioBytes, err := GameData("assets/smb_coin.wav")
	if err != nil {
		panic(err)
	}
	d, err := wav.Decode(AudioContext, bytes.NewReader(audioBytes))
	if err != nil {
		panic(err)
	}
	pickupCoinSound, err = AudioContext.NewPlayer(d)
	if err != nil {
		panic(err)
	}
}

func DropCoin(count int) {
	dropCoinSound[count%5].Rewind()
	dropCoinSound[count%5].Play()
}

func PickupCoin() {
	pickupCoinSound.Rewind()
	pickupCoinSound.Play()
}

func PlayLevelAmbience(level int) {
	audioBytes, err := GameData(fmt.Sprintf("assets/ambience-level-%d.ogg", level))
	if err == nil {
		d, err := vorbis.Decode(AudioContext, bytes.NewReader(audioBytes))
		if err == nil {
			s := audio.NewInfiniteLoop(d, d.Length())
			ambientSound, err = AudioContext.NewPlayer(s)
			if err == nil {
				ambientSound.Play()
			} else {
				panic(err)
			}
		} else {
			panic(err)
		}
	} else {
		panic(err)
	}
}

func StopLevelAmbience() {
	if ambientSound.IsPlaying() {
		ambientSound.Pause()
	}
}
