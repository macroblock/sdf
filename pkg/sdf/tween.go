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
	if o.pval == nil {
		return 0, true
	}
	t -= o.startTime
	if t <= 0 {
		*o.pval = o.v0
		return 0, false
	}
	if t >= o.duration {
		// fmt.Println("tween : ", o.v1)
		*o.pval = o.v1
		return t - o.duration, true
	}
	k := float64(t) / float64(o.duration)
	// fmt.Printf("k %v\n", k)
	*o.pval = int(float64(o.v1-o.v0)*k) + o.v0
	return 0, false
}
