package ui

import (
	"image/color"

	"github.com/macroblock/sdf/pkg/resources"
)

// ITheme -
type ITheme interface {
	BackgroundColor() color.Color
	ButtonColor() color.Color
	HyperlinkColor() color.Color
	TextColor() color.Color
	PlaceHolderColor() color.Color
	PrimaryColor() color.Color
	FocusColor() color.Color
	ScrollBarColor() color.Color

	TextSize() int
	TextFont() resources.IResource
	TextBoldFont() resources.IResource
	TextItalicFont() resources.IResource
	TextBoldItalicFont() resources.IResource
	TextMonospaceFont() resources.IResource

	Padding() int
	IconInlineSize() int
	ScrollBarSize() int
}
