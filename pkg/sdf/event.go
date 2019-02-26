package sdf

type (
	// Event -
	Event struct {
		Align uint8
		Type  uint8
		Mod   uint16
		Rune  rune
	}
)

// BinaryKey -
func (o Event) BinaryKey() int64 {
	return (((int64(o.Align)<<8)|int64(o.Type))<<16|int64(o.Mod))<<32 | int64(o.Rune)
}
