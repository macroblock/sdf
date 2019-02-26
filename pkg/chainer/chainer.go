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
	BinaryKey int64
)
