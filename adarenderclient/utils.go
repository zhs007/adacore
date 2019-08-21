package adarenderclient

import (
	"bytes"
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

	var lst []*adarender.MarkdownStream

	st := 0
	for st < bl {
		isend := false
		cl := adacorebase.BigMsgLength
		if cl > bl-st {
			cl = bl - st
			isend = true
		}

		cb := buf[st:(st + cl)]

		cs := &adarender.MarkdownStream{
			TotalLength: int32(bl),
			CurStart:    int32(st),
			CurLength:   int32(cl),
			HashData:    adacorebase.MD5Buffer(cb),
			Data:        cb,
			Token:       token,
		}

		st += cl
		if isend {
			cs.TotalHashData = adacorebase.MD5Buffer(buf)
		}

		lst = append(lst, cs)
	}

	return lst, nil
}

// BuildHTMLData - []HTMLStream => HTMLData
func BuildHTMLData(lst []*adarender.HTMLStream) (*adarender.HTMLData, error) {
	if len(lst) == 1 {
		if lst[0].HtmlData != nil {
			return lst[0].HtmlData, nil
		}

		return nil, adacorebase.ErrEmptyHTMLData
	}

	var lstbytes [][]byte
	totalmd5inmsg := ""
	st := 0
	ct := 0
	for i, v := range lst {
		if st != int(v.CurStart) {
			return nil, adacorebase.ErrInvalidCurStartAdaRender
		}

		if len(v.Data) != int(v.CurLength) {
			return nil, adacorebase.ErrInvalidCurLengthAdaRender
		}

		curmd5 := adacorebase.MD5Buffer(v.Data)
		if curmd5 != v.HashData {
			return nil, adacorebase.ErrInvalidHashDataAdaRender
		}

		lstbytes = append(lstbytes, v.Data)

		st += len(v.Data)
		ct += len(v.Data)

		if i == len(lst)-1 {
			totalmd5inmsg = v.TotalHashData

			if ct != int(v.TotalLength) {
				return nil, adacorebase.ErrInvalidTotalLengthAdaRender
			}
		}
	}

	buf := bytes.Join(lstbytes, []byte(""))

	totalmd5 := adacorebase.MD5Buffer(buf)
	if totalmd5 != totalmd5inmsg {
		return nil, adacorebase.ErrInvalidTotalHashDataAdaRender
	}

	htmld := &adarender.HTMLData{}
	err := proto.Unmarshal(buf, htmld)
	if err != nil {
		return nil, err
	}

	return htmld, nil
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
