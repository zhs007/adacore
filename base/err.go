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

	// ErrServNoConfig - The config is invalid
	ErrServNoConfig = errors.New("The config is invalid")
	// ErrServInvalidToken - This token is invalid
	ErrServInvalidToken = errors.New("This token is invalid")
	// ErrServInvalidErrString - The err string is invalid
	ErrServInvalidErrString = errors.New("The err string is invalid")

	// ErrEmptyHTMLData - The HTMLData is empty
	ErrEmptyHTMLData = errors.New("The HTMLData is empty")

	// ErrAdaCoreClientNoServAddr - There is no ServAddr in AdaCoreClient
	ErrAdaCoreClientNoServAddr = errors.New("There is no ServAddr in AdaCoreClient")
	// ErrAdaCoreClientNoToken - There is no Token in AdaCoreClient
	ErrAdaCoreClientNoToken = errors.New("There is no Token in AdaCoreClient")
)
