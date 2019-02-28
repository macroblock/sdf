package sdf

import (
	"fmt"
	"runtime"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const constNameSeparator = "/"

var (
	sdf          Type
	on           handler
	assets       Assets
	programStart = time.Now()
	fixedTime    = time.Since(programStart)
)

type (
	// Type -
	Type struct {
		err         error
		isRunning   bool
		fps         float64
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

	sdf.window, err = sdl.CreateWindow("test", on.x, on.y, on.w, on.h, flags)
	if err == nil {
		flags = sdl.RENDERER_ACCELERATED | sdl.RENDERER_PRESENTVSYNC
		sdf.renderer, err = sdl.CreateRenderer(sdf.window, -1, flags)
	}
	setError(err)

	// sdf.window = window
	// sdf.renderer = renderer

	on.obj = obj

	// if i, ok := obj.(iInit); ok {
	// 	i.Init()
	// }
	callInit(obj)

	sdf.isRunning = true

	lastUpdate := time.Now()
	lastRender := time.Now()
	for Ok() && Running() {
		fixedTime = time.Since(programStart)
		HandleEvents()
		sdf.deltaUpdate = time.Since(lastUpdate)
		lastUpdate = time.Now()
		Update()
		sdf.deltaRender = time.Since(lastRender)
		lastRender = time.Now()
		Render()
		sdf.renderer.Present()
		// sdl.Delay(1000 / 60)
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
func Running() bool { return sdf.isRunning }

// Quit -
func Quit() { sdf.isRunning = false }

// Ok -
func Ok() bool { return sdf.err == nil }

// HasError -
func HasError() bool { return sdf.err != nil }

// Error -
func Error() error { return sdf.err }

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
	processInput()
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

// FPS -
func FPS() float64 {
	return sdf.fps
}
