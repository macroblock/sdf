package general

// RefCounter -
type RefCounter struct {
	ref int
}

// IncRef -
func (o *RefCounter) IncRef() { o.ref++ }

// DecRef -
func (o *RefCounter) DecRef() { o.ref-- }

// NumRefs -
func (o *RefCounter) NumRefs() int { return o.ref }

// URI -
type URI struct {
	uri string
}

// SetURI -
func (o *URI) SetURI(uri string) {
	o.uri = uri
}

// GetURI -
func (o *URI) GetURI() string {
	return o.uri
}
