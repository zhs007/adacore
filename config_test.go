package adacore

import (
	"testing"

	adacoredef "github.com/zhs007/adacore/basedef"
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
	if err != adacoredef.ErrConfigNoAdaRenderServAddr {
		t.Fatalf("TestLoadConfig LoadConfig %v err is not ErrConfigNoAdaRenderServAddr", "./unittest/config002.yaml")
	}

	cfg, err = LoadConfig("./unittest/config003.yaml")
	if err != adacoredef.ErrConfigNoClientTokens {
		t.Fatalf("TestLoadConfig LoadConfig %v err is not ErrConfigNoClientTokens", "./unittest/config003.yaml")
	}

	cfg, err = LoadConfig("./unittest/config004.yaml")
	if err != adacoredef.ErrConfigNoAdaRenderToken {
		t.Fatalf("TestLoadConfig LoadConfig %v err is not ErrConfigNoAdaRenderToken", "./unittest/config004.yaml")
	}

	t.Logf("TestLoadConfig OK")
}
