package geom

// Rect2i -
type Rect2i struct {
	X, Y, W, H int
}

// InitRect2i -
func InitRect2i(x, y, w, h int) Rect2i {
	return Rect2i{x, y, w, h}
}

// Normalize -
func (o Rect2i) Normalize() Rect2i {
	if o.W < 0 {
		o.W = -o.W
		o.X -= o.W
	}
	if o.H < 0 {
		o.H = -o.H
		o.Y -= o.H
	}
	return o
}

// Add -
func (o Rect2i) Add(pt Point2i) Rect2i {
	o.X += pt.X
	o.Y += pt.Y
	return o
}

// Mul -
func (o Rect2i) Mul(kxy Point2i) Rect2i {
	o.X *= kxy.X
	o.Y *= kxy.Y
	o.W *= kxy.X
	o.H *= kxy.Y
	return o
}

// AddInt -
func (o Rect2i) AddInt(x, y int) Rect2i {
	o.X += x
	o.Y += y
	return o
}

// MulInt -
func (o Rect2i) MulInt(kx, ky int) Rect2i {
	o.X *= kx
	o.Y *= ky
	o.W *= kx
	o.H *= ky
	return o
}
