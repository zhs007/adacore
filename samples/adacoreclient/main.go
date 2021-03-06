// Zerro
package main

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/zhs007/adacore"
	adacorepb "github.com/zhs007/adacore/adacorepb"
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

	md.AppendParagraph("")

	c := &adacore.Commodity{
		ID: "commodity001",
		Items: []*adacore.CommodityItem{
			&adacore.CommodityItem{
				Title:       "这是一件商品1",
				CurPrice:    999.99,
				ImgFileName: "./c.jpg",
				URL:         "https://ada.heyalgo.io/p1",
				Shop: adacore.CommodityShop{
					Name: "这是一个店铺1",
					URL:  "https://ada.heyalgo.io/shop1",
				},
			},
			&adacore.CommodityItem{
				Title:       "这是一件商品2",
				CurPrice:    1999.99,
				ImgFileName: "./c.jpg",
				URL:         "https://ada.heyalgo.io/p2",
				Shop: adacore.CommodityShop{
					Name: "这是一个店铺2",
					URL:  "https://ada.heyalgo.io/shop2",
				},
			},
		},
	}

	im, err := c.LoadImageMap(false)
	if err != nil {
		return nil, err
	}

	_, err = md.AppendCommodity(c, im, mddata)
	if err != nil {
		return nil, err
	}

	// mddata.BinaryData["sample001.jpg"] = buf

	mddata.StrData = md.GetMarkdownString(km)

	// fmt.Print(mddata.StrData)

	return mddata, nil
}

func startClient(cfg *adacore.Config) error {
	client := adacore.NewClient("47.91.209.141:7201", "x7sSGGHgmKwUMoa5S4VZlr9bUF2lCCzF")

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
		// fmt.Print(reply.HashName)
		fmt.Print(reply.Url)
	}

	return nil
}

func main() {
	adacore.InitTemplates()

	cfg, err := adacore.LoadConfig("./config.yaml")
	if err != nil {
		fmt.Printf("startServ LoadConfig %v", err)

		return
	}

	adacore.InitLogger(cfg)

	startClient(cfg)
}
