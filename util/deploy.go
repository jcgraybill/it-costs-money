//go:build deploy
// +build deploy

package util

import(
	"embed"
)

// FIXME "util\deploy.go:10:12: pattern assets: no matching files found"
// embed only works on subdirectories of a package, so re-think this

//go:embed assets level/leveldata
var gd embed.FS

func GameData(path string) []byte, error {
	return gd.ReadFile(path)
}