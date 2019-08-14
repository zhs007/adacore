package adacore

import (
	"bytes"
	"io/ioutil"
	"path"
	"strings"
	"text/template"

	"gopkg.in/yaml.v2"

	adacorebase "github.com/zhs007/adacore/base"
	adacorepb "github.com/zhs007/adacore/proto"
)

var tempDataset *template.Template
var tempLine *template.Template
var tempPie *template.Template
var tempBar *template.Template
var tempTreeMap *template.Template

// InitTemplates - init templates
func InitTemplates(dir string) error {
	tmp, err := template.ParseFiles(path.Join(dir, "adadataset.md"))
	if err != nil {
		return err
	}

	tempDataset = tmp

	tmp, err = template.ParseFiles(path.Join(dir, "adapie.md"))
	if err != nil {
		return err
	}

	tempPie = tmp

	tmp, err = template.ParseFiles(path.Join(dir, "adabar.md"))
	if err != nil {
		return err
	}

	tempBar = tmp

	tmp, err = template.ParseFiles(path.Join(dir, "adaline.md"))
	if err != nil {
		return err
	}

	tempLine = tmp

	tmp, err = template.ParseFiles(path.Join(dir, "adatreemap.md"))
	if err != nil {
		return err
	}

	tempTreeMap = tmp

	return nil
}

// Dataset - dataset
type Dataset struct {
	Name string      `yaml:"name"`
	Data interface{} `yaml:"data"`
}

// // Chart - chart basic infomation
// type Chart struct {
// 	ID          string `yaml:"id"`
// 	DatasetName string `yaml:"datasetname"`
// 	Title       string `yaml:"title"`
// 	SubText     string `yaml:"subtext"`
// 	Width       int    `yaml:"width"`
// 	Height      int    `yaml:"height"`
// }

// ChartBasicData - chart basic data
type ChartBasicData struct {
	Name string `yaml:"name"`
	Data string `yaml:"data"`
}

// ChartTreeMapData - chart treemap data
type ChartTreeMapData struct {
	Name     string             `yaml:"name"`
	Value    int                `yaml:"value"`
	Children []ChartTreeMapData `yaml:"children"`
}

// ChartPie - chart pie infomation
type ChartPie struct {
	ID          string `yaml:"id"`
	DatasetName string `yaml:"datasetname"`
	Title       string `yaml:"title"`
	SubText     string `yaml:"subtext"`
	Width       int    `yaml:"width"`
	Height      int    `yaml:"height"`
	A           string `yaml:"a"`
	BVal        string `yaml:"bval"`
	CVal        string `yaml:"cval"`
}

// ChartBar - chart bar infomation
type ChartBar struct {
	ID          string           `yaml:"id"`
	DatasetName string           `yaml:"datasetname"`
	Title       string           `yaml:"title"`
	SubText     string           `yaml:"subtext"`
	Width       int              `yaml:"width"`
	Height      int              `yaml:"height"`
	LegendData  []string         `yaml:"legenddata"`
	XType       string           `yaml:"xtype"`
	XData       string           `yaml:"xdata"`
	XShowAll    bool             `yaml:"xshowall"`
	YType       string           `yaml:"ytype"`
	YData       []ChartBasicData `yaml:"ydata"`
}

// ChartTreeMap - chart treemap infomation
type ChartTreeMap struct {
	ID      string             `yaml:"id"`
	Title   string             `yaml:"title"`
	SubText string             `yaml:"subtext"`
	Width   int                `yaml:"width"`
	Height  int                `yaml:"height"`
	Data    []ChartTreeMapData `yaml:"data"`
}

// baseObj -
type baseObj struct {
	Yaml string
}

// Markdown - markdown
type Markdown struct {
	// Title - title
	Title string
	// str - markdown string
	str string
}

// NewMakrdown - new Markdown
func NewMakrdown(title string) *Markdown {
	md := &Markdown{
		Title: title,
	}

	md.AppendParagraph("# " + title)

	return md
}

// isTitle - is this line a title?
func isTitle(strline string) bool {
	ns := strings.TrimLeft(strline, " ")
	if len(ns) > 0 {
		return ns[0] == '#'
	}

	return false
}

// isCodeBegin - is this line a code begin?
func isCodeBegin(strline string) bool {
	ns := strings.TrimLeft(strline, " ")
	if len(ns) >= 3 {
		return ns[0] == '`' && ns[1] == '`' && ns[2] == '`'
	}

	return false
}

// isCodeEnd - is this line a code end?
func isCodeEnd(strline string) bool {
	ns := strings.TrimLeft(strline, " ")
	if len(ns) >= 3 {
		return ns[0] == '`' && ns[1] == '`' && ns[2] == '`'
	}

	return false
}

