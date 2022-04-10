//go:build deploy
// +build deploy

package sys

import (
	"embed"
)

func GameData(path string) ([]byte, error) {
	return gd.ReadFile(path)
}

//go:embed assets leveldata
var gd embed.FS
