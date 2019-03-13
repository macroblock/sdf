package sdf

import (
	"fmt"
	"time"
)

type (
	// BinaryKey -
	BinaryKey int64

	// IEvent -
	IEvent interface {
		BinaryKey() BinaryKey
	}

	// Event -
	Event struct {
		Timestamp time.Duration
	}

	// KeyboardEvent -
	KeyboardEvent struct {
		Event
		Key  int
		Rune rune
		Mod  uint16
	}

	// MouseClickEvent -
	MouseClickEvent struct {
		Event
		Pressed bool
		Button  uint32
		X, Y    int
	}

	// MouseMotionEvent -
	MouseMotionEvent struct {
		Event
		Buttons uint32
		X, Y    int
		DX, DY  int
	}
)

// BinaryKey -
func (o KeyboardEvent) BinaryKey() BinaryKey {
	return BinaryKey(o.Mod)<<32 | BinaryKey(o.Rune)
}

// BinaryKey -
func (o MouseClickEvent) BinaryKey() BinaryKey {
	return -1
}

// BinaryKey -
func (o MouseMotionEvent) BinaryKey() BinaryKey {
	return -1
}

func (o KeyboardEvent) String() string {
	return fmt.Sprintf("kbd %q", o.Rune)
}

func (o MouseClickEvent) String() string {
	return fmt.Sprintf("click (%v, %v) %v %v", o.X, o.Y, o.Pressed, o.Button)
}

func (o MouseMotionEvent) String() string {
	return fmt.Sprintf("motion (%v, %v) [%v, %v] %v", o.X, o.Y, o.DX, o.DY, o.Buttons)
}
