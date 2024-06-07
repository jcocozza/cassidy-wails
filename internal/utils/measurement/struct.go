package measurement

import (
	"fmt"
)

// A length of units.
type Measurement struct {
	Unit   Unit  `json:"unit"`
	Length float64 `json:"length"`
}
// Empty Measurement is a measurement with no unit and a length of -1.
func EmptyMeasurement() *Measurement {
	return &Measurement{
		Unit: "",
		Length: -1,
	}
}
// The default measurement is 0 meters.
func DefaultMeasurement() *Measurement {
	return &Measurement{
		Unit: Meter,
		Length: 0,
	}
}
// Return the zero length for a passed measurement.
func ZeroLength(unit Unit) *Measurement {
	return &Measurement{
		Unit: unit,
		Length: 0,
	}
}
// Return a measurement with passed units and length.
func CreateMeasurement(unit Unit, length float64) *Measurement {
	return &Measurement{
		Unit: unit,
		Length: length,
	}
}
// Return a measurement with passed length and standard unit
func StandardMeasurement(length float64) *Measurement {
	return &Measurement{
		Unit: StandardUnit,
		Length: length,
	}
}
// Validate a measurement
//
// A valid measurement has a valid unit and a length of greater than or equal to 0
func (m *Measurement) Validate() error {
	err := ValidateUnit(m.Unit)
	if err != nil {
		return fmt.Errorf("unit is invalid")
	}
	if m.Length < 0 {
		return fmt.Errorf("length cannot be negative")
	}
	return nil
}

type Pace struct {
	Unit  PaceUnit `json:"unit"`
	Speed float64  `json:"speed"`
}
func NewPace(unit PaceUnit, speed float64) *Pace {
	return &Pace{
		Unit: unit,
		Speed: speed,
	}
}
// take a measurment (with a length in meters) and a duration in seconds and convert that into a duration.
//
// Return 0 for speed if the distance or duration is 0 (the alternative in N/A, but 0 is good enough for now)
func CalculatePace(distance *Measurement, duration float64, paceUnit PaceUnit) *Pace {

	if distance.Length == 0 {
		return &Pace{
			Unit: paceUnit,
			Speed: 0,
		}
	}
	if duration == 0 {
		return &Pace{
			Unit: paceUnit,
			Speed: 0,
		}
	}

	switch paceUnit {
	case MilesPerHour:
		hour := duration * secondsToHours
		return &Pace{
			Speed: RoundToHundredth((distance.Length * mToMi) / hour),
			Unit: MilesPerHour,
		}
	case KilometerPerHour:
		hour := duration * secondsToHours
		return &Pace{
			Speed: RoundToHundredth((distance.Length * mToKm) / hour),
			Unit: KilometerPerHour,
		}
	case MintuesPer100Meters:
		minutes := duration * secondsToMinutes
		return &Pace{
			Speed: RoundToHundredth((minutes * 100) / distance.Length),
			Unit: MintuesPer100Meters,
		}
	case MintuesPer100Yards:
		minutes := duration * secondsToMinutes
		return &Pace{
			Speed: RoundToHundredth((minutes * 100) / (distance.Length * mToYd)),
			Unit: MintuesPer100Yards,
		}
	case MinutesPerMile:
		minutes := duration * secondsToMinutes
		return &Pace{
			Speed: RoundToHundredth((minutes / (distance.Length * mToMi))),
			Unit: MinutesPerMile,
		}
	case MinutesPerKilometer:
		minutes := duration * secondsToMinutes
		return &Pace{
			Speed: RoundToHundredth((minutes / (distance.Length * mToKm))),
			Unit: MinutesPerKilometer,
		}
	}
	panic("invaid unit")
}
// Convert a pace to a string
// TODO Fix this to properly handle non minute/mile values (there should be no 10:00 mph)
func (p *Pace) String() string {

		if p.Unit == MilesPerHour || p.Unit == KilometerPerHour {
			paceString := fmt.Sprintf("%.2f", p.Speed)
			return paceString + " " + string(p.Unit)
		}

		// Calculate the minutes and seconds
		minutes := int(p.Speed)
		seconds := int((p.Speed - float64(minutes)) * 60)

		paceString := fmt.Sprintf("%d:%02d", minutes, seconds)
		return paceString + " " + string(p.Unit)
}
// The id's of activity types
// TODO: FIX THIS TO DEPEND ON THE ACTIVITY TYPE FILE
const (
	Run = 1
	Bike = 2
	Swim = 3
	Hike = 4
	RestDay = 5
	Strength = 6
	Recovery = 7
)
// return the pace associated with an activity type id
func PaceUnitByActivityType(activityTypeId int, userUnitClass UnitClass) PaceUnit {
	paceUnit := PaceUnit("")
	switch activityTypeId {
	case Bike:
		switch userUnitClass {
		case Imperial:
			paceUnit = MilesPerHour
		case Metric:
			paceUnit = KilometerPerHour
		}
	case Swim:
		switch userUnitClass {
		case Imperial:
			paceUnit = MintuesPer100Yards
		case Metric:
			paceUnit = MintuesPer100Meters
		}
    default: // Run, Hike, RestDay, Strength, Recovery and others
		switch userUnitClass {
		case Imperial:
			paceUnit = MinutesPerMile
		case Metric:
			paceUnit = MinutesPerKilometer
		}
    }
	return paceUnit
}
func PaceUnitByDistanceUnit(activityTypeId int, distanceUnit Unit) PaceUnit {
	paceUnit := PaceUnit("")

	switch distanceUnit {
	case Meter:
		paceUnit = MintuesPer100Meters
	case Yard:
		paceUnit = MintuesPer100Yards
	
	case Kilometer:
		switch activityTypeId {
		case Bike:
			paceUnit = KilometerPerHour
		default:
			paceUnit = MinutesPerKilometer
		}
    default: // Mile is default	
		switch activityTypeId {
		case Bike:
			paceUnit = MilesPerHour
		default:
			paceUnit = MinutesPerMile
		}
    }
	return paceUnit
}
