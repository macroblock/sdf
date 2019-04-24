package event

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

	// Keyboard -
	Keyboard struct {
		Event
		// scancode
		Key  int
		Rune rune
		Mod  uint16
	}

	// MouseClick -
	MouseClick struct {
		Event
		Pressed bool
		Button  uint32
		X, Y    int
	}

	// MouseMotion -
	MouseMotion struct {
		Event
		Buttons uint32
		X, Y    int
		DX, DY  int
	}
)

// BinaryKey -
func (o Keyboard) BinaryKey() BinaryKey {
	return BinaryKey(o.Mod)<<32 | BinaryKey(o.Rune)
}

// BinaryKey -
func (o MouseClick) BinaryKey() BinaryKey {
	return -1
}

// BinaryKey -
func (o MouseMotion) BinaryKey() BinaryKey {
	return -1
}

func (o Keyboard) String() string {
	return fmt.Sprintf("kbd %q", o.Rune)
}

func (o MouseClick) String() string {
	return fmt.Sprintf("click (%v, %v) %v %v", o.X, o.Y, o.Pressed, o.Button)
}

func (o MouseMotion) String() string {
	return fmt.Sprintf("motion (%v, %v) [%v, %v] %v", o.X, o.Y, o.DX, o.DY, o.Buttons)
}
