package sdf

import "github.com/veandco/go-sdl2/sdl"

type (
	iInit interface {
		Init()
	}
	iCleanUp interface {
		CleanUp()
	}
	iUpdate interface {
		Update()
	}
	iRender interface {
		Render()
	}
	iHandleEvents interface {
		HandleEvent(ev sdl.Event)
	}

	// IElem -
	IElem interface {
		// Update(delta time.Duration) bool
		Copy(x, y int)
	}
)

func callInit(i interface{}) {
	if i, ok := i.(iInit); ok {
		i.Init()
	}
}

func callCleanUp(i interface{}) {
	if i, ok := i.(iCleanUp); ok {
		i.CleanUp()
	}
}

func callUpdate(i interface{}) {
	if i, ok := i.(iUpdate); ok {
		i.Update()
	}
}

func callRender(i interface{}) {
	if i, ok := i.(iRender); ok {
		i.Render()
	}
}

func callHandleEvent(i interface{}, ev sdl.Event) {
	if i, ok := i.(iHandleEvents); ok {
		i.HandleEvent(ev)
	}
}
