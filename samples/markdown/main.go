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
	err := adacore.InitTemplates("../../templates")
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
	})
	if err != nil {
		fmt.Printf("AppendChartPie error %v", err)

		return
	}

	fmt.Printf("%v", md.GetMarkdownString(km))
}
