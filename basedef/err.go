package adacoredef

import "errors"

var (
	// ErrConfigNoAdaRenderServAddr - There is no AdaRenderServAddr in the configuration file
	ErrConfigNoAdaRenderServAddr = errors.New("There is no AdaRenderServAddr in the configuration file")
)
