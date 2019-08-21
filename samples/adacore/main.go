package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/zhs007/adacore"
	adacorepb "github.com/zhs007/adacore/proto"
)

func genMarkdown() (*adacorepb.MarkdownData, error) {
	mddata := &adacorepb.MarkdownData{
		TemplateName: "default",
		BinaryData:   make(map[string][]byte),
	}

	km, err := adacore.LoadKeywordMappingList("./keywordmapping.yaml")
	if err != nil {
		fmt.Printf("load keywordmapping error %v", err)
	}

	md := adacore.NewMakrdown("Ada Core")

	// for i := 0; i < 500; i++ {
	md.AppendParagraph("This is a Markdown API for Ada.")
	md.AppendParagraph("This libraray is write by Zerro.\nThis is a multi-line text.")

	md.AppendTable([]string{"head0", "head1", "head2"}, [][]string{
		[]string{"text0_0", "text1_0", "text2_0"},
		[]string{"text0_1", "text1_1", "text2_1"}})

	fd, err := ioutil.ReadFile("./main.go")
	if err != nil {
		return nil, err
	}

	md.AppendCode(string(fd), "golang")

	md.AppendParagraph("This libraray is write by Zerro.\nThis is a multi-line text.")

	_, _, err = md.AppendImage("This is a image", "sample001.jpg", mddata)
	if err != nil {
		return nil, err
	}

	// _, _, err = md.AppendImage("This is a image", "jpg.png", mddata)
	// if err != nil {
	// 	return nil, err
	// }

	// _, _, err = md.AppendImage("This is a image", "output.png", mddata)
	// if err != nil {
	// 	return nil, err
	// }

	// _, _, err = md.AppendImage("This is a image", "q100webp.png", mddata)
	// if err != nil {
	// 	return nil, err
	// }
	// }

	// mddata.BinaryData["sample001.jpg"] = buf

	mddata.StrData = md.GetMarkdownString(km)

	// fmt.Print(str)

	return mddata, nil
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

	md, err := genMarkdown()
	if err != nil {
		fmt.Printf("startClient genMarkdown %v", err)

		return err
	}

	reply, err := client.BuildWithMarkdown(context.Background(), md)
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
