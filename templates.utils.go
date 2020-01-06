package adacore

import (
	"encoding/base64"
	"text/template"
)

var tempDataset *template.Template
var tempLine *template.Template
var tempPie *template.Template
var tempBar *template.Template
var tempTreeMap *template.Template
var tempCommodity *template.Template

func parseBase64(name string, str string) (*template.Template, error) {
	buf, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return nil, err
	}

	temp := template.New(name)
	temp.Parse(string(buf))

	return temp, nil
}

// InitTemplates - init templates
func InitTemplates() error {
	tmp, err := parseBase64("adadataset", templateadadataset)
	if err != nil {
		return err
	}

	tempDataset = tmp

	tmp, err = parseBase64("adapie", templateadapie)
	if err != nil {
		return err
	}

	tempPie = tmp

	tmp, err = parseBase64("adabar", templateadabar)
	if err != nil {
		return err
	}

	tempBar = tmp

	tmp, err = parseBase64("adaline", templateadaline)
	if err != nil {
		return err
	}

	tempLine = tmp

	tmp, err = parseBase64("adatreemap", templateadatreemap)
	if err != nil {
		return err
	}

	tempTreeMap = tmp

	tmp, err = parseBase64("adacommodity", templatecommodity)
	if err != nil {
		return err
	}

	tempCommodity = tmp

	return nil
}
