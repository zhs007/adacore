package adacore

import (
	adarender "github.com/zhs007/adacore/adarenderpb"
	adacorepb "github.com/zhs007/adacore/proto"
)

// BuildMarkdownData - []HTMLStream => HTMLData
func BuildMarkdownData(lst []*adacorepb.MarkdownStream) (*adacorepb.MarkdownData, error) {
	if len(lst) == 1 {
		if lst[0].MarkdownData != nil {
			return lst[0].MarkdownData, nil
		}
	}

	return nil, nil
}

// SaveHTMLData - save html
func SaveHTMLData(htmldata *adarender.HTMLData) error {
	return nil
}
