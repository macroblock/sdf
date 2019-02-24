package sdf

import (
	"fmt"
	"time"

	"github.com/macroblock/sdf/pkg/geom"
)

type (
	// IGameObject -
	IGameObject interface {
		Do(string) bool
		Ready() bool
	}

	// GameObject -
	GameObject struct {
		name     string
		pos      geom.Point2i
		oldState string
		curState string
		oldTile  IElem
		curTile  IElem

		curAnim *Animation
		anims   map[string]*Animation

		suspended    bool
		startTime    time.Duration
		internalTime time.Duration
	}
)

// NewGameObject -
func NewGameObject(name string) *GameObject {
	ret := &GameObject{
		name:  name,
		anims: map[string]*Animation{},
	}
	return ret
}

// AddAnimation -
func (o *GameObject) AddAnimation(alias string, name string) *GameObject {
	if o == nil {
		return nil
	}
	if alias == "" {
		alias = name
	}
	if _, ok := o.anims[alias]; ok {
		setError(fmt.Errorf("state %q already exists", alias))
		return nil
	}
	if !animationExists(name) {
		setError(fmt.Errorf("animation %q does not exists", name))
		return nil
	}
	anim := assets.anims[name]
	o.anims[alias] = anim
	if o.curAnim == nil {
		o.curAnim = anim
	}
	return o
}

// Tile -
func (o *GameObject) Tile() *Tile {
	return o.curAnim.Tile(FixedTime() - o.startTime)
}

// Copy -
func (o *GameObject) Copy(x, y int) {
	// o.curTile.Copy(x, y)
	o.Tile().Copy(x, y)
}

// Play -
func (o *GameObject) Play(name string) {
	anim, ok := o.anims[name]
	if !ok {
		setError(fmt.Errorf("animation %q does not exists", name))
		return
	}
	if o.curAnim == anim {
		// fmt.Printf("name %q\n", o.curAnim.name)
		// fmt.Printf("startTime %v, fixedTime %v\n", o.startTime, FixedTime())
		// fmt.Printf("internal %v\n", o.internalTime)
		return
	}
	o.curAnim = anim
	o.Reset()
	// fmt.Printf("name %q\n", o.curAnim.name)
	// fmt.Printf("startTime %v, fixedTime %v\n", o.startTime, FixedTime())
	// fmt.Printf("internal %v\n", o.internalTime)
}

// Continue -
func (o *GameObject) Continue() {
	if o.suspended {
		o.startTime = FixedTime() - o.internalTime
		o.suspended = false
		return
	}
}

// Reset -
func (o *GameObject) Reset() {
	o.startTime = FixedTime()
	o.internalTime = 0
}

// Suspend -
func (o *GameObject) Suspend() {
	o.suspended = true
}

// Suspended -
func (o *GameObject) Suspended() bool {
	return o.suspended
}