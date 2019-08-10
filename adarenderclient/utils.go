package adarenderclient

import (
	"io/ioutil"
	"path"

	"github.com/golang/protobuf/proto"

	adarender "github.com/zhs007/adacore/adarenderpb"
	adacorebase "github.com/zhs007/adacore/base"
)

// BuildMarkdownStream - MarkdownData => []MarkdownStream
func BuildMarkdownStream(mddata *adarender.MarkdownData, token string) ([]*adarender.MarkdownStream, error) {
	buf, err := proto.Marshal(mddata)
	if err != nil {
		return nil, err
	}

	bl := len(buf)
	if bl <= adacorebase.BigMsgLength {
		stream := &adarender.MarkdownStream{}

		stream.MarkdownData = mddata
		stream.Token = token

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

// BuildMarkdownData - MarkdownData => []MarkdownStream
func BuildMarkdownData(inpath string, filename string) (*adarender.MarkdownData, error) {
	mdbuf, err := ioutil.ReadFile(path.Join(inpath, filename))
	if err != nil {
		return nil, err
	}

	md := &adarender.MarkdownData{
		StrData:    string(mdbuf),
		BinaryData: make(map[string][]byte),
	}

	files, err := ioutil.ReadDir(inpath)
	if err != nil {
		return nil, err
	}

	for _, fn := range files {
		if fn.IsDir() {

		} else if fn.Name() != "__debug_bin" {
			buf, err := ioutil.ReadFile(fn.Name())
			if err != nil {
				return nil, err
			}

			md.BinaryData[fn.Name()] = buf
		}
	}

	return md, nil
}
