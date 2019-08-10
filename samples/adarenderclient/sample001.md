# Ada Render Sample

This is a ``markdown`` file.

``` go
package main

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/zhs007/adacore"
	adacorepb "github.com/zhs007/adacore/proto"
)

func genMarkdown() string {
	km, err := adacore.LoadKeywordMappingList("./keywordmapping.yaml")
	if err != nil {
		fmt.Printf("load keywordmapping error %v", err)
	}

	md := adacore.NewMakrdown("Ada Core")

	md.AppendParagraph("This is a Markdown API for Ada.")
	md.AppendParagraph("This libraray is write by Zerro.\nThis is a multi-line text.")

	md.AppendTable([]string{"head0", "head1", "head2"}, [][]string{
		[]string{"text0_0", "text1_0", "text2_0"},
		[]string{"text0_1", "text1_1", "text2_1"}})

	fd, err := ioutil.ReadFile("./main.go")
	if err != nil {
		return ""
	}

	md.AppendCode(string(fd), "golang")

	return md.GetMarkdownString(km)
}

func startClient(cfg *adacore.Config) error {
	client := adacore.NewClient("47.91.209.141:7201", "x7sSGGHgmKwUMoa5S4VZlr9bUF2lCCzF")

	reply, err := client.BuildWithMarkdown(context.Background(), &adacorepb.MarkdownData{
		StrData:      genMarkdown(),
		TemplateName: "default",
	})
	if err != nil {
		fmt.Printf("startClient BuildWithMarkdownFile %v", err)

		return err
	}

	if reply != nil {
		// fmt.Print(reply.HashName)
		fmt.Print(reply.Url)
	}

	return nil
}

func main() {
	cfg, err := adacore.LoadConfig("./config.yaml")
	if err != nil {
		fmt.Printf("startServ LoadConfig %v", err)

		return
	}

	adacore.InitLogger(cfg)

	startClient(cfg)
}
```

This is a ``markdown`` file.  
This a multiline string.

![This a image](sample001.jpg)

This is a ``markdown`` file.  
This a multiline string too.