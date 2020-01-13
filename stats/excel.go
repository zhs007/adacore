package stats

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	excel "github.com/zhs007/adacore/excel"
)

// memberDataSetStats - DataSetStats member
var memberDataSetStats = []string{
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

// ExportExcel - export a excel file
func ExportExcel(f *excelize.File, sheet string, lst []*DataSetStats, floatFormat string) {

	// write head
	excel.SetSheet(f, sheet, 1, 1, memberDataSetStats, len(lst),
		func(i int, member string) string {
			return ""
		},
		func(i int, member string) (interface{}, error) {
			v := lst[i]

			if member == "Name" {
				return v.Name, nil
			} else if member == "Nums" {
				return v.Nums, nil
			} else if member == "MeanSDev1" {
				return fmt.Sprintf("%.2f", float32(v.MeanSDev1)/float32(v.Nums)), nil
			} else if member == "MeanSDev2" {
				return fmt.Sprintf("%.2f", float32(v.MeanSDev2)/float32(v.Nums)), nil
			} else if member == "MeanSDev3" {
				return fmt.Sprintf("%.2f", float32(v.MeanSDev3)/float32(v.Nums)), nil
			} else if member == "Min" {
				return fmt.Sprintf(floatFormat, v.Min), nil
			} else if member == "Max" {
				return fmt.Sprintf(floatFormat, v.Max), nil
			} else if member == "Median" {
				return fmt.Sprintf(floatFormat, v.Median), nil
			} else if member == "MedianAbsoluteDeviation" {
				return fmt.Sprintf(floatFormat, v.MedianAbsoluteDeviation), nil
			} else if member == "MedianAbsoluteDeviationPopulation" {
				return fmt.Sprintf(floatFormat, v.MedianAbsoluteDeviationPopulation), nil
			} else if member == "Midhinge" {
				return fmt.Sprintf(floatFormat, v.Midhinge), nil
			} else if member == "Mean" {
				return fmt.Sprintf(floatFormat, v.Mean), nil
			} else if member == "GeometricMean" {
				return fmt.Sprintf(floatFormat, v.GeometricMean), nil
			} else if member == "HarmonicMean" {
				return fmt.Sprintf(floatFormat, v.HarmonicMean), nil
			} else if member == "InterQuartileRange" {
				return fmt.Sprintf(floatFormat, v.InterQuartileRange), nil
			} else if member == "StandardDeviation" {
				return fmt.Sprintf(floatFormat, v.StandardDeviation), nil
			} else if member == "StandardDeviationPopulation" {
				return fmt.Sprintf(floatFormat, v.StandardDeviationPopulation), nil
			} else if member == "StandardDeviationSample" {
				return fmt.Sprintf(floatFormat, v.StandardDeviationSample), nil
			} else if member == "Trimean" {
				return fmt.Sprintf(floatFormat, v.Trimean), nil
			} else if member == "Variance" {
				return fmt.Sprintf(floatFormat, v.Variance), nil
			} else if member == "PopulationVariance" {
				return fmt.Sprintf(floatFormat, v.PopulationVariance), nil
			} else if member == "SampleVariance" {
				return fmt.Sprintf(floatFormat, v.SampleVariance), nil
			} else if member == "COV" {
				return fmt.Sprintf(floatFormat, v.COV), nil
			}

			return nil, nil
		})
}
