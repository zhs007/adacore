package adacore

import (
	"io/ioutil"
	"os"

	"go.uber.org/zap/zapcore"
	yaml "gopkg.in/yaml.v2"

	adacoredef "github.com/zhs007/adacore/basedef"
)

// Config - config
type Config struct {

	//------------------------------------------------------------------
	// base configuration

	AdaRenderServAddr string
	AdaRenderToken    string
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
