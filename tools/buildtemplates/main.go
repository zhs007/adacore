package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

// TemplateInfo - template infomation
type TemplateInfo struct {
	Name string
	Str  string
}

// ParamObj - object
type ParamObj struct {
	Templates []TemplateInfo
}

// buildTemplate - template => base64 string
func buildTemplate(fn string) (string, error) {
	buf, err := ioutil.ReadFile(fn)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(buf), nil
}

// buildTemplate - template => base64 string
func buildInPath(dir string) (*ParamObj, error) {
	param := &ParamObj{}

	lst, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, fi := range lst {
		if fi.IsDir() {

		} else {
			arr := strings.Split(fi.Name(), ".")
			if len(arr) == 2 && arr[1] == "md" {
				str, err := buildTemplate(dir + string(os.PathSeparator) + fi.Name())
				if err != nil {
					return nil, err
				}

				param.Templates = append(param.Templates, TemplateInfo{
					Name: arr[0],
					Str:  str,
				})
			}
		}
	}

	return param, nil
}

func main() {
	param, err := buildInPath("./templates")
	if err != nil {
		fmt.Printf("buildInPath err %v", err)

		return
	}

	tmp, err := template.ParseFiles("./tools/buildtemplates/gofile.md")
	if err != nil {
		fmt.Printf("ParseFiles err %v", err)

		return
	}

	var b bytes.Buffer
	err = tmp.Execute(&b, *param)
	if err != nil {
		fmt.Printf("Execute err %v", err)

		return
	}

	err = ioutil.WriteFile("./templates.go", b.Bytes(), 0644)
	if err != nil {
		fmt.Printf("WriteFile err %v", err)

		return
	}
}
