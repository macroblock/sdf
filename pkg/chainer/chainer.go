package chainer

type (
	// Chainer -
	Chainer struct {
		entryType
		current *entryType
		history string
		actions map[string]IAction
	}

	entryType struct {
		entries map[BinaryKey]*entryType
		action  IAction
	}

	// BinaryKey -
	BinaryKey [4]int
)

// MakeBinaryKey -
func MakeBinaryKey(v3, v2, v1, v0 int) BinaryKey {
	return BinaryKey{v3, v2, v1, v0}
}
