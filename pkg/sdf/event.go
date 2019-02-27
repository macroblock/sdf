package sdf

import "time"

type (
	// BinaryKey -
	BinaryKey int64

	// KeyboardEvent -
	KeyboardEvent struct {
		Timestamp time.Duration
		Key       int
		Rune      rune
		Mod       uint16
	}
)

// BinaryKey -
func (o KeyboardEvent) BinaryKey() BinaryKey {
	return BinaryKey(o.Mod)<<32 | BinaryKey(o.Rune)
}
