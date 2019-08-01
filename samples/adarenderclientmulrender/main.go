package main

import (
	"context"
	"fmt"

	"github.com/zhs007/adacore/adarenderclient"
	adarender "github.com/zhs007/adacore/adarenderpb"
)

func render(ctx context.Context, client *adarenderclient.Client, mddata *adarender.MarkdownData, index int, endChan chan int) {
	htmldata, err := client.Render(ctx, mddata)
	if err != nil {
		fmt.Printf("TestRenderClient Render err is %v", err)
	}

	if htmldata == nil {
		fmt.Print("TestRenderClient Render non HTMLData")
	}

	fmt.Printf("render %v done.", index)

	endChan <- index
}

func main() {
	client := adarenderclient.NewClient("127.0.0.1:7052", "RVhVrt13P6i5xCrL5Fc3GcuHC03kaunA")
	// client := adarenderclient.NewClient("47.91.209.141:7052", "RVhVrt13P6i5xCrL5Fc3GcuHC03kaunA")

	mddata := &adarender.MarkdownData{}

	mddata.StrData = `# Ada Render Sample

	This is a ` + "``markdown``" + ` file.
	`

	mddata.TemplateName = "default"

	c := make(chan int)

	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()

	for i := 0; i < 100; i++ {
		// ctx, _ := context.WithCancel(context.Background())
		go render(context.Background(), client, mddata, i, c)
	}

	lastnums := 100

	for {
		<-c

		lastnums--
		if lastnums == 0 {
			return
		}
	}

	// fmt.Print("TestRenderClient OK")
}
