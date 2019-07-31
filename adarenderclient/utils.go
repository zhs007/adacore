package adarenderclient

import (
	"github.com/golang/protobuf/proto"

	adarender "github.com/zhs007/adacore/adarenderpb"
	"github.com/zhs007/jarviscore/basedef"
)

// BuildMarkdownStream - MarkdownData => []MarkdownStream
func BuildMarkdownStream(mddata *adarender.MarkdownData) ([]*adarender.MarkdownStream, error) {
	buf, err := proto.Marshal(mddata)
	if err != nil {
		return nil, err
	}

	bl := len(buf)
	if bl <= basedef.BigMsgLength {
		stream := &adarender.MarkdownStream{}

		stream.MarkdownData = mddata

		return []*adarender.MarkdownStream{stream}, nil
	}

	return nil, nil
}

// BuildHTMLData - []HTMLStream => HTMLData
func BuildHTMLData(lst []*adarender.HTMLStream) (*adarender.HTMLData, error) {
	if len(lst) == 1 {
		if lst[0].HtmlData != nil {
			return lst[0].HtmlData, nil
		}
	}

	return nil, nil
}
