package sdf

import (
	"fmt"
	"path/filepath"
)

type (
	// Assets -
	Assets struct {
		byURI map[string]*Texture
	}

	// IResource -
	IResource interface {
		Load(uri string) error
		Loaded() bool
	}
)

func newAssets() Assets {
	return Assets{byURI: map[string]*Texture{}}
}

// LoadResource -
func (o *Assets) loadResource(path string) (*Texture, error) {
	uri, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	res, ok := o.byURI[uri]
	if ok {
		res.IncRef()
		return res, nil
	}

	res, err = loadTexture(uri)
	if err != nil {
		return nil, err
	}
	res.IncRef()
	o.byURI[uri] = res
	return res, nil

}

// UnloadResource -
func (o *Assets) unloadResource(res *Texture) error {
	if res == nil {
		return fmt.Errorf("resource is nil")
	}
	res.DecRef()
	uri := res.GetURI()
	refs := res.NumRefs()

	if refs > 0 {
		return nil
	}
	if refs < 0 {
		return fmt.Errorf("reference counter of a resource %q is %v", uri, refs)
	}

	uri = res.GetURI()
	_, ok := o.byURI[uri]
	if !ok {
		return fmt.Errorf("resource %q was not found", uri)
	}
	unloadTexture(res)
	return nil
}
