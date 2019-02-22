package sdf

import (
	"fmt"
	"strconv"
	"time"

	"github.com/macroblock/sdf/pkg/misc"
)

type (
	// Sprite -
	Sprite struct {
		// name         string
		time      time.Duration
		duration  time.Duration
		speed     float64
		tile      *Tile
		keyframes []tKeyframe
		tileSet   *TileSheet
	}

	tKeyframe struct {
		time time.Duration
		tile *Tile
	}
)

// DeltaTimeFunc -
func DeltaTimeFunc() func() time.Duration {
	lastUpdate := time.Now()
	dt := time.Since(lastUpdate)
	return func() time.Duration {
		dt = time.Since(lastUpdate)
		lastUpdate = time.Now()
		return dt
	}
}

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
func (o *Sprite) Tile(delta time.Duration) *Tile {
	d := time.Duration(float64(delta) * o.speed)
	o.time += d
	if o.time >= o.duration {
		o.time = o.time % o.duration //math.Remainder(o.time, o.duration)
	}
	ret := &o.tile
	str := "init"
	for i := range o.keyframes {
		keyframe := &o.keyframes[i]
		if o.time < keyframe.time {
			break
		}
		str = strconv.Itoa(i)
		ret = &keyframe.tile
	}
	fmt.Println("state: ", str, "time ", o.time)
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

// Copy -
func (o *Sprite) Copy(x, y int, delta time.Duration) {
	tile := o.Tile(delta)
	tile.Copy(x, y, -1)
}
