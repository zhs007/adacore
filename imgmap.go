package adacore

import "io/ioutil"

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
func (im *ImageMap) AddImage(fn string) error {
	_, isok := im.MapImgs[fn]
	if isok {
		return ErrDuplicateFNInImageMap
	}

	buf, err := ioutil.ReadFile(fn)
	if err != nil {
		return err
	}

	im.MapImgs[fn] = buf

	return nil
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
