package adacore

import "errors"

var (
	// ErrDuplicateFNInImageMap - Duplicate filename in ImageMap
	ErrDuplicateFNInImageMap = errors.New("Duplicate filename in ImageMap")
	// ErrDuplicateBNInImageMap - Duplicate buffname in ImageMap
	ErrDuplicateBNInImageMap = errors.New("Duplicate buffname in ImageMap")

	// ErrNilImageMap - ImageMap is nil
	ErrNilImageMap = errors.New("ImageMap is nil")

	// ErrInvalidImageFileType - invalid image file type
	ErrInvalidImageFileType = errors.New("invalid image file type")
)
