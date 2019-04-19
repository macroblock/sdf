package resources

// IResource -
type IResource interface {
	Name() string
	Content() []byte
}

// StaticResource -
type StaticResource struct {
	StaticName    string
	StaticContent []byte
}

// Name -
func (r *StaticResource) Name() string {
	return r.StaticName
}

// Content -
func (r *StaticResource) Content() []byte {
	return r.StaticContent
}

// NewStaticResource -
func NewStaticResource(name string, content []byte) *StaticResource {
	return &StaticResource{
		StaticName:    name,
		StaticContent: content,
	}
}
