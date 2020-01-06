package adacore

import (
	"io/ioutil"
	"path/filepath"
)

// ImageMap - image mapping
type ImageMap struct {
	MapImgs map[string][]byte
}

// NewImageMap - new ImageMap
func NewImageMap() *ImageMap {
	return &ImageMap{
		MapImgs: make(map[string][]byte),
	}
}

// AddImage - add a image
func (im *ImageMap) AddImage(fn string, fullfn bool) (string, error) {
	var key string

	if !fullfn {
		_, cfn := filepath.Split(fn)
		key = cfn
	} else {
		key = fn
	}

	_, isok := im.MapImgs[key]
	if isok {
		return key, ErrDuplicateFNInImageMap
	}

	buf, err := ioutil.ReadFile(fn)
	if err != nil {
		return "", err
	}

	im.MapImgs[key] = buf

	return key, nil
}

// AddImageBuff - add a image buffer
func (im *ImageMap) AddImageBuff(name string, buf []byte) error {
	_, isok := im.MapImgs[name]
	if isok {
		return ErrDuplicateBNInImageMap
	}

	im.MapImgs[name] = buf

	return nil
}
