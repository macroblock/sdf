package sdf

import (
	"time"
	"unicode/utf8"

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
	InputDelete
	InputInsert
	InputCopy
	InputPaste
	InputCut
	maxInputKey
)

type (
	tInput    uint
	tKeyState struct {
		pressed     bool
		justPressed bool
		timestamp   time.Duration
	}
)

var textInput string

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

// JustPressed -
func JustPressed(input tInput) bool {
	if input >= maxInputKey {
		return false
	}
	ret := !keyState[input].justPressed
	keyState[input].justPressed = false
	return ret
}

// JustPressedInt -
func JustPressedInt(input tInput) int {
	if input >= maxInputKey || !keyState[input].justPressed {
		return 0
	}
	keyState[input].justPressed = false
	return 1
}

// StartTextInput -
func StartTextInput() {
	textInput = ""
	sdl.StartTextInput()
}

// StopTextInput -
func StopTextInput() {
	sdl.StopTextInput()
}

// TextInput -
func TextInput() string {
	return textInput
}

// IsTextInput -
func IsTextInput() {
	sdl.IsTextInputActive()
}

func processInput() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch ev := event.(type) {
		case *sdl.QuitEvent:
			_ = ev
			sdf.isRunning = false
		case *sdl.KeyboardEvent:
			input := InputAny
			switch {
			case ev.Keysym.Mod&sdl.KMOD_CTRL != 0:
				ev.Keysym.Mod |= sdl.KMOD_CTRL
			case ev.Keysym.Mod&sdl.KMOD_SHIFT != 0:
				ev.Keysym.Mod |= sdl.KMOD_SHIFT
			case ev.Keysym.Mod&sdl.KMOD_ALT != 0:
				ev.Keysym.Mod |= sdl.KMOD_ALT
			}
			switch ev.Keysym.Mod {
			case 0:
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
				case sdl.K_DELETE, sdl.K_BACKSPACE:
					input = InputDelete
				case sdl.K_INSERT:
					input = InputInsert
				} // switch ev.Keysym.Sym
			case sdl.KMOD_CTRL:
				switch ev.Keysym.Sym {
				case sdl.K_INSERT, sdl.K_c:
					input = InputCopy
				case sdl.K_v:
					input = InputPaste
				case sdl.K_x:
					input = InputCut
				} // switch ev.Keysym.Sym {
			case sdl.KMOD_SHIFT:
				switch ev.Keysym.Sym {
				case sdl.K_INSERT:
					input = InputPaste
				} // switch ev.Keysym.Sym {
			case sdl.KMOD_ALT:
				switch ev.Keysym.Sym {
				case sdl.K_INSERT:
					input = InputCut
				} // switch ev.Keysym.Sym {
			} // switch ev.Keysym.Mod
			pressed := (ev.Type == sdl.KEYDOWN)
			keyState[input].pressed = pressed
			keyState[input].timestamp = time.Since(programStart)
			keyState[input].justPressed = pressed
			// keyState[input].timestamp = time.Duration(ev.Timestamp) * time.Millisecond
		case *sdl.TextInputEvent:
			slice := ev.Text[:]
			for len(slice) > 0 {
				r, size := utf8.DecodeRune(slice)
				// fmt.Printf("%c %v\n", r, size)
				if r == '\x00' {
					break
				}
				textInput += string(r)
				slice = slice[size:]
			}
			// textInput += string(slice)
		case *sdl.TextEditingEvent:
		} // switch ev := event.(type)
	}
}
