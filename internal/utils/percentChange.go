package utils

import (
	"math"

	"github.com/jcocozza/cassidy-wails/internal/utils/measurement"
)

// Calcaulte the percent change between to floats.
//
// from previous TO current
func PercentChange(current float64, previous float64) float64 {
	if previous == 0 && current != 0 {
		return math.NaN()
	} else if previous == 0 && current == 0 {
		return 0
	}
	return measurement.RoundToHundredth(((current - previous) / math.Abs(previous)) * 100)
}
