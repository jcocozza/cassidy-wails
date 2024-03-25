package measurement

import "fmt"

type UnitClass string
type Unit string
type PaceUnit string
type UnitType string
const (
	Imperial UnitClass = "imperial"
	Metric UnitClass = "metric"
	StandardUnit Unit = "m"
	Distance UnitType = "distance"
	Vertical UnitType = "vertical"
	Minute UnitType = "minute"
	Hour UnitType = "hour"

	Feet Unit  = "ft"
	Yard Unit = "yd"
	Meter Unit = "m"
	Mile Unit = "mi"
	Kilometer Unit = "km"

	MilesPerHour PaceUnit = "mph"
	KilometerPerHour PaceUnit = "kmph"
	MintuesPer100Meters PaceUnit = "min/100m"
	MintuesPer100Yards PaceUnit = "min/100yd"
	MinutesPerMile PaceUnit = "min/mile"
	MinutesPerKilometer PaceUnit = "min/km"

	secondsToHours = (1.0/60/60)
	secondsToMinutes = (1.0/60)

	kmToMi = 0.621371
	kmToM = 1000
	kmToYd = 1093.613
	kmToft = 3280.84

 	mToKm = 0.001
 	mToMi = 0.000621371
	mToYd = 1.093613
	mToft = 3.28084

	miToKm = 1.609344
	miToM = 1609.344
	miToYd = 1760
	miToft = 5280

	ydToM = 0.9144
	ydToKm = 0.000914
	ydToMi = 0.0005681818
	ydToft = 3

	ftToM = 0.3048
	ftToKm = 0.0003048
	ftToMi = 0.0001893939
	ftToYd = 0.3333333333
)
// Validate a unit.
//
// This just checks to ensure that the unit is in the known units described in the consts.
func ValidateUnit(unit Unit) error {
	if unit == "" {
		return fmt.Errorf("empty unit is not valid")
	}
	if unit != Feet && unit != Yard && unit != Mile && unit != Kilometer && unit != Meter {
		return fmt.Errorf("unknown unit")
	}

	return nil
}
// Validate a unit class.
//
// This just checks to ensure that a unit class is either imperial or metric.
func ValidateUnitClass(unitClass UnitClass) error {
	if unitClass == "" {
		return fmt.Errorf("empty unit class is not valid")
	}
	if unitClass != Imperial && unitClass != Metric {
		return fmt.Errorf("invalid unit class")
	}
	return nil
}

// return units for given unit classes
//
// unitClass: imperial, metric
//
// unitType: vertical, distance
func UnitClassControl(unitClass UnitClass, unitType UnitType) Unit {
	switch unitClass {
	case Imperial:
		switch unitType {
		case Vertical:
			return Feet
		case Distance:
			return Mile
		}
	case Metric:
		switch unitType {
		case Vertical:
			return Meter
		case Distance:
			return Kilometer
		}
	}
	// this better not panic
	panic("unknown unit")
}
// Return the distance and vertical units for a unit class
func DistanceVerticalFromClass(unitClass UnitClass) (Unit, Unit) {
	switch unitClass {
	case Imperial:
		return Mile, Feet
	case Metric:
		return Kilometer, Meter
	}
	// this better not panic
	panic("unknown unit class")
}