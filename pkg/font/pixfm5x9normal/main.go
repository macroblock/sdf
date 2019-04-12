package pixfm5x9normal

import "github.com/macroblock/sdf/pkg/font"

// Font -
var Font = font.PixelFontSettings{
	FileName:        "../../assets/fonts/font-5x9-(16x7).png",
	MinRune:         ' ',
	MaxRune:         '~',
	UnavailableRune: '~' + 2,
	// TilesX:          16,
	// TilesY:          7,
	GlyphW:   5,
	GlyphH:   9,
	BearingX: 0,
	BearingY: 7,
	AdvanceX: 6,
	AdvanceY: 11,
}
