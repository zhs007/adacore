package adarenderclient

import (
	"context"
	"testing"

	adarender "github.com/zhs007/adacore/adarenderpb"
)

func TestRenderClient(t *testing.T) {
	client := NewClient("47.91.209.141:7052", "RVhVrt13P6i5xCrL5Fc3GcuHC03kaunA")

	mddata := &adarender.MarkdownData{}

	mddata.StrData = `# Ada Render Sample

	This is a ` + "``markdown``" + ` file.
	`

	mddata.TemplateName = "default"

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	htmldata, err := client.Render(ctx, mddata)
	if err != nil {
		t.Fatalf("TestRenderClient Render err is %v", err)
	}

	if htmldata == nil {
		t.Fatalf("TestRenderClient Render non HTMLData")
	}

	t.Logf("TestRenderClient OK")
}
