package sdf

import (
	"fmt"
	"path/filepath"
)

type (
	// Assets -
	Assets struct {
		byURI  map[string]*Texture
		sheets map[string]*TileSheet
		tiles  map[string]*Tile
		anims  map[string]*Animation

		currentTileSheet *TileSheet
	}

	// IResource -
	IResource interface {
		Load(uri string) error
		Loaded() bool
	}
)

func newAssets() Assets {
	return Assets{
		byURI:  map[string]*Texture{},
		sheets: map[string]*TileSheet{},
		tiles:  map[string]*Tile{},
		anims:  map[string]*Animation{},
	}
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

func (o *Assets) listURIs() []string {
	list := []string{}
	for key := range o.byURI {
		list = append(list, key)
	}
	return list
}

func (o *Assets) listTileSheets() []string {
	list := []string{}
	for key := range o.sheets {
		list = append(list, key)
	}
	return list
}

func (o *Assets) listTiles() []string {
	list := []string{}
	for key := range o.tiles {
		list = append(list, key)
	}
	return list
}
func (o *Assets) listAnimations() []string {
	list := []string{}
	for key := range o.anims {
		list = append(list, key)
	}
	return list
}

// ListAssets -
func ListAssets() []string {
	list := assets.listURIs()
	for i := range list {
		list[i] = "  uri: " + list[i]
	}
	ret := list
	list = assets.listTileSheets()
	for i := range list {
		list[i] = "sheet: " + list[i]
	}
	ret = append(ret, list...)
	list = assets.listTiles()
	for i := range list {
		list[i] = " tile: " + list[i]
	}
	ret = append(ret, list...)
	list = assets.listAnimations()
	for i := range list {
		list[i] = " anim: " + list[i]
	}
	ret = append(ret, list...)
	return ret
}
