package adacore

// NewRateOfChgInt - new
func NewRateOfChgInt(ival []int) []float32 {
	if len(ival) < 2 {
		return nil
	}

	var lstper []float32

	lastival := ival[0]
	for i := 1; i < len(ival); i++ {
		lstper = append(lstper, float32(ival[i])/float32(lastival))
		lastival = ival[i]
	}

	return lstper
}

// NewRateOfChgFloat - new
func NewRateOfChgFloat(fval []float32) []float32 {
	if len(fval) < 2 {
		return nil
	}

	var lstper []float32

	lastival := fval[0]
	for i := 1; i < len(fval); i++ {
		lstper = append(lstper, fval[i]/lastival)
		lastival = fval[i]
	}

	return lstper
}

// NewRateOfChgFloat64 - new
func NewRateOfChgFloat64(fval []float64) []float32 {
	if len(fval) < 2 {
		return nil
	}

	var lstper []float32

	lastival := fval[0]
	for i := 1; i < len(fval); i++ {
		lstper = append(lstper, float32(fval[i]/lastival))
		lastival = fval[i]
	}

	return lstper
}
