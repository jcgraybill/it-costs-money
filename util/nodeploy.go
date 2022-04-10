//go:build !deploy
// +build !deploy

package util

import "os"

func GameData(path string) ([]byte, error) {
	return os.ReadFile(path)
}
