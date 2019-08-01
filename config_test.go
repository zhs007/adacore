package adacore

import (
	"testing"

	adacorebase "github.com/zhs007/adacore/base"
)

func TestLoadConfig(t *testing.T) {
	cfg, err := LoadConfig("./unittest/config001.yaml")
	if err != nil {
		t.Fatalf("TestLoadConfig LoadConfig %v err is %v", "./unittest/config001.yaml", err)
	}

	if cfg == nil {
		t.Fatalf("TestLoadConfig LoadConfig %v cfg is nil", "./unittest/config001.yaml")
	}

	cfg, err = LoadConfig("./unittest/config002.yaml")
	if err != adacorebase.ErrConfigNoAdaRenderServAddr {
		t.Fatalf("TestLoadConfig LoadConfig %v err is not ErrConfigNoAdaRenderServAddr", "./unittest/config002.yaml")
	}

	cfg, err = LoadConfig("./unittest/config003.yaml")
	if err != adacorebase.ErrConfigNoClientTokens {
		t.Fatalf("TestLoadConfig LoadConfig %v err is not ErrConfigNoClientTokens", "./unittest/config003.yaml")
	}

	cfg, err = LoadConfig("./unittest/config004.yaml")
	if err != adacorebase.ErrConfigNoAdaRenderToken {
		t.Fatalf("TestLoadConfig LoadConfig %v err is not ErrConfigNoAdaRenderToken", "./unittest/config004.yaml")
	}

	cfg, err = LoadConfig("./unittest/config005.yaml")
	if err != adacorebase.ErrConfigNoFilePath {
		t.Fatalf("TestLoadConfig LoadConfig %v err is not ErrConfigNoFilePath", "./unittest/config005.yaml")
	}

	t.Logf("TestLoadConfig OK")
}
