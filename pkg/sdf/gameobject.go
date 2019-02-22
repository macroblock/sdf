package sdf

import "github.com/macroblock/sdf/pkg/geom"

type (
	// IGameObject -
	IGameObject interface {
		Do(string) bool
		Ready() bool
	}

	// GameObject -
	GameObject struct {
		pos      geom.Point2i
		oldState string
		curState string
		oldTile  IElem
		curTile  IElem
		sprites  map[string]IElem
	}
)
