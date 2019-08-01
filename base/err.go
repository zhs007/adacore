package adacorebase

import "errors"

var (
	// ErrConfigNoAdaRenderServAddr - There is no AdaRenderServAddr in the configuration file
	ErrConfigNoAdaRenderServAddr = errors.New("There is no AdaRenderServAddr in the configuration file")
	// ErrConfigNoClientTokens - There is no ClientTokens in the configuration file
	ErrConfigNoClientTokens = errors.New("There is no ClientTokens in the configuration file")
	// ErrConfigNoAdaRenderToken - There is no AdaRenderToken in the configuration file
	ErrConfigNoAdaRenderToken = errors.New("There is no AdaRenderToken in the configuration file")
	// ErrConfigNoFilePath - There is no FilePath in the configuration file
	ErrConfigNoFilePath = errors.New("There is no FilePath in the configuration file")
	// ErrConfigNoBindAddr - There is no BindAddr in the configuration file
	ErrConfigNoBindAddr = errors.New("There is no BindAddr in the configuration file")

	// ErrAdaRenderClientNoServAddr - There is no ServAddr in AdaRenderClient
	ErrAdaRenderClientNoServAddr = errors.New("There is no ServAddr in AdaRenderClient")
	// ErrAdaRenderClientNoToken - There is no Token in AdaRenderClient
	ErrAdaRenderClientNoToken = errors.New("There is no Token in AdaRenderClient")

	// ErrAdaCoreDBInvalidHashName - There is invalid hashname in AdaCoreDB
	ErrAdaCoreDBInvalidHashName = errors.New("There is invalid hashname in AdaCoreDB")
	// ErrAdaCoreDBInvalidResList - There is invalid reslist in AdaCoreDB
	ErrAdaCoreDBInvalidResList = errors.New("There is invalid reslist in AdaCoreDB")
)
