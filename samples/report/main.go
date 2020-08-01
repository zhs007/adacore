// Zerro
package main

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/zhs007/adacore"
	adacorepb "github.com/zhs007/adacore/adacorepb"
)

type commondataset struct {
	NameData []string `yaml:"namedata"`
	ValData  []int    `yaml:"valdata"`
}

func genReport() (*adacorepb.MarkdownData, error) {
	// new a markdown data
	mddata := &adacorepb.MarkdownData{
		TemplateName: "default",               // the template is default
		BinaryData:   make(map[string][]byte), // If you need to upload some images, you need to initialize the binary data.
	}

	// You can automate all keywords with keyword mapping
	km, err := adacore.LoadKeywordMappingList("./keywordmapping.yaml")
	if err != nil {
		fmt.Printf("load keywordmapping error %v", err)
	}

	// new a markdown
	md := adacore.NewMakrdown("Report")

	md.AppendParagraph("> This libraray is write by Zerro.\nIf you have any questions, you can contact Zerro.")

	md.AppendParagraph("### Table")

	md.AppendTable([]string{"head0", "head1", "head2"}, [][]string{
		[]string{"text0_0", "text1_0", "text2_0"},
		[]string{"text0_1", "text1_1", "text2_1"},
		[]string{"text0_2", "text1_2", "text2_2"}})

	md.AppendParagraph("")

	md.AppendParagraph("### Code")

	md.AppendParagraph("> This is core code.")

	fd, err := ioutil.ReadFile("./main.go")
	if err != nil {
		return nil, err
	}

	md.AppendCode(string(fd), "golang")

	md.AppendParagraph("### Image")

	_, _, err = md.AppendImage("This is a image", "sample001.jpg", mddata)
	if err != nil {
		return nil, err
	}

	md.AppendParagraph("### Tree Map")

	_, err = md.AppendChartTreeMapFloat(&adacore.ChartTreeMapFloat{
		ID:         "treemap001",
		Title:      "TreeMap",
		SubText:    "test treemap chart",
		Width:      1280,
		Height:     800,
		LegendData: []string{"test1"},
		TreeMap: []adacore.ChartTreeMapSeriesNodeFloat{
			adacore.ChartTreeMapSeriesNodeFloat{
				Name: "test1",
				Data: []adacore.ChartTreeMapDataFloat{
					adacore.ChartTreeMapDataFloat{
						Name: "nodeA",
						Children: []adacore.ChartTreeMapDataFloat{
							adacore.ChartTreeMapDataFloat{
								Name:  "nodeAa",
								Value: 6.5,
							},
							adacore.ChartTreeMapDataFloat{
								Name:  "nodeAa",
								Value: 6,
							},
						},
					},
					adacore.ChartTreeMapDataFloat{
						Name: "nodeB",
						Children: []adacore.ChartTreeMapDataFloat{
							adacore.ChartTreeMapDataFloat{
								Name: "nodeBa",
								Children: []adacore.ChartTreeMapDataFloat{
									adacore.ChartTreeMapDataFloat{
										Name:  "nodeBa1",
										Value: 18.8,
									},
								},
							},
						},
					},
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}

	md.AppendParagraph("### Other Charts")

	_, err = md.AppendDataset("pieds", &commondataset{
		NameData: []string{"dat0", "dat1", "dat2", "dat3", "dat4", "dat5", "dat6", "dat7"},
		ValData:  []int{100, 123, 89, 120, 50, 37, 21, 87},
	})
	if err != nil {
		return nil, err
	}

	_, err = md.AppendChartPie(&adacore.ChartPie{
		ID:          "pie001",
		DatasetName: "pieds",
		Title:       "Pie",
		SubText:     "test pie chart",
		Width:       1280,
		Height:      800,
		A:           "pie name",
		BVal:        "namedata",
		CVal:        "valdata",
		Sort:        adacore.ChartSortReverse,
	})
	if err != nil {
		return nil, err
	}

	_, err = md.AppendChartBar(&adacore.ChartBar{
		ID:          "bar001",
		DatasetName: "pieds",
		Title:       "Bar",
		SubText:     "test bar chart",
		LegendData:  []string{"val1"},
		XType:       "category",
		XData:       "namedata",
		XShowAll:    true,
		YType:       "value",
		YData: []adacore.ChartBasicData{
			adacore.ChartBasicData{
				Name: "val1",
				Data: "valdata",
			},
		},
	})
	if err != nil {
		return nil, err
	}

	mddata.StrData = md.GetMarkdownString(km)

	return mddata, nil
}

func startClient() error {
	client := adacore.NewClient("47.91.209.141:7201", "x7sSGGHgmKwUMoa5S4VZlr9bUF2lCCzF")

	md, err := genReport()
	if err != nil {
		fmt.Printf("startClient genReport %v", err)

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
	startClient()
}
