package sdf

import (
	"fmt"
	"runtime"
	"time"

	"github.com/macroblock/sdf/pkg/fonts/pixfm5x9normal"

	"github.com/macroblock/sdf/pkg/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

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
		err              error
		isRunning        bool
		fps              float64
		deltaRender      time.Duration
		deltaUpdate      time.Duration
		window           *sdl.Window
		renderer         *gfx.Renderer
		curTilePath      string
		curAnimationPath string
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
	sdf = Type{curTilePath: "/", curAnimationPath: "/"}

	flags := uint32(0)

	assets = newAssets()

	err := sdl.Init(sdl.INIT_EVERYTHING)
	defer sdl.Quit()
	setError(err)

	sdf.window, err = sdl.CreateWindow("test", on.x, on.y, on.w, on.h, flags)
	if err != nil {
		setError(err)
		return err
	}

	flags = sdl.RENDERER_ACCELERATED | sdl.RENDERER_PRESENTVSYNC
	renderer, err := sdl.CreateRenderer(sdf.window, -1, flags)
	setError(err)

	sdf.renderer = gfx.NewRenderer(renderer)
	defaultFont := CreatePixelFont(pixfm5x9normal.Font)
	sdf.renderer.SetDefaultFont(defaultFont)

	on.obj = obj

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

// SetScale -
func SetScale(x, y float64) {
	if !Ok() {
		return
	}
	sdf.renderer.SetScale(x, y)
}

// Renderer -
func Renderer() *gfx.Renderer {
	return sdf.renderer
}

// Size -
func Size() (int, int) {
	if !Ok() {
		return -1, -1
	}
	// w, h, err := sdf.renderer.GetOutputSize()
	// if err != nil {
	// 	setError(err)
	// 	return -1, -1
	// }
	w, h := sdf.renderer.Size()
	return w, h
}
