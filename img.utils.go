package adacore

import (
	"image"
	"os"
	"path/filepath"
	"strings"

	"image/gif"
	"image/jpeg"
	"image/png"

	"github.com/disintegration/imaging"
)

// SaveImageFile - save image file
func SaveImageFile(fn string, img image.Image) error {
	f, err := os.Create(fn)
	if err != nil {
		return err
	}
	defer f.Close()

	_, rfn := filepath.Split(strings.ToLower(fn))
	arr := strings.Split(rfn, ".")
	if len(arr) > 1 {
		if arr[len(arr)-1] == "jpg" || arr[len(arr)-1] == "jpeg" {
			return jpeg.Encode(f, img, nil)
		} else if arr[len(arr)-1] == "gif" {
			return gif.Encode(f, img, nil)
		} else if arr[len(arr)-1] == "png" {
			return png.Encode(f, img)
		}
	}

	return ErrInvalidImageFileType
}

// LoadImageFile - load image file
func LoadImageFile(fn string) (image.Image, error) {
	reader, err := os.Open(fn)
	if err != nil {
		return nil, err
	}

	defer reader.Close()

	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}

	return img, nil
}

// ResizeImage - resize image file
func ResizeImage(destfn string, w int, h int, srcfn string, filter imaging.ResampleFilter) error {
	img, err := LoadImageFile(srcfn)
	if err != nil {
		return err
	}

	nimg := imaging.Resize(img, w, h, filter)
	// nimg := imaging.Fill(img, w, h, imaging.Center, imaging.NearestNeighbor)

	return SaveImageFile(destfn, nimg)
}

// FitImage - fit image file
func FitImage(destfn string, w int, h int, srcfn string, filter imaging.ResampleFilter) error {
	img, err := LoadImageFile(srcfn)
	if err != nil {
		return err
	}

	nimg := imaging.Fit(img, w, h, imaging.NearestNeighbor)
	// nimg := imaging.Fill(img, w, h, imaging.Center, imaging.NearestNeighbor)

	return SaveImageFile(destfn, nimg)
}

// FillImage - fill image file
func FillImage(destfn string, w int, h int, srcfn string, anchor imaging.Anchor, filter imaging.ResampleFilter) error {
	img, err := LoadImageFile(srcfn)
	if err != nil {
		return err
	}

	// nimg := imaging.Resize(img, w, h, imaging.NearestNeighbor)
	nimg := imaging.Fill(img, w, h, anchor, filter)
	// nimg := imaging.Fill(img, w, h, imaging.Center, imaging.NearestNeighbor)

	return SaveImageFile(destfn, nimg)
}
