package sdf

import (
	"fmt"
	"path"
	"sort"
	"time"
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

	// Animation -
	Animation struct {
		name string
		// speed     float64
		duration  time.Duration
		keyframes []tKeyframe
		// tileSet   *TileSheet
	}
)

// CreateAnimation -
func CreateAnimation(name string) *Animation {
	if animationExists(name) {
		setError(fmt.Errorf("animation %q already exists", name))
		return nil
	}
	anim := &Animation{name: name}
	assets.anims[name] = anim
	return anim
}

// KeyframeT -
func (o *Animation) KeyframeT(t time.Duration, tileName string) *Animation {
	if !Ok() {
		return nil
	}
	tile, ok := assets.tiles[tileName]
	if !ok && tileName != "" {
		setError(fmt.Errorf("tile %q does not exist", tileName))
		return nil
	}
	o.duration = t
	if tile == nil {
		return o
	}
	o.keyframes = append(o.keyframes, tKeyframe{t, tile})
	return o
}

// Keyframe -
func (o *Animation) Keyframe(t float64, tileName string) *Animation {
	if !Ok() {
		return nil
	}
	return o.KeyframeT(time.Duration(t*float64(time.Second)), tileName)
}

// Plain -
func (o *Animation) Plain(tileNames ...string) *Animation {
	if !Ok() {
		return nil
	}
	const defDuration = time.Second
	o.keyframes = nil
	o.duration = 0
	for _, name := range tileNames {
		realNames := getRealTileNames(name)
		if len(realNames) == 0 {
			setError(fmt.Errorf("tile %q does not exist", name))
			return nil
		}
		for _, name := range realNames {
			tile, _ := assets.tiles[name]
			// if !ok {
			// 	setError(fmt.Errorf("tile %q does not exist", name))
			// 	return nil
			// }
			o.keyframes = append(o.keyframes, tKeyframe{o.duration, tile})
			o.duration += defDuration
		} // realNames
	} // tileNames
	return o
}

// StretchTo -
func (o *Animation) StretchTo(dur float64) *Animation {
	if !Ok() {
		return nil
	}
	if o.duration == 0 {
		setError(fmt.Errorf("duration is 0"))
		return nil
	}
	k := dur * float64(time.Second) / float64(o.duration)
	o.duration = time.Duration(k * float64(o.duration))
	for i := range o.keyframes {
		kf := &o.keyframes[i]
		kf.time = time.Duration(k * float64(kf.time))
	}
	return o
}

// StretchToT -
func (o *Animation) StretchToT(dur time.Duration) *Animation {
	if !Ok() {
		return nil
	}
	return o.StretchTo(float64(dur) / float64(time.Second))
}

// Tile -
func (o *Animation) Tile(t time.Duration) *Tile {
	if !Ok() {
		return nil
	}
	t = t % o.duration //math.Remainder(o.time, o.duration)
	len := len(o.keyframes)
	if len == 0 {
		return nil
	}
	ret := o.keyframes[len-1].tile
	// str := "init"
	for i := range o.keyframes {
		kf := &o.keyframes[i]
		if t < kf.time {
			break
		}
		// str = strconv.Itoa(i)
		ret = kf.tile
	}
	// fmt.Println("state: ", str, "time ", t)
	return ret
}

func getRealTileNames(name string) []string {
	tiles := assets.listTiles()
	name = AbsPath(name)
	ret := []string{}
	for _, s := range tiles {
		// fmt.Printf("%v; %v\n", name, s)
		if ok, _ := path.Match(name, s); ok {
			ret = append(ret, s)
		}
	}
	sort.Strings(ret)
	return ret
}
