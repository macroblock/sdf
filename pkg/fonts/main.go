package fonts

// PixelFontSettings -
type PixelFontSettings struct {
	FileName           string
	MinRune, MaxRune   rune
	UnavailableRune    rune
	TilesX, TilesY     int
	GlyphW, GlyphH     int
	AdvanceX, AdvanceY int
	BearingX, BearingY int
}
