package misc

import "time"

// MinInt -
func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// MaxInt -
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// FloatToTime -
func FloatToTime(t float64) time.Duration {
	return time.Duration(t * float64(time.Second))
}
