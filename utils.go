package adacore

import (
	"io/ioutil"
	"path/filepath"

	"github.com/golang/protobuf/proto"
	adarender "github.com/zhs007/adacore/adarenderpb"
	adacorebase "github.com/zhs007/adacore/base"
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
func SaveHTMLData(htmldata *adarender.HTMLData, cfg *Config) (string, error) {
	if htmldata != nil && len(htmldata.StrData) > 0 {
		hashname := adacorebase.HashBuffer([]byte(htmldata.StrData)) + ".html"

		fn := filepath.Join(cfg.FilePath, hashname)

		err := ioutil.WriteFile(fn, []byte(htmldata.StrData), 0644)
		if err != nil {
			return "", err
		}

		if len(htmldata.BinaryData) > 0 {
			for hn, buf := range htmldata.BinaryData {
				rfn := filepath.Join(cfg.FilePath, hn)
				err := ioutil.WriteFile(rfn, buf, 0644)
				if err != nil {
					// 	return "", err
				}
			}
		}

		return hashname, nil
	}

	return "", adacorebase.ErrEmptyHTMLData
}

// BuildMarkdownStream - MarkdownData => []MarkdownStream
func BuildMarkdownStream(mddata *adacorepb.MarkdownData, token string) ([]*adacorepb.MarkdownStream, error) {
	buf, err := proto.Marshal(mddata)
	if err != nil {
		return nil, err
	}

	bl := len(buf)
	if bl <= adacorebase.BigMsgLength {
		stream := &adacorepb.MarkdownStream{}

		stream.MarkdownData = mddata
		stream.Token = token

		return []*adacorepb.MarkdownStream{stream}, nil
	}

	return nil, nil
}
