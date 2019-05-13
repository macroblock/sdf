package theme

import (
	"bytes"
	"image/color"
	"io/ioutil"

	"github.com/golang/freetype/truetype"
	"github.com/macroblock/sdf/pkg/gfx"
	"github.com/macroblock/sdf/pkg/unifont"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
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

var fontFace font.Face

// AcquireFontFace -
func AcquireFontFace(renderer *gfx.Renderer) font.Face {
	if fontFace != nil {
		return fontFace
	}
	r := bytes.NewReader(goregular.TTF)
	data, err := ioutil.ReadAll(r)
	if err != nil {
		// return nil, err
		panic("!!!")
	}

	ttf, err := truetype.Parse(data)
	if err != nil {
		// return nil, err
		panic("???")
	}

	face, err := unifont.NewHWFace(renderer, ttf, &truetype.Options{
		Size: float64(20),
		// Hinting:    font.HintingFull,
		SubPixelsX: 1,
		SubPixelsY: 1,
	})
	if err != nil {
		panic("!?!? " + err.Error())
	}
	return face
}

// ReleaseFontFace -
func ReleaseFontFace() {}
