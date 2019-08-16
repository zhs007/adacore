package adacore

import (
	"io/ioutil"
	"os"

	"go.uber.org/zap/zapcore"
	"gopkg.in/yaml.v2"

	adacorebase "github.com/zhs007/adacore/base"
)

// Config - config
type Config struct {

	//------------------------------------------------------------------
	// adarender configuration

	// AdaRenderServAddr - Ada render service address
	AdaRenderServAddr string
	// AdaRenderToken - This is a valid adarenderserv token
	AdaRenderToken string

	//------------------------------------------------------------------
	// adanode service configuration

	// ClientTokens - There are the valid clienttokens for this node
	ClientTokens []string
	// MaxExpireTime - max expire time in seconds
	MaxExpireTime int32
	// IsAllowTemplateData - Whether to allow templatedata
	IsAllowTemplateData bool
	// Templates - This is all the templates available for this role.
	Templates []string
	// ResNums - This is the amount of resources available for this role
	ResNums int32

	// FilePath - Output file path
	FilePath string
	// BindAddr - bind addr
	BindAddr string
	// BaseURL - base URL
	BaseURL string
	// TemplatesPath - templates file path
	// Deprecated: The configuration of the template path is no longer needed.
	TemplatesPath string

	//------------------------------------------------------------------
	// logger configuration

	Log struct {
		// LogPath - log path
		LogPath string
		// LogLevel - log level, it can be debug, info, warn, error
		LogLevel string
		// LogConsole - it can be output to console
		LogConsole bool
	}
}

func getLogLevel(str string) zapcore.Level {
	switch str {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	default:
		return zapcore.ErrorLevel
	}
}

func checkConfig(cfg *Config) error {
	if cfg.AdaRenderServAddr == "" {
		return adacorebase.ErrConfigNoAdaRenderServAddr
	}

	if cfg.AdaRenderToken == "" {
		return adacorebase.ErrConfigNoAdaRenderToken
	}

	if len(cfg.ClientTokens) == 0 {
		return adacorebase.ErrConfigNoClientTokens
	}

	if cfg.FilePath == "" {
		return adacorebase.ErrConfigNoFilePath
	}

	if cfg.BindAddr == "" {
		return adacorebase.ErrConfigNoBindAddr
	}

	if cfg.MaxExpireTime <= 0 {
		cfg.MaxExpireTime = adacorebase.DefaultMaxExpireTime
	}

	if len(cfg.Templates) == 0 {
		cfg.Templates = append(cfg.Templates, adacorebase.DefaultTemplate)
	}

	if cfg.ResNums < 0 {
		cfg.ResNums = adacorebase.DefaultResNums
	}

	if cfg.BaseURL == "" {
		cfg.BaseURL = adacorebase.DefaultBaseURL
	}

	if cfg.TemplatesPath == "" {
		cfg.TemplatesPath = adacorebase.DefaultTemplatesPath
	}

	return nil
}

// LoadConfig - load config
func LoadConfig(filename string) (*Config, error) {
	fi, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}

	err = yaml.Unmarshal(fd, cfg)
	if err != nil {
		return nil, err
	}

	err = checkConfig(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

// HasToken - has token
func (cfg *Config) HasToken(token string) bool {
	for _, v := range cfg.ClientTokens {
		if v == token {
			return true
		}
	}

	return false
}

// InitLogger - init logger
func InitLogger(cfg *Config) {
	adacorebase.InitLogger(getLogLevel(cfg.Log.LogLevel), cfg.Log.LogConsole, cfg.Log.LogPath)
}
