package stats

import (
	"fmt"

	adacore "github.com/zhs007/adacore"
)

var tableHead = []string{
	"Name",
	"Nums",
	"MeanSDev1",
	"MeanSDev2",
	"MeanSDev3",
	"Min",
	"Max",
	"COV",
	"Median",
	"MedianAbsoluteDeviation",
	"MedianAbsoluteDeviationPopulation",
	"Midhinge",
	"Mean",
	"GeometricMean",
	"HarmonicMean",
	"InterQuartileRange",
	"StandardDeviation",
	"StandardDeviationPopulation",
	"StandardDeviationSample",
	"Trimean",
	"Variance",
	"PopulationVariance",
	"SampleVariance",
}

// AppendTable - append table
func AppendTable(md *adacore.Markdown, lst []*DataSetStats) error {

	var arr [][]string
	for _, v := range lst {
		cl := []string{
			v.Name,
			fmt.Sprintf("%v", v.Nums),
			fmt.Sprintf("%.2f", float32(v.MeanSDev1)/float32(v.Nums)),
			fmt.Sprintf("%.2f", float32(v.MeanSDev2)/float32(v.Nums)),
			fmt.Sprintf("%.2f", float32(v.MeanSDev3)/float32(v.Nums)),
			fmt.Sprintf("%v", v.Min),
			fmt.Sprintf("%v", v.Max),
			fmt.Sprintf("%v", v.COV),
			fmt.Sprintf("%v", v.Median),
			fmt.Sprintf("%v", v.MedianAbsoluteDeviation),
			fmt.Sprintf("%v", v.MedianAbsoluteDeviationPopulation),
			fmt.Sprintf("%v", v.Midhinge),
			fmt.Sprintf("%v", v.Mean),
			fmt.Sprintf("%v", v.GeometricMean),
			fmt.Sprintf("%v", v.HarmonicMean),
			fmt.Sprintf("%v", v.InterQuartileRange),
			fmt.Sprintf("%v", v.StandardDeviation),
			fmt.Sprintf("%v", v.StandardDeviationPopulation),
			fmt.Sprintf("%v", v.StandardDeviationSample),
			fmt.Sprintf("%v", v.Trimean),
			fmt.Sprintf("%v", v.Variance),
			fmt.Sprintf("%v", v.PopulationVariance),
			fmt.Sprintf("%v", v.SampleVariance),
		}

		arr = append(arr, cl)

	}

	if len(arr) > 0 {
		md.AppendTable(tableHead, arr)
	}

	return nil
}
