package main

import (
	"fmt"
	"context"

	"github.com/zhs007/adacore/adarenderclient"
	"github.com/zhs007/adacore/adarenderpb"
)

func main() {
	client := adarenderclient.NewClient("127.0.0.1:7052", "RVhVrt13P6i5xCrL5Fc3GcuHC03kaunA")
	// client := adarenderclient.NewClient("47.91.209.141:7052", "RVhVrt13P6i5xCrL5Fc3GcuHC03kaunA")

	mddata := &adarender.MarkdownData{}

	mddata.StrData = `# Ada Render Sample

	This is a ` + "``markdown``" + ` file.
	`

	mddata.TemplateName = "default"

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	htmldata, err := client.Render(ctx, mddata)
	if err != nil {
		fmt.Print("TestRenderClient Render err is %v", err)
	}

	if htmldata == nil {
		fmt.Print("TestRenderClient Render non HTMLData")
	}

	// fmt.Print("TestRenderClient OK")
}