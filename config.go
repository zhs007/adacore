package adacore

import (
	"io/ioutil"
	"os"

	"go.uber.org/zap/zapcore"
	"gopkg.in/yaml.v2"

	adacoredef "github.com/zhs007/adacore/basedef"
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
		return adacoredef.ErrConfigNoAdaRenderServAddr
	}

	if cfg.AdaRenderToken == "" {
		return adacoredef.ErrConfigNoAdaRenderToken
	}

	if len(cfg.ClientTokens) == 0 {
		return adacoredef.ErrConfigNoClientTokens
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
