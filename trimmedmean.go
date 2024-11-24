package trimmedmean

import (
	"errors"
	"sort"
)

// TrimmedMean calculates the trimmed mean of a slice of numbers.
func TrimmedMean(data []float64, trimProportion ...float64) (float64, error) {
	if len(data) == 0 {
		return 0, errors.New("data slice is empty")
	}

	// Determine the trimming proportions
	lowerTrim, upperTrim := 0.0, 0.0
	if len(trimProportion) > 0 {
		lowerTrim = trimProportion[0]
		upperTrim = lowerTrim
	}
	if len(trimProportion) > 1 {
		upperTrim = trimProportion[1]
	}

	if lowerTrim < 0 || upperTrim < 0 || lowerTrim+upperTrim >= 1 {
		return 0, errors.New("invalid trimming proportions")
	}

	// Sort data
	sortedData := append([]float64{}, data...)
	sort.Float64s(sortedData)

	// Calculate indices
	n := len(sortedData)
	start := int(float64(n) * lowerTrim)
	end := n - int(float64(n)*upperTrim)

	if start >= end {
		return 0, errors.New("trimming removed all data points")
	}

	trimmed := sortedData[start:end]
	sum := 0.0
	for _, value := range trimmed {
		sum += value
	}

	return sum / float64(len(trimmed)), nil
}
