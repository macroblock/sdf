package sdf

import (
	"path"
	"time"
)

func setError(err error) {
	if sdf.err != nil {
		return
	}
	sdf.err = err
}

// Time -
func Time() time.Duration {
	return time.Since(programStart)
}

// FixedTime -
func FixedTime() time.Duration {
	return fixedTime
}

func tileSheetExists(name string) bool {
	if _, ok := assets.sheets[name]; ok {
		return true
	}
	return false
}

func tileExists(name string) bool {
	if _, ok := assets.tiles[name]; ok {
		return true
	}
	return false
}

func animationExists(name string) bool {
	if _, ok := assets.anims[name]; ok {
		return true
	}
	return false
}

func joinPaths(a, b string) string {
	if path.IsAbs(b) {
		return path.Clean(b)
	}
	return path.Join(a, b)
}

// AbsTilePath -
func AbsTilePath(path string) string {
	return joinPaths(sdf.curTilePath, path)
}

// TilePath -
func TilePath() string {
	return sdf.curTilePath
}

// SetTilePath -
func SetTilePath(path string) {
	sdf.curTilePath = AbsTilePath(path)
}

// AbsAnimationPath -
func AbsAnimationPath(path string) string {
	return joinPaths(sdf.curAnimationPath, path)
}

// AnimationPath -
func AnimationPath() string {
	return sdf.curAnimationPath
}

// SetAnimationPath -
func SetAnimationPath(path string) {
	sdf.curAnimationPath = AbsAnimationPath(path)
}

// func absName(name string) []string {
// 	if assets.currentTileSheet == nil {
// 		return nil
// 	}
// }
