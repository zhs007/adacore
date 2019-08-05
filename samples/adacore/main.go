package main

import (
	"context"
	"fmt"
	"time"

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
	md.AppendParagraph("This libraray is write by Zerro.")

	md.AppendTable([]string{"head0", "head1", "head2"}, [][]string{
		[]string{"text0_0", "text1_0", "text2_0"},
		[]string{"text0_1", "text1_1", "text2_1"}})

	return md.GetMarkdownString(km)
}

func startServ(cfg *adacore.Config, endchan chan int) (*adacore.Serv, error) {

	serv, err := adacore.NewAdaCoreServ(cfg)
	if err != nil {
		fmt.Printf("startServ NewAdaCoreServ %v", err)

		return nil, err
	}

	go func() {
		err := serv.Start(context.Background())
		if err != nil {
			fmt.Printf("startServ Start %v", err)
		}

		endchan <- 0
	}()

	return serv, nil
}

func startClient(cfg *adacore.Config) error {
	client := adacore.NewClient("127.0.0.1:7201", "x7sSGGHgmKwUMoa5S4VZlr9bUF2lCCzF")

	reply, err := client.BuildWithMarkdown(context.Background(), &adacorepb.MarkdownData{
		StrData:      genMarkdown(),
		TemplateName: "default",
	})
	if err != nil {
		fmt.Printf("startClient BuildWithMarkdownFile %v", err)

		return err
	}

	if reply != nil {
		fmt.Print(reply.HashName)
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

	servchan := make(chan int)
	serv, err := startServ(cfg, servchan)
	if err != nil {
		return
	}

	time.Sleep(1000 * 3)

	startClient(cfg)

	serv.Stop()

	<-servchan

	// serv.Stop()
}