// GetMarkdownString - get markdown string
func (md *Markdown) GetMarkdownString(lst *KeywordMappingList) string {
	if lst != nil && len(lst.Keywords) > 0 {
		lstline := strings.Split(md.str, "\n")
		incode := false

		for i, cl := range lstline {
			if len(cl) == 0 {
				continue
			}

			if !incode && !isTitle(cl) {
				if isCodeBegin(cl) {
					incode = true
				}

				if !incode {
					for _, v := range lst.Keywords {
						if v.URL == "" {
							lstline[i] = strings.Replace(lstline[i], v.Keyword,
								adacorebase.AppendString("``", v.Keyword, "``"), -1)
						} else {
							lstline[i] = strings.Replace(lstline[i], v.Keyword,
								adacorebase.AppendString("[", v.Keyword, "]("+v.URL+")"), -1)
						}
					}
				}
			} else if incode {
				if isCodeEnd(cl) {
					incode = false
				}
			}
		}

		md.str = strings.Join(lstline, "\n")
	}

	return md.str
}

// AppendParagraph - append paragraph
func (md *Markdown) AppendParagraph(str string) string {
	str = strings.Replace(str, "  \n", "\n", -1)
	str = strings.Replace(str, "\n", "  \n", -1)

	md.str = adacorebase.AppendString(md.str, str+"\n\n")

	return md.str
}

// AppendTable - append a table
func (md *Markdown) AppendTable(head []string, data [][]string) string {
	if len(head) > 0 {
		str := "|"

		for _, hv := range head {
			str += hv + "|"
		}

		str += "\n|"

		for range head {
			str += "---|"
		}

		str += "\n"

		for _, li := range data {
			str += "|"
			for _, ld := range li {
				str += ld + "|"
			}
			str += "\n"
		}

		md.str = adacorebase.AppendString(md.str, str+"\n\n")
	}

	return md.str
}

// AppendCode - append code
func (md *Markdown) AppendCode(code string, codetype string) string {
	md.str = adacorebase.AppendString(md.str, "``` ", codetype, "\n", code, "\n```\n\n")

	return md.str
}

// AppendImage - append image
func (md *Markdown) AppendImage(text string, fn string, mddata *adacorepb.MarkdownData) (
	[]byte, string, error) {

	buf, err := ioutil.ReadFile(fn)
	if err != nil {
		return nil, "", err
	}

	md.str = adacorebase.AppendString(md.str, "![", text, "](", fn, ")")

	if mddata.BinaryData == nil {
		mddata.BinaryData = make(map[string][]byte)
	}

	mddata.BinaryData[fn] = buf

	return buf, md.str, nil
}

// AppendImageBuf - append image buf
func (md *Markdown) AppendImageBuf(text string, name string, buf []byte, mddata *adacorepb.MarkdownData) (
	[]byte, string, error) {

	md.str = adacorebase.AppendString(md.str, "![", text, "](", name, ")")

	if mddata.BinaryData == nil {
		mddata.BinaryData = make(map[string][]byte)
	}

	mddata.BinaryData[name] = buf

	return buf, md.str, nil
}

// AppendDataset - append dataset, the obj should be an object that can be encoded by yaml
func (md *Markdown) AppendDataset(name string, data interface{}) (
	string, error) {

	obj := &Dataset{
		Name: name,
		Data: data,
	}

	d, err := yaml.Marshal(obj)
	if err != nil {
		return "", err
	}

	var b bytes.Buffer
	err = tempDataset.Execute(&b, &baseObj{
		Yaml: string(d),
	})
	if err != nil {
		return "", err
	}

	md.str += b.String()

	return md.str, nil
}

// AppendChartPie - append chart pie, the obj should be an object that can be encoded by yaml
func (md *Markdown) AppendChartPie(pie *ChartPie) (
	string, error) {

	d, err := yaml.Marshal(pie)
	if err != nil {
		return "", err
	}

	var b bytes.Buffer
	err = tempPie.Execute(&b, baseObj{
		Yaml: string(d),
	})
	if err != nil {
		return "", err
	}

	md.str += b.String()

	return md.str, nil
}

// AppendChartLine - append chart line, the obj should be an object that can be encoded by yaml
func (md *Markdown) AppendChartLine(obj interface{}) (
	string, error) {

	d, err := yaml.Marshal(obj)
	if err != nil {
		return "", err
	}

	var b bytes.Buffer

	err = tempLine.Execute(&b, baseObj{
		Yaml: string(d),
	})
	if err != nil {
		return "", err
	}

	md.str += b.String()

	return md.str, nil
}

// AppendChartBar - append chart bar, the obj should be an object that can be encoded by yaml
func (md *Markdown) AppendChartBar(bar *ChartBar) (
	string, error) {

	d, err := yaml.Marshal(bar)
	if err != nil {
		return "", err
	}

	var b bytes.Buffer
	err = tempBar.Execute(&b, baseObj{
		Yaml: string(d),
	})
	if err != nil {
		return "", err
	}

	md.str += b.String()

	return md.str, nil
}

// AppendChartTreeMap - append chart treemap, the obj should be an object that can be encoded by yaml
func (md *Markdown) AppendChartTreeMap(treemap *ChartTreeMap) (
	string, error) {

	d, err := yaml.Marshal(treemap)
	if err != nil {
		return "", err
	}

	var b bytes.Buffer
	err = tempTreeMap.Execute(&b, baseObj{
		Yaml: string(d),
	})
	if err != nil {
		return "", err
	}

	md.str += b.String()

	return md.str, nil
}
