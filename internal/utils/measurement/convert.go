package measurement

import (
	"fmt"
	"math"
)

// Round a number to the hundreth
func RoundToHundredth(num float64) float64 {
	return math.Round(num * 100) / 100
}
// Modify a measurement in place.
//
// Takes a measurement and converts it into the outgoingUnit.
//
// If the outgoing unit is the same as the incoming measurement, this function does nothing.
func ModifyMeasurement(measurement *Measurement, incomingUnit Unit, outgoingUnit Unit) error {
	switch incomingUnit {
	case "km":
		switch outgoingUnit {
		case "mi":
			//measurement.Unit = "mi"
			measurement.Length = RoundToHundredth(measurement.Length * kmToMi)
			return nil
		case "m":
			//measurement.Unit = "m"
			measurement.Length = RoundToHundredth(measurement.Length * kmToM)
			return nil
		case "yd":
			//measurement.Unit = "yd"
			measurement.Length = RoundToHundredth(measurement.Length * kmToYd)
			return nil
		case "ft":
			//measurement.Unit = "ft"
			measurement.Length = RoundToHundredth(measurement.Length * kmToft)
			return nil
		default:
			return nil
		}
	case "m":
		switch outgoingUnit {
		case "km":
			//measurement.Unit = "km"
			measurement.Length = RoundToHundredth(measurement.Length * mToKm)
			return nil
		case "mi":
			//measurement.Unit = "mi"
			measurement.Length = RoundToHundredth(measurement.Length * mToMi)
			return nil
		case "yd":
			//measurement.Unit = "yd"
			measurement.Length = RoundToHundredth(measurement.Length * mToYd)
			return nil
		case "ft":
			//measurement.Unit = "ft"
			measurement.Length = RoundToHundredth(measurement.Length * mToft)
			return nil
		default:
			return nil
		}
	case "mi":
		switch outgoingUnit {
		case "km":
			//measurement.Unit = "km"
			measurement.Length = RoundToHundredth(measurement.Length * miToKm)
			return nil
		case "m":
			//measurement.Unit = "m"
			measurement.Length = RoundToHundredth(measurement.Length * miToM)
			return nil
		case "yd":
			//measurement.Unit = "yd"
			measurement.Length = RoundToHundredth(measurement.Length * miToYd)
			return nil
		case "ft":
			//measurement.Unit = "ft"
			measurement.Length = RoundToHundredth(measurement.Length * miToft)
			return nil
		default:
			return nil
		}
	case "ft":
		switch outgoingUnit {
		case "km":
			//measurement.Unit = "km"
			measurement.Length = RoundToHundredth(measurement.Length * ftToKm)
			return nil
		case "m":
			//measurement.Unit = "m"
			measurement.Length = RoundToHundredth(measurement.Length * ftToM)
			return nil
		case "yd":
			//measurement.Unit = "yd"
			measurement.Length = RoundToHundredth(measurement.Length * ftToYd)
			return nil
		case "mi":
			//measurement.Unit = "mi"
			measurement.Length = RoundToHundredth(measurement.Length * ftToMi)
			return nil
		default:
			return nil
		}
	case "yd":
		switch outgoingUnit {
		case "m":
			//measurement.Unit = "m"
			measurement.Length = RoundToHundredth(measurement.Length * ydToM)
			return nil
		case "km":
			//measurement.Unit = "km"
			measurement.Length = RoundToHundredth(measurement.Length * ydToKm)
			return nil
		case "mi":
			//measurement.Unit = "mi"
			measurement.Length = RoundToHundredth(measurement.Length * ydToMi)
			return nil
		case "ft":
			//measurement.Unit = "ft"
			measurement.Length = RoundToHundredth(measurement.Length * ydToft)
			return nil
		default:
			return nil
		}
	}
	return fmt.Errorf("measurement conversion failed for an unknown reason. are you sure units are correct?")
}
// Set a measurement in place. Will modify the unit
//
// Takes a measurement and converts it into the outgoingUnit.
//
// If the outgoing unit is the same as the incoming measurement, this function does nothing.
func SetMeasurement(measurement *Measurement, incomingUnit Unit, outgoingUnit Unit) error {
	switch incomingUnit {
	case "km":
		switch outgoingUnit {
		case "mi":
			measurement.Unit = "mi"
			measurement.Length = RoundToHundredth(measurement.Length * kmToMi)
			return nil
		case "m":
			measurement.Unit = "m"
			measurement.Length = RoundToHundredth(measurement.Length * kmToM)
			return nil
		case "yd":
			measurement.Unit = "yd"
			measurement.Length = RoundToHundredth(measurement.Length * kmToYd)
			return nil
		case "ft":
			measurement.Unit = "ft"
			measurement.Length = RoundToHundredth(measurement.Length * kmToft)
			return nil
		default:
			return nil
		}
	case "m":
		switch outgoingUnit {
		case "km":
			measurement.Unit = "km"
			measurement.Length = RoundToHundredth(measurement.Length * mToKm)
			return nil
		case "mi":
			measurement.Unit = "mi"
			measurement.Length = RoundToHundredth(measurement.Length * mToMi)
			return nil
		case "yd":
			measurement.Unit = "yd"
			measurement.Length = RoundToHundredth(measurement.Length * mToYd)
			return nil
		case "ft":
			measurement.Unit = "ft"
			measurement.Length = RoundToHundredth(measurement.Length * mToft)
			return nil
		default:
			return nil
		}
	case "mi":
		switch outgoingUnit {
		case "km":
			measurement.Unit = "km"
			measurement.Length = RoundToHundredth(measurement.Length * miToKm)
			return nil
		case "m":
			measurement.Unit = "m"
			measurement.Length = RoundToHundredth(measurement.Length * miToM)
			return nil
		case "yd":
			measurement.Unit = "yd"
			measurement.Length = RoundToHundredth(measurement.Length * miToYd)
			return nil
		case "ft":
			measurement.Unit = "ft"
			measurement.Length = RoundToHundredth(measurement.Length * miToft)
			return nil
		default:
			return nil
		}
	case "ft":
		switch outgoingUnit {
		case "km":
			measurement.Unit = "km"
			measurement.Length = RoundToHundredth(measurement.Length * ftToKm)
			return nil
		case "m":
			measurement.Unit = "m"
			measurement.Length = RoundToHundredth(measurement.Length * ftToM)
			return nil
		case "yd":
			measurement.Unit = "yd"
			measurement.Length = RoundToHundredth(measurement.Length * ftToYd)
			return nil
		case "mi":
			measurement.Unit = "mi"
			measurement.Length = RoundToHundredth(measurement.Length * ftToMi)
			return nil
		default:
			return nil
		}
	case "yd":
		switch outgoingUnit {
		case "m":
			measurement.Unit = "m"
			measurement.Length = RoundToHundredth(measurement.Length * ydToM)
			return nil
		case "km":
			measurement.Unit = "km"
			measurement.Length = RoundToHundredth(measurement.Length * ydToKm)
			return nil
		case "mi":
			measurement.Unit = "mi"
			measurement.Length = RoundToHundredth(measurement.Length * ydToMi)
			return nil
		case "ft":
			measurement.Unit = "ft"
			measurement.Length = RoundToHundredth(measurement.Length * ydToft)
			return nil
		default:
			return nil
		}
	}
	return fmt.Errorf("measurement conversion failed for an unknown reason. are you sure units are correct?")
}

// Modify a measurement in place so that its magnitude is reflected in SI units
func IntoDatabaseConversion(measurement *Measurement) error {
	err := ModifyMeasurement(measurement, measurement.Unit, StandardUnit)
	return err
}
// Modify a measurement in place so that its magnitude is reflected in its stored units
func LeavingDatabaseConversion(measurement *Measurement) error {
	err := ModifyMeasurement(measurement, StandardUnit, measurement.Unit)
	return err
}
// Modify a measurement given a user unit class; Also will modify the UNIT
func LeavingDatabaseConversionUnitClass(measurement *Measurement, userUnitClass UnitClass, unitType UnitType) error {
	unit := UnitClassControl(userUnitClass, unitType)
	err := SetMeasurement(measurement, StandardUnit, unit)
	return err
}