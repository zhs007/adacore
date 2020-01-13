package stats

import (
	statsbase "github.com/montanaflynn/stats"
	adacorebase "github.com/zhs007/adacore/base"
	"go.uber.org/zap"
)

// DataSetStats - data set stats
type DataSetStats struct {
	Name                              string  `json:"Name"`
	Nums                              int     `json:"Nums"`
	MeanSDev1                         int     `json:"MeanSDev1"`                         // 1个标准差以内的数量
	MeanSDev2                         int     `json:"MeanSDev2"`                         // 2个标准差以内的数量
	MeanSDev3                         int     `json:"MeanSDev3"`                         // 3个标准差以内的数量
	COV                               float64 `json:"COV"`                               // coefficient of variation, StandardDeviation / Mean
	Min                               float64 `json:"Min"`                               // 最小值
	Max                               float64 `json:"Max"`                               // 最大值
	Median                            float64 `json:"Median"`                            // 中位数
	MedianAbsoluteDeviation           float64 `json:"MedianAbsoluteDeviation"`           // 中值绝对偏差
	MedianAbsoluteDeviationPopulation float64 `json:"MedianAbsoluteDeviationPopulation"` // 离散中值绝对偏差
	Midhinge                          float64 `json:"Midhinge"`                          // 中枢纽，第1四分位数和第3四分位数的算术平均值
	Mean                              float64 `json:"Mean"`                              // 平均数
	GeometricMean                     float64 `json:"GeometricMean"`                     // 几何平均数
	HarmonicMean                      float64 `json:"HarmonicMean"`                      // 谐波均值
	InterQuartileRange                float64 `json:"InterQuartileRange"`                // 4分位间距
	StandardDeviation                 float64 `json:"StandardDeviation"`                 // 标准差
	StandardDeviationPopulation       float64 `json:"StandardDeviationPopulation"`
	StandardDeviationSample           float64 `json:"StandardDeviationSample"`
	Trimean                           float64 `json:"Trimean"`            // 3均值
	Variance                          float64 `json:"Variance"`           // 方差
	PopulationVariance                float64 `json:"PopulationVariance"` // 离散方差
	SampleVariance                    float64 `json:"SampleVariance"`     // 采样方差
	// Entropy                           float64 // 熵
}

// FuncGetValue - get value
type FuncGetValue func(index int) (float64, bool)

// AnalayzeDataSet - analayze data set
func AnalayzeDataSet(length int, name string, funcGetValue FuncGetValue) (*DataSetStats, error) {
	lstd := []float64{}

	for i := 0; i < length; i++ {
		v, isok := funcGetValue(i)
		if isok {
			lstd = append(lstd, v)
		}
	}

	dss := &DataSetStats{
		Name: name,
		Nums: len(lstd),
	}

	min, err := statsbase.Min(lstd)
	if err != nil {
		adacorebase.Warn("AnalayzeDataSet.Min",
			zap.Error(err))
		// return nil, err
	}

	dss.Min = min

	max, err := statsbase.Max(lstd)
	if err != nil {
		adacorebase.Warn("AnalayzeDataSet.Max",
			zap.Error(err))
		// return nil, err
	}

	dss.Max = max

	median, err := statsbase.Median(lstd)
	if err != nil {
		adacorebase.Warn("AnalayzeDataSet.Median",
			zap.Error(err))
		// return nil, err
	}

	dss.Median = median

	mad, err := statsbase.MedianAbsoluteDeviation(lstd)
	if err != nil {
		adacorebase.Warn("AnalayzeDataSet.MedianAbsoluteDeviation",
			zap.Error(err))
		// return nil, err
	}

	dss.MedianAbsoluteDeviation = mad

	madp, err := statsbase.MedianAbsoluteDeviationPopulation(lstd)
	if err != nil {
		adacorebase.Warn("AnalayzeDataSet.MedianAbsoluteDeviationPopulation",
			zap.Error(err))
		// return nil, err
	}

	dss.MedianAbsoluteDeviationPopulation = madp

	midhinge, err := statsbase.Midhinge(lstd)
	if err != nil {
		adacorebase.Warn("AnalayzeDataSet.Midhinge",
			zap.Error(err))
		// return nil, err
	}

	dss.Midhinge = midhinge

	mean, err := statsbase.Mean(lstd)
	if err != nil {
		adacorebase.Warn("AnalayzeDataSet.Mean",
			zap.Error(err))
		// return nil, err
	}

	dss.Mean = mean

	gm, err := statsbase.GeometricMean(lstd)
	if err != nil {
		adacorebase.Warn("AnalayzeDataSet.GeometricMean",
			zap.Error(err))
		// return nil, err
	}

	dss.GeometricMean = gm

	if CanHarmonicMean(lstd) {
		hm, err := statsbase.HarmonicMean(lstd)
		if err != nil {
			adacorebase.Warn("AnalayzeDataSet.HarmonicMean",
				zap.Error(err))
			// return nil, err
		}

		dss.HarmonicMean = hm
	}

	iqr, err := statsbase.InterQuartileRange(lstd)
	if err != nil {
		adacorebase.Warn("AnalayzeDataSet.InterQuartileRange",
			zap.Error(err))
		// return nil, err
	}

	dss.InterQuartileRange = iqr

	sd, err := statsbase.StandardDeviation(lstd)
	if err != nil {
		adacorebase.Warn("AnalayzeDataSet.StandardDeviation",
			zap.Error(err))
		// return nil, err
	}

	dss.StandardDeviation = sd

	sdp, err := statsbase.StandardDeviationPopulation(lstd)
	if err != nil {
		adacorebase.Warn("AnalayzeDataSet.StandardDeviationPopulation",
			zap.Error(err))
		// return nil, err
	}

	dss.StandardDeviationPopulation = sdp

	sds, err := statsbase.StandardDeviationSample(lstd)
	if err != nil {
		adacorebase.Warn("AnalayzeDataSet.StandardDeviationSample",
			zap.Error(err))
		// return nil, err
	}

	dss.StandardDeviationSample = sds

	trimean, err := statsbase.Trimean(lstd)
	if err != nil {
		adacorebase.Warn("AnalayzeDataSet.Trimean",
			zap.Error(err))
		// return nil, err
	}

	dss.Trimean = trimean

	variance, err := statsbase.Variance(lstd)
	if err != nil {
		adacorebase.Warn("AnalayzeDataSet.Variance",
			zap.Error(err))
		// return nil, err
	}

	dss.Variance = variance

	pv, err := statsbase.PopulationVariance(lstd)
	if err != nil {
		adacorebase.Warn("AnalayzeDataSet.PopulationVariance",
			zap.Error(err))
		// return nil, err
	}

	dss.PopulationVariance = pv

	sv, err := statsbase.SampleVariance(lstd)
	if err != nil {
		adacorebase.Warn("AnalayzeDataSet.SampleVariance",
			zap.Error(err))
		// return nil, err
	}

	dss.SampleVariance = sv

	if dss.Mean != 0 {
		dss.COV = dss.StandardDeviation / dss.Mean
	}

	// entropy, err := statsbase.Entropy(lstd)
	// if err != nil {
	// 	return nil, err
	// }

	// dss.Entropy = entropy

	for _, v := range lstd {
		ov := v - dss.Mean
		if ov < 0 {
			ov = -ov
		}

		mdv := ov / dss.StandardDeviation
		if mdv < 1 {
			dss.MeanSDev1++
		} else if mdv < 2 {
			dss.MeanSDev2++
		} else if mdv < 3 {
			dss.MeanSDev3++
		}
	}

	return dss, nil
}
