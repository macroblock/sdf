package ui

import (
	"github.com/macroblock/sdf/pkg/theme"
)

type (
	// Button -
	Button struct {
		Panel
		pressed bool
	}
)

// NewButton -
func NewButton() *Button {
	o := Button{}
	o.self = &o
	return &o
}

// Draw -
func (o *Button) Draw() {
	// self := o.self
	r := o.Renderer()

	r.SetColor(theme.Button.Color())
	r.Clear()
}
