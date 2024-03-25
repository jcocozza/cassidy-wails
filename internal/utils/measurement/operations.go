package measurement

import (
	"fmt"
	"math"
)

// Subtract two measurements. (The first minus the second).
//
// This operation assumes the length of the unit is in the standard length, thus it does NOT convert units.
//
// That said, the two measurements must be of the same unit to combine them.
//
// e.g. thus function works under the assumption that everything is in SI units and merely contains information about what unit it should be converted to.
func DifferenceSI(m1 *Measurement, m2 *Measurement) (*Measurement, error) {
	if m1.Unit != m2.Unit {
		return nil, fmt.Errorf("failed to subtract measurements")
	}
	return &Measurement{
		Unit:   m1.Unit,
		Length: m1.Length - m2.Length,
	}, nil
}

// Calculate the percent change from the current measurement to the previous measurement to the second. Return the result in % form.
//
// e.g. curr = 2, prev = 1 => 100%.
//
// This operation assumes the length of the unit is in the standard length, thus it does NOT convert units.
//
// That said, the two measurements must be of the same unit to combine them.
//
// e.g. thus function works under the assumption that everything is in SI units and merely contains information about what unit it should be converted to.
func PercentChange(currM *Measurement, prevM *Measurement) (float64, error) {
	if currM.Unit != prevM.Unit {
		return math.NaN(), fmt.Errorf("failed to calculate percent change in measurement")
	}

	if prevM.Length == 0 && currM.Length != 0 {
		return math.NaN(), nil
	} else if prevM.Length == 0 && currM.Length == 0 {
		return 0, nil
	}
	return RoundToHundredth(((currM.Length - prevM.Length) / math.Abs(prevM.Length)) * 100), nil
}
// Calculate the average length of a list of measurements
//
// Error if the length is 0 or they have different units
func AverageMeasurement(mList []*Measurement) (*Measurement, error) {
	if len(mList) == 0 {
		return nil, fmt.Errorf("unable to calcualte average measurment of 0 measurements")
	}
	u1 := mList[0].Unit
	sum := 0.0
	for _, measurement := range mList {
		if measurement.Unit != u1 {
			return nil, fmt.Errorf("failed to calculate average measurement: must have same units")
		} else {
			sum += measurement.Length
		}
	}
	average := sum / float64(len(mList))
	return &Measurement{
		Length: average,
		Unit: u1,
	}, nil
}
// For a measurement that represents an sum, compute its average
func CalculateAverageMeasurement(numMeasures int, measurement *Measurement) (*Measurement) {
	return &Measurement{
		Length: measurement.Length/float64(numMeasures),
		Unit: measurement.Unit,
	}
}