package sdf

import (
	"fmt"
	"time"
	"unicode/utf8"

	"github.com/veandco/go-sdl2/sdl"
)

// -
const (
	InputAny    = 0 // tInput = iota
	InputAccept = sdl.SCANCODE_KP_ENTER
	InputCancel = sdl.SCANCODE_ESCAPE
	InputSelect = sdl.SCANCODE_TAB
	InputUp     = sdl.SCANCODE_UP
	InputDown   = sdl.SCANCODE_DOWN
	InputLeft   = sdl.SCANCODE_LEFT
	InputRight  = sdl.SCANCODE_RIGHT
	InputPause  = sdl.SCANCODE_P
	InputDelete = sdl.SCANCODE_DELETE
	InputInsert = sdl.SCANCODE_INSERT
	InputCopy   = sdl.SCANCODE_COPY
	InputPaste  = sdl.SCANCODE_PASTE
	InputCut    = sdl.SCANCODE_CUT
	maxInputKey
)

type (
	tInput    uint
	tKeyState struct {
		justPressed bool
		timestamp   time.Duration
		scan        int
	}
)

var (
	eventQueue []KeyboardEvent
	scanbuf    []tKeyState
)

func init() {
	eventQueue = make([]KeyboardEvent, 0, 32)
}

// Pressed -
func Pressed(scan int) bool {
	for i := range scanbuf {
		if scanbuf[i].scan == scan {
			return true
		}
	}
	return false
}

// PressedInt -
func PressedInt(scan int) int {
	for i := range scanbuf {
		if scanbuf[i].scan == scan {
			return 1
		}
	}
	return 0
}

// JustPressed -
func JustPressed(scan int) bool {
	for i := range scanbuf {
		if scanbuf[i].scan == scan && scanbuf[i].justPressed {
			scanbuf[i].justPressed = false
			return true
		}
	}
	return false
}

// JustPressedInt -
func JustPressedInt(scan int) int {
	for i := range scanbuf {
		if scanbuf[i].scan == scan && scanbuf[i].justPressed {
			scanbuf[i].justPressed = false
			return 1
		}
	}
	return 0
}

func processInput() {
	// lastKbdEventIndex := -1
	lastKbdEvent := (*KeyboardEvent)(nil)
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch ev := event.(type) {
		case *sdl.QuitEvent:
			_ = ev
			sdf.isRunning = false
		case *sdl.KeyboardEvent:
			pressed := (ev.Type == sdl.KEYDOWN)
			time := time.Since(programStart)
			if pressed {
				// lastKbdEventIndex = len(eventQueue)
				eventQueue = append(eventQueue, KeyboardEvent{})
				lastKbdEvent = &eventQueue[len(eventQueue)-1]
				lastKbdEvent.Key = int(ev.Keysym.Scancode)
				lastKbdEvent.Rune = utf8.RuneError
				lastKbdEvent.Mod = ev.Keysym.Mod
			}
			scanbuf = updateKey(scanbuf, int(ev.Keysym.Scancode), time, pressed)

		case *sdl.TextInputEvent:
			r, _ := utf8.DecodeRune(ev.Text[:])
			// if lastKbdEventIndex == -1 {
			if lastKbdEvent == nil {
				fmt.Printf("kbd input event warning %q\n", r)
				// lastKbdEventIndex = len(eventQueue)
				eventQueue = append(eventQueue, KeyboardEvent{})
				lastKbdEvent = &eventQueue[len(eventQueue)-1]
				// eventQueue[lastKbdEventIndex].Mod = 1
			}
			if r != utf8.RuneError {
				// eventQueue[lastKbdEventIndex].Rune = r
				lastKbdEvent.Rune = r
			}
		case *sdl.TextEditingEvent:
		} // switch ev := event.(type)
	} // for PollEvent
	for i := range eventQueue {
		callHandleEvent(on.obj, eventQueue[i])
	}
	scanbuf = packKey(scanbuf)
	eventQueue = eventQueue[:0]
}

func decodeRuneBuffer(buf []byte) string {
	s := ""
	for len(buf) > 0 {
		r, size := utf8.DecodeRune(buf)
		// fmt.Printf("%c %v\n", r, size)
		if r == '\x00' {
			break
		}
		s += string(r)
		buf = buf[size:]
	}
	return s
}

func updateKey(buf []tKeyState, scan int, time time.Duration, pressed bool) []tKeyState {
	for i := range buf {
		sc := buf[i].scan
		if sc < 0 || sc != scan {
			continue
		}
		if pressed {
			buf[i].timestamp = time
			buf[i].justPressed = true
		} else {
			buf[i].scan = -1
		}
		return buf
	}
	if pressed {
		buf = append(buf, tKeyState{scan: scan, timestamp: time, justPressed: true})
	}
	return buf
}

func packKey(buf []tKeyState) []tKeyState {
	o := 0
	for i := range buf {
		if buf[i].scan < 0 {
			continue
		}
		if i != o {
			buf[o] = buf[i]
		}
		o++
	}
	buf = buf[:o]
	return buf
}

// Scanbuf -
func Scanbuf() (string, bool) {
	if len(scanbuf) == 0 {
		return "", false
	}
	s := "--------------\n"
	for i := range scanbuf {
		s += fmt.Sprintf("%v, %v\n", scanbuf[i].scan, scanbuf[i].justPressed)
	}
	return s, true
}
