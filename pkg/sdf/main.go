package sdf

import (
	"fmt"
	"runtime"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

var (
	sdf    Type
	on     handler
	assets Assets
)

type (
	// Type -
	Type struct {
		err         error
		isRunning   bool
		deltaRender time.Duration
		deltaUpdate time.Duration
		window      *sdl.Window
		renderer    *sdl.Renderer
	}

	handler struct {
		x, y, w, h int32
		obj        interface{}
	}
)

func init() {
	runtime.LockOSThread()
	on.x = sdl.WINDOWPOS_UNDEFINED
	on.y = sdl.WINDOWPOS_UNDEFINED
	on.w = 640
	on.h = 480
}

// Run -
func Run(obj interface{}) error {
	if sdf.err != nil {
		return sdf.err
	}
	if sdf.isRunning {
		sdf.err = fmt.Errorf("sdf is already running")
	}

	if sdf.window != nil || sdf.renderer != nil {
		sdf.err = fmt.Errorf("sdf was not cleaned up correctly")
		return sdf.err
	}
	sdf = Type{}

	flags := uint32(0)

	assets = newAssets()

	err := sdl.Init(sdl.INIT_EVERYTHING)
	defer sdl.Quit()
	setError(err)

	window, renderer, err := sdl.CreateWindowAndRenderer(on.w, on.h, flags)
	setError(err)

	sdf.window = window
	sdf.renderer = renderer

	on.obj = obj

	if i, ok := obj.(iInit); ok {
		i.Init()
	}

	sdf.isRunning = true

	lastUpdate := time.Now()
	lastRender := time.Now()
	for Ok() && Running() {
		HandleEvents()
		sdf.deltaUpdate = time.Since(lastUpdate)
		lastUpdate = time.Now()
		Update()
		sdf.deltaRender = time.Since(lastRender)
		lastRender = time.Now()
		Render()
		sdf.renderer.Present()
	}

	CleanUp()
	return sdf.err
}

// CleanUp -
func CleanUp() {
	sdf.isRunning = false

	if i, ok := on.obj.(iCleanUp); ok {
		i.CleanUp()
	}

	err := error(nil)
	if sdf.renderer != nil {
		err = sdf.renderer.Destroy()
		sdf.renderer = nil
	}
	setError(err)
	if sdf.window != nil {
		err = sdf.window.Destroy()
		sdf.window = nil
	}
	setError(err)
}

// Running -
func Running() bool {
	return sdf.isRunning
}

// Ok -
func Ok() bool {
	return sdf.err == nil
}

// HasError -
func HasError() bool {
	return sdf.err != nil
}

// Error -
func Error() error {
	return sdf.err
}

// Warning -
func Warning() error {
	if sdf.err != nil {
		return sdf.err
	}
	if !sdf.isRunning {
		return fmt.Errorf("sdf is not running")
	}
	return nil
}

// HandleEvents -
func HandleEvents() {
	if !Ok() {
		return
	}
	event := sdl.PollEvent()
	switch ev := event.(type) {
	case *sdl.QuitEvent:
		_ = ev
		sdf.isRunning = false
	}
}

// Update -
func Update() {
	if !Ok() {
		return
	}
	if i, ok := on.obj.(iUpdate); ok {
		i.Update()
	}
}

// Render -
func Render() {
	if !Ok() {
		return
	}
	if i, ok := on.obj.(iRender); ok {
		i.Render()
	}
}

// DeltaUpdate -
func DeltaUpdate() time.Duration {
	return sdf.deltaUpdate
}

// DeltaRender -
func DeltaRender() time.Duration {
	return sdf.deltaRender
}
