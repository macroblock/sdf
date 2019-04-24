package theme

import (
	"image/color"
)

type (
	// ITheme -
	ITheme interface {
		Palette() *palette

		// TextSize() int
		// TextFont() resources.IResource
		// TextBoldFont() resources.IResource
		// TextItalicFont() resources.IResource
		// TextBoldItalicFont() resources.IResource
		// TextMonospaceFont() resources.IResource

		// Padding() int
		// IconInlineSize() int
		// ScrollBarSize() int
	}

	// PaletteIndex -
	PaletteIndex int

	palette [paletteLen]color.Color
)

// Color -
func (o PaletteIndex) Color() color.Color { return current.Palette()[o] }

var (
	current ITheme = &builtinDarkTheme
)

// -
const (
	Background PaletteIndex = iota
	Button
	Hyperlink
	Text
	PlaceHolder
	Primary
	ScrollBar
	paletteLen

	FocusColor = Primary
)

var builtinDarkTheme = palette{
	Background:  color.RGBA{0x42, 0x42, 0x42, 0xff},
	Button:      color.RGBA{0x21, 0x21, 0x21, 0xff},
	Text:        color.RGBA{0x99, 0x99, 0xff, 0xff},
	Hyperlink:   color.RGBA{0xff, 0xff, 0xff, 0xff},
	PlaceHolder: color.RGBA{0xb2, 0xb2, 0xb2, 0xff},
	Primary:     color.RGBA{0x1a, 0x23, 0x7e, 0xff},
	ScrollBar:   color.RGBA{0x0, 0x0, 0x0, 0x99},
}

func (o *palette) Palette() *palette {
	if o == nil {
		return &builtinDarkTheme
	}
	return o
}

// Current -
func Current() ITheme {
	return current
}
