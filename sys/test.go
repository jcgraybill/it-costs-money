//go:build !deploy
// +build !deploy

package sys

import "os"

func GameData(path string) ([]byte, error) {
	return os.ReadFile("sys/" + path)
}
