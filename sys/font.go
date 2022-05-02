package sys

import (
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var Ttf font.Face

func init() {
	ttbytes, err := GameData("assets/Modak-Regular.ttf")
	if err == nil {
		tt, err := opentype.Parse(ttbytes)
		if err == nil {
			var err error
			Ttf, err = opentype.NewFace(tt, &opentype.FaceOptions{
				Size:    36,
				DPI:     72,
				Hinting: font.HintingFull,
			})
			if err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}

	} else {
		panic(err)
	}
}
