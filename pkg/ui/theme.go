package ui

import (
	"image/color"
)

type (
	// ITheme -
	ITheme interface {
		Color(ColorIndex) color.Color

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

	// ColorIndex -
	ColorIndex int

	builtinTheme [maxColors]color.Color
)

// -
const (
	BackgroundColor ColorIndex = iota
	ButtonColor
	HyperlinkColor
	TextColor
	PlaceHolderColor
	PrimaryColor
	FocusColor
	ScrollBarColor
	maxColors
)

var builtinDarkTheme = [maxColors]color.Color{
	// Dark theme -
	color.RGBA{0x42, 0x42, 0x42, 0xff},
	color.RGBA{0x21, 0x21, 0x21, 0xff},
	color.RGBA{0xff, 0xff, 0xff, 0xff},
	color.RGBA{0x99, 0x99, 0xff, 0xff},
	color.RGBA{0xb2, 0xb2, 0xb2, 0xff},
	color.RGBA{0x1a, 0x23, 0x7e, 0xff},
	color.RGBA{0x0, 0x0, 0x0, 0x99},
}
