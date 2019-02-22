package sdf

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

// -
const (
	InputAny tInput = iota
	InputAccept
	InputCancel
	InputSelect
	InputUp
	InputDown
	InputLeft
	InputRight
	InputPause
	maxInputKey
)

type (
	tInput    uint
	tKeyState struct {
		pressed   bool
		timestamp time.Duration
	}
)

var keyState []tKeyState

func init() {
	keyState = make([]tKeyState, maxInputKey)
}

// Pressed -
func Pressed(input tInput) bool {
	if input >= maxInputKey {
		return false
	}
	return keyState[input].pressed
}

// PressedInt -
func PressedInt(input tInput) int {
	if input >= maxInputKey || !keyState[input].pressed {
		return 0
	}
	return 1
}

func processInput() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch ev := event.(type) {
		case *sdl.QuitEvent:
			_ = ev
			sdf.isRunning = false
		case *sdl.KeyboardEvent:
			input := InputAny
			switch ev.Keysym.Sym {
			case sdl.K_SPACE, sdl.K_RETURN:
				input = InputAccept
			case sdl.K_ESCAPE:
				input = InputCancel
			case sdl.K_TAB:
				input = InputSelect
			case sdl.K_UP:
				input = InputUp
			case sdl.K_DOWN:
				input = InputDown
			case sdl.K_LEFT:
				input = InputLeft
			case sdl.K_RIGHT:
				input = InputRight
			case sdl.K_p, sdl.K_PAUSE:
				input = InputPause
			}
			keyState[input].pressed = (ev.Type == sdl.KEYDOWN)
			keyState[input].timestamp = time.Since(programStart)
			// keyState[input].timestamp = time.Duration(ev.Timestamp) * time.Millisecond
		} // switch ev := event.(type)

	}
}
