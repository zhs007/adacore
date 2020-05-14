package adacore

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/golang/protobuf/proto"
	adarender "github.com/zhs007/adacore/adarenderpb"
	adacorebase "github.com/zhs007/adacore/base"
	adacorepb "github.com/zhs007/adacore/adacorepb"
)

// BuildMarkdownData - []HTMLStream => HTMLData
func BuildMarkdownData(lst []*adacorepb.MarkdownStream) (*adacorepb.MarkdownData, error) {
	if len(lst) == 1 {
		if lst[0].MarkdownData != nil {
			return lst[0].MarkdownData, nil
		}

		return nil, adacorebase.ErrNoMarkdownData
	}

	var lstbytes [][]byte
	for _, v := range lst {
		lstbytes = append(lstbytes, v.Data)
	}

	buf := bytes.Join(lstbytes, nil)

	md := &adacorepb.MarkdownData{}
	err := proto.Unmarshal(buf, md)
	if err != nil {
		return nil, err
	}

	return md, nil
}

// SaveHTMLData - save html
func SaveHTMLData(htmldata *adarender.HTMLData, hashname string, cfg *Config) (string, error) {
	if htmldata != nil && len(htmldata.StrData) > 0 {
		if hashname == "" {
			hashname = adacorebase.HashBuffer([]byte(htmldata.StrData)) + ".html"
		}

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

	var lst []*adacorepb.MarkdownStream

	st := 0
	for st < bl {
		isend := false
		cl := adacorebase.BigMsgLength
		if cl > bl-st {
			cl = bl - st
			isend = true
		}

		cb := buf[st:(st + cl)]

		cs := &adacorepb.MarkdownStream{
			TotalLength: int32(bl),
			CurStart:    int32(st),
			CurLength:   int32(cl),
			HashData:    adacorebase.HashBuffer(cb),
			Data:        cb,
			Token:       token,
		}

		st += cl
		if isend {
			cs.TotalHashData = adacorebase.HashBuffer(buf)
		}

		lst = append(lst, cs)
	}

	return lst, nil
}

var oldTableString = []string{
	"[",
	"]",
	"|",
}

var newTableString = []string{
	"{",
	"}",
	"/",
}

// FixTableString - fix table string
func FixTableString(str string) string {
	for i, v := range oldTableString {
		str = strings.ReplaceAll(str, v, newTableString[i])
	}

	return str
}

// LoadMarkdownAndFiles - load a markdown file & somes files
func LoadMarkdownAndFiles(fn string, globpattern string) (*adacorepb.MarkdownData, error) {
	buf, err := ioutil.ReadFile(fn)
	if err != nil {
		return nil, err
	}

	mdd := &adacorepb.MarkdownData{
		StrData:      string(buf),
		BinaryData:   make(map[string][]byte),
		TemplateName: "default",
	}

	matches, err := filepath.Glob(globpattern)
	if err != nil {
		return nil, err
	}

	for _, v := range matches {
		_, cfn := filepath.Split(v)
		buf, err := ioutil.ReadFile(v)
		if err != nil {
			return nil, err
		}

		mdd.BinaryData[cfn] = buf
	}

	return mdd, nil
}
