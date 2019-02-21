package sdf

import "time"

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
		HandleEvents()
	}

	// IElem -
	IElem interface {
		Copy(x, y int, delta time.Duration)
	}
)
