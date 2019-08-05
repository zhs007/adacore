package main

import (
	"fmt"

	"github.com/zhs007/adacore"
)

func main() {
	km, err := adacore.LoadKeywordMappingList("./keywordmapping.yaml")
	if err != nil {
		fmt.Printf("load keywordmapping error %v", err)
	}

	md := adacore.NewMakrdown("Ada Core")

	md.AppendParagraph("This is a Markdown API for Ada.")
	md.AppendParagraph("This libraray is write by Zerro.")

	fmt.Printf("%v", md.GetMarkdownString(km))
}
