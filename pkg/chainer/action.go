package chainer

type (
	// Action -
	Action struct {
		name        string
		keychain    string
		description string
		handler     ActionHandler
	}

	// ActionHandler -
	ActionHandler func(keychain string) bool

	// IAction -
	IAction interface {
		Name() string
		Keychain() string
		Description() string
		Handler() ActionHandler
		BinaryKey() ([]BinaryKey, error)
	}

	// Builder -
	Builder struct{}
)

// NewAction -
func NewAction(name, keychain, desc string, handler ActionHandler) IAction {
	return &Action{
		name:        name,
		keychain:    keychain,
		description: desc,
		handler:     handler,
	}
}

// Keychain -
func (o *Action) Keychain() string { return o.keychain }

// Name -
func (o *Action) Name() string { return o.name }

// Description -
func (o *Action) Description() string { return o.description }

// Handler -
func (o *Action) Handler() ActionHandler { return o.handler }

// BinaryKey -
func (o *Action) BinaryKey() ([]BinaryKey, error) { return keychainToBinary(o.keychain) }
