package main

import (
	"fmt"

	"github.com/zhs007/adacore"
)

type piedataset struct {
	NameData []string `yaml:"namedata"`
	ValData  []int    `yaml:"valdata"`
}

func main() {
	err := adacore.InitTemplates()
	if err != nil {
		fmt.Printf("InitTemplates error %v", err)

		return
	}

	km, err := adacore.LoadKeywordMappingList("./keywordmapping.yaml")
	if err != nil {
		fmt.Printf("load keywordmapping error %v", err)

		return
	}

	md := adacore.NewMakrdown("Ada Core")

	md.AppendParagraph("This is a Markdown API for Ada.")
	md.AppendParagraph("This libraray is write by Zerro.")

	md.AppendTable([]string{"head0", "head1", "head2"}, [][]string{
		[]string{"text0_0", "text1_0", "text2_0"},
		[]string{"text0_1", "text1_1", "text2_1"}})

	_, err = md.AppendDataset("pieds", &piedataset{
		NameData: []string{"dat0", "dat1", "dat2"},
		ValData:  []int{100, 123, 89},
	})
	if err != nil {
		fmt.Printf("AppendDataset error %v", err)

		return
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
		fmt.Printf("AppendChartPie error %v", err)

		return
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
		fmt.Printf("AppendChartPie error %v", err)

		return
	}

	_, err = md.AppendChartTreeMap(&adacore.ChartTreeMap{
		ID:         "treemap001",
		Title:      "TreeMap",
		SubText:    "test treemap chart",
		Width:      1280,
		Height:     800,
		LegendData: []string{"test1"},
		TreeMap: []adacore.ChartTreeMapSeriesNode{
			adacore.ChartTreeMapSeriesNode{
				Name: "test1",
				Data: []adacore.ChartTreeMapData{
					adacore.ChartTreeMapData{
						Name: "nodeA",
						Children: []adacore.ChartTreeMapData{
							adacore.ChartTreeMapData{
								Name:  "nodeAa",
								Value: 6,
							},
							adacore.ChartTreeMapData{
								Name:  "nodeAa",
								Value: 6,
							},
						},
					},
					adacore.ChartTreeMapData{
						Name: "nodeB",
						Children: []adacore.ChartTreeMapData{
							adacore.ChartTreeMapData{
								Name: "nodeBa",
								Children: []adacore.ChartTreeMapData{
									adacore.ChartTreeMapData{
										Name:  "nodeBa1",
										Value: 20,
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
		fmt.Printf("AppendChartTreeMap error %v", err)

		return
	}

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
		fmt.Printf("AppendChartTreeMapFloat error %v", err)

		return
	}

	fmt.Printf("%v", md.GetMarkdownString(km))
}
