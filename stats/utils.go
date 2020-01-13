package stats

// CanHarmonicMean - can HarmonicMean
func CanHarmonicMean(lst []float64) bool {
	for _, v := range lst {
		if v == 0 {
			return false
		}
	}

	return true
}
