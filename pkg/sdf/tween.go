package sdf

import (
	"time"
)

// Tween -
type Tween struct {
	startTime time.Duration
	duration  time.Duration
	v0, v1    int
	pval      *int
}

// NewTween -
func NewTween() *Tween {
	return &Tween{}
}

// Reset -
func (o *Tween) Reset(pval *int, t0, dur time.Duration, v0, v1 int) {
	o.startTime = t0
	o.duration = dur
	o.v0 = v0
	o.v1 = v1
	o.pval = pval
}

// Process -
func (o *Tween) Process(t time.Duration) (time.Duration, bool) {
	t -= o.startTime
	if t < 0 {
		return 0, false
	}
	if t >= o.duration {
		return t - o.duration, true
	}
	k := float64(t) / float64(o.duration)
	// fmt.Printf("k %v\n", k)
	*o.pval = int(float64(o.v1-o.v0)*k) + o.v0
	return 0, false
}
