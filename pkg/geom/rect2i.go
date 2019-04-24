package geom

import "github.com/macroblock/sdf/pkg/misc"

// Rect2i -
type Rect2i struct {
	A Point2i
	B Point2i
}

// InitRect2i -
func InitRect2i(x, y, w, h int) Rect2i {
	return Rect2i{Point2i{x, y}, Point2i{x + w, y + h}}
}

// InitRect2iAbs -
func InitRect2iAbs(x1, y1, x2, y2 int) Rect2i {
	return Rect2i{Point2i{x1, y1}, Point2i{x2, y2}}
}

// Size -
func (o Rect2i) Size() Point2i {
	return InitPoint2i(o.B.X-o.A.X, o.B.Y-o.A.Y)
}

// W -
func (o Rect2i) W() int {
	return o.B.X - o.A.X
}

// H -
func (o Rect2i) H() int {
	return o.B.Y - o.A.Y
}

// Canon -
func (o Rect2i) Canon() Rect2i {
	if o.B.X < o.A.X {
		o.A.X, o.B.X = o.B.X, o.A.X
	}
	if o.B.Y < o.B.Y {
		o.A.Y, o.B.Y = o.B.Y, o.A.Y
	}
	return o
}

// SetPos -
func (o Rect2i) SetPos(pos Point2i) Rect2i {
	dX := pos.X - o.A.X
	dY := pos.Y - o.A.Y
	return o.AddInt(dX, dY)
}

// Add -
func (o Rect2i) Add(pt Point2i) Rect2i {
	o.A.X += pt.X
	o.A.Y += pt.Y
	o.B.X += pt.X
	o.B.Y += pt.Y
	return o
}

// Sub -
func (o Rect2i) Sub(pt Point2i) Rect2i {
	o.A.X -= pt.X
	o.A.Y -= pt.Y
	o.B.X -= pt.X
	o.B.Y -= pt.Y
	return o
}

// Mul -
func (o Rect2i) Mul(kxy Point2i) Rect2i {
	o.A.X *= kxy.X
	o.A.Y *= kxy.Y
	o.B.X *= kxy.X
	o.B.Y *= kxy.Y
	return o
}

// AddInt -
func (o Rect2i) AddInt(x, y int) Rect2i {
	o.A.X += x
	o.A.Y += y
	o.B.X += x
	o.B.Y += y
	return o
}

// SubInt -
func (o Rect2i) SubInt(x, y int) Rect2i {
	o.A.X -= x
	o.A.Y -= y
	o.B.X -= x
	o.B.Y -= y
	return o
}

// MulInt -
func (o Rect2i) MulInt(kx, ky int) Rect2i {
	o.A.X *= kx
	o.A.Y *= ky
	o.B.X *= kx
	o.B.Y *= ky
	return o
}

// Intersect -
func (o Rect2i) Intersect(r Rect2i) Rect2i {
	o.A.X = misc.MaxInt(o.A.X, r.A.X)
	o.A.Y = misc.MaxInt(o.A.Y, r.A.Y)
	o.B.X = misc.MinInt(o.B.X, r.B.X)
	o.B.Y = misc.MinInt(o.B.Y, r.B.Y)
	return o
}

// Consists -
func (o Rect2i) Consists(p Point2i) bool {
	if o.A.X <= p.X && p.X < o.B.X && o.A.Y <= p.Y && p.Y < o.B.Y {
		return true
	}
	return false
}

// ConsistsInt -
func (o Rect2i) ConsistsInt(x, y int) bool {
	if o.A.X <= x && x < o.B.X && o.A.Y <= y && y < o.B.Y {
		return true
	}
	return false
}
