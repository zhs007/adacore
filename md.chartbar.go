package adacore

import (
	"bytes"

	"gopkg.in/yaml.v2"
)

// ChartBar - chart bar infomation
type ChartBar struct {
	ID          string           `yaml:"id"`
	DatasetName string           `yaml:"datasetname"`
	Title       string           `yaml:"title"`
	SubText     string           `yaml:"subtext"`
	LegendData  []string         `yaml:"legenddata"`
	XType       string           `yaml:"xtype"`
	XData       string           `yaml:"xdata"`
	XShowAll    bool             `yaml:"xshowall"`
	YType       string           `yaml:"ytype"`
	YData       []ChartBasicData `yaml:"ydata"`
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
