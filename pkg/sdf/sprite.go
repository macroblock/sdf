package sdf

import (
	"time"

	"github.com/macroblock/sdf/pkg/misc"
)

type (
	// Sprite -
	Sprite struct {
		// name         string
		startTime    time.Duration
		internalTime time.Duration
		duration     time.Duration
		speed        float64
		curTile      *Tile
		tile         *Tile
		keyframes    []tKeyframe
		tileSet      *TileSheet
		suspended    bool
	}

	tKeyframe struct {
		time time.Duration
		tile *Tile
	}
)

// // DeltaTimeFunc -
// func DeltaTimeFunc() func() time.Duration {
// 	lastUpdate := time.Now()
// 	dt := time.Since(lastUpdate)
// 	return func() time.Duration {
// 		dt = time.Since(lastUpdate)
// 		lastUpdate = time.Now()
// 		return dt
// 	}
// }

func newSprite(tileSet *TileSheet, dur float64, tile *Tile, keyframes []tKeyframe) *Sprite {
	return &Sprite{
		// name:         name,
		duration:  misc.FloatToTime(dur),
		tile:      tile,
		keyframes: keyframes,
		speed:     1.0,
		tileSet:   tileSet,
	}
}

// SetSpeed -
func (o *Sprite) SetSpeed(speed float64) *Sprite {
	if speed < 0 || o == nil {
		return o
	}
	o.speed = speed
	return o
}

// Tile -
func (o *Sprite) Tile() *Tile {
	time := o.internalTime
	if !o.suspended {
		delta := FixedTime() - o.startTime
		// d := time.Duration(float64(delta) * o.speed)
		time = delta % o.duration //math.Remainder(o.time, o.duration)
		o.internalTime = time
	}
	ret := &o.tile
	// str := "init"
	for i := range o.keyframes {
		keyframe := &o.keyframes[i]
		if time < keyframe.time {
			break
		}
		// str = strconv.Itoa(i)
		ret = &keyframe.tile
	}
	// fmt.Println("state: ", str, "time ", o.time)
	return *ret
}

// AddKeyframe -
func (o *Sprite) AddKeyframe(time float64, tileName string) *Sprite {
	if !Ok() || o == nil {
		return nil
	}
	tile := o.tileSet.Tile(tileName)
	if tile == nil {
		return nil
	}
	o.keyframes = append(o.keyframes, tKeyframe{misc.FloatToTime(time), tile})
	return o
}

// // Update -
// func (o *Sprite) Update(delta time.Duration) bool {
// 	tile := o.curTile
// 	o.curTile = o.Tile(delta)
// 	return o.curTile != tile
// }

// Copy -
func (o *Sprite) Copy(x, y int) {
	// o.curTile.Copy(x, y)
	o.Tile().Copy(x, y)
}

// Run -
func (o *Sprite) Run() {
	if !o.suspended {
		return
	}
	o.startTime = FixedTime() - o.internalTime
	o.suspended = false
}

// Suspend -
func (o *Sprite) Suspend() {
	o.suspended = true
}

// Reset -
func (o *Sprite) Reset() {
	o.startTime = FixedTime()
	o.internalTime = 0
}

// Suspended -
func (o *Sprite) Suspended() bool {
	return o.suspended
}
