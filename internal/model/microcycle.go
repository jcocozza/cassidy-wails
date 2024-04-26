package model

import (
	"math"
	"time"

	"github.com/jcocozza/cassidy-wails/internal/utils"
	"github.com/jcocozza/cassidy-wails/internal/utils/measurement"
)

// The totals for a microcycle.
type Totals struct {
	TotalPlannedDistance *measurement.Measurement `json:"total_planned_distance"`
	TotalPlannedDuration float64                  `json:"total_planned_duration"`
	TotalPlannedVertical *measurement.Measurement `json:"total_planned_vertical"`

	TotalCompletedDistance *measurement.Measurement `json:"total_completed_distance"`
	TotalCompletedDuration float64                  `json:"total_completed_duration"`
	TotalCompletedVertical *measurement.Measurement `json:"total_completed_vertical"`
}

// An empty total has all standard units for planned and completed as well as 0 duration
func EmptyTotals(userUnitClass measurement.UnitClass) *Totals {
	distanceUnit, verticalUnit := measurement.DistanceVerticalFromClass(userUnitClass)
	return &Totals{
		TotalPlannedDistance: measurement.CreateMeasurement(distanceUnit, 0),
		TotalPlannedDuration: 0,
		TotalPlannedVertical: measurement.CreateMeasurement(verticalUnit, 0),

		TotalCompletedDistance: measurement.CreateMeasurement(distanceUnit, 0),
		TotalCompletedDuration: 0,
		TotalCompletedVertical: measurement.CreateMeasurement(verticalUnit, 0),
	}
}

// Contains actual and percent changes in totals between two microcycles
type TotalsDifferences struct {
	DifferencePlannedDistance *measurement.Measurement `json:"difference_planned_distance"`
	PercentPlannedDistance    utils.JsonFloat          `json:"percent_difference_planned_distance"`

	DifferencePlannedDuration float64         `json:"difference_planned_duration"`
	PercentPlannedDuration    utils.JsonFloat `json:"percent_difference_planned_duration"`

	DifferencePlannedVertical *measurement.Measurement `json:"difference_planned_vertical"`
	PercentPlannedVertical    utils.JsonFloat          `json:"percent_difference_planned_vertical"`

	DifferenceCompletedDistance *measurement.Measurement `json:"difference_completed_distance"`
	PercentCompletedDistance    utils.JsonFloat          `json:"percent_difference_completed_distance"`

	DifferenceCompletedDuration float64         `json:"difference_completed_duration"`
	PercentCompletedDuration    utils.JsonFloat `json:"percent_difference_completed_duration"`

	DifferenceCompletedVertical *measurement.Measurement `json:"difference_completed_vertical"`
	PercentCompletedVertical    utils.JsonFloat          `json:"percent_difference_completed_vertical"`
}

// Totals by activity type for a given microcycle.
type TotalByActivityType struct {
	ActivityType         *ActivityType            `json:"activity_type"`
	TotalPlannedDistance *measurement.Measurement `json:"total_planned_distance"`
	TotalPlannedDuration float64                  `json:"total_planned_duration"`
	TotalPlannedVertical *measurement.Measurement `json:"total_planned_vertical"`
	PlannedPace          string                   `json:"planned_pace"`

	TotalCompletedDistance *measurement.Measurement `json:"total_completed_distance"`
	TotalCompletedDuration float64                  `json:"total_completed_duration"`
	TotalCompletedVertical *measurement.Measurement `json:"total_completed_vertical"`
	CompletedPace          string                   `json:"completed_pace"`
}

// An empty total by activity type is an empty activity type, standard units with 0 length and 0 duration.
func EmptyTotalByActivityType(userUnitClass measurement.UnitClass) *TotalByActivityType {
	distanceUnit, verticalUnit := measurement.DistanceVerticalFromClass(userUnitClass)
	return &TotalByActivityType{
		ActivityType:         EmptyActivityType(),
		TotalPlannedDistance: measurement.CreateMeasurement(distanceUnit, 0),
		TotalPlannedDuration: 0,
		TotalPlannedVertical: measurement.CreateMeasurement(verticalUnit, 0),

		TotalCompletedDistance: measurement.CreateMeasurement(distanceUnit, 0),
		TotalCompletedDuration: 0,
		TotalCompletedVertical: measurement.CreateMeasurement(verticalUnit, 0),
	}
}

// calculate total planned and completed paces
func (tbat *TotalByActivityType) CalculatePaces(userUnitClass measurement.UnitClass) {
	paceUnit := measurement.PaceUnitByActivityType(tbat.ActivityType.Id, userUnitClass)

	plannedPace := measurement.CalculatePace(tbat.TotalPlannedDistance, tbat.TotalPlannedDuration, paceUnit)
	completedPace := measurement.CalculatePace(tbat.TotalCompletedDistance, tbat.TotalCompletedDuration, paceUnit)

	tbat.PlannedPace = plannedPace.String()
	tbat.CompletedPace = completedPace.String()
}

// Contains actual and percent changes in totals by activity type between two microcycles
type TotalByActivityTypeDifferences struct {
	ActivityType              *ActivityType            `json:"activity_type"`
	DifferencePlannedDistance *measurement.Measurement `json:"difference_planned_distance"`
	PercentPlannedDistance    utils.JsonFloat          `json:"percent_difference_planned_distance"`

	DifferencePlannedDuration float64         `json:"difference_planned_duration"`
	PercentPlannedDuration    utils.JsonFloat `json:"percent_difference_planned_duration"`

	DifferencePlannedVertical *measurement.Measurement `json:"difference_planned_vertical"`
	PercentPlannedVertical    utils.JsonFloat          `json:"percent_difference_planned_vertical"`

	DifferenceCompletedDistance *measurement.Measurement `json:"difference_completed_distance"`
	PercentCompletedDistance    utils.JsonFloat          `json:"percent_difference_completed_distance"`

	DifferenceCompletedDuration float64         `json:"difference_completed_duration"`
	PercentCompletedDuration    utils.JsonFloat `json:"percent_difference_completed_duration"`

	DifferenceCompletedVertical *measurement.Measurement `json:"difference_completed_vertical"`
	PercentCompletedVertical    utils.JsonFloat          `json:"percent_difference_completed_vertical"`
}

// The set of totals for an activity type on a given date.
type TotalByActivityTypeAndDate struct {
	ActivityType         *ActivityType            `json:"activity_type"`
	Date                 time.Time			      `json:"date" ts_type:"Date" ts_transform:"new Date(__VALUE__)"`
	TotalPlannedDistance *measurement.Measurement `json:"total_planned_distance"`
	TotalPlannedDuration float64                  `json:"total_planned_duration"`
	TotalPlannedVertical *measurement.Measurement `json:"total_planned_vertical"`

	TotalCompletedDistance *measurement.Measurement `json:"total_completed_distance"`
	TotalCompletedDuration float64                  `json:"total_completed_duration"`
	TotalCompletedVertical *measurement.Measurement `json:"total_completed_vertical"`
}

// An empty total by activity type and date is an empty activity type, empty date, and standard measurements with 0 length and 0 duration.
func EmptyTotalByActivityTypeAndDate(userUnitClass measurement.UnitClass) *TotalByActivityTypeAndDate {
	distanceUnit, verticalUnit := measurement.DistanceVerticalFromClass(userUnitClass)
	return &TotalByActivityTypeAndDate{
		ActivityType:         EmptyActivityType(),
		Date:                 time.Now(),
		TotalPlannedDistance: measurement.CreateMeasurement(distanceUnit, 0),
		TotalPlannedDuration: 0,
		TotalPlannedVertical: measurement.CreateMeasurement(verticalUnit, 0),

		TotalCompletedDistance: measurement.CreateMeasurement(distanceUnit, 0),
		TotalCompletedDuration: 0,
		TotalCompletedVertical: measurement.CreateMeasurement(verticalUnit, 0),
	}
}

// The summary information of a microcycle.
type MicrocycleSummary struct {
	Totals                          *Totals                           `json:"totals"`
	PreviousTotals                  *Totals                           `json:"previous_totals"`
	AveragePreviousTotals           *Totals                           `json:"average_previous_totals"`
	TotalsDifferences               *TotalsDifferences                `json:"totals_differences"`
	WeightedTotalsDifferences       *TotalsDifferences                `json:"weighted_totals_differences"`
	TotalsByActivityType            []*TotalByActivityType            `json:"totals_by_activity_type"`
	TotalsByActivityTypeDifferences []*TotalByActivityTypeDifferences `json:"totals_by_activity_type_differences"`
	PreviousTotalsByActivityType    []*TotalByActivityType            `json:"previous_totals_by_activity_type"`
	TotalsByActivityTypeAndDate     []*TotalByActivityTypeAndDate     `json:"totals_by_activity_type_and_date"`
}

// Calculate changes in from previous cycle to current cycle
//
// Sets the relevant fields in the microcycle summary or throws an error
func (ms *MicrocycleSummary) CalculateCycleChanges() error {
	// total changes
	plannedTotalDistanceChange, err := measurement.DifferenceSI(ms.Totals.TotalPlannedDistance, ms.PreviousTotals.TotalPlannedDistance)
	if err != nil {
		return err
	}
	plannedTotalDistancePercentChange, _ := measurement.PercentChange(ms.Totals.TotalPlannedDistance, ms.PreviousTotals.TotalPlannedDistance) // we can ignore error b/c it would have triggered in the previous function
	plannedTotalDurationChange := ms.Totals.TotalPlannedDuration - ms.PreviousTotals.TotalPlannedDuration
	plannedTotalDurationPercentChange := utils.PercentChange(ms.Totals.TotalPlannedDuration, ms.PreviousTotals.TotalPlannedDuration)
	plannedTotalVerticalChange, err1 := measurement.DifferenceSI(ms.Totals.TotalPlannedVertical, ms.PreviousTotals.TotalPlannedVertical)
	if err1 != nil {
		return err1
	}
	plannedTotalVerticalPercentChange, _ := measurement.PercentChange(ms.Totals.TotalPlannedVertical, ms.PreviousTotals.TotalPlannedVertical)

	completedTotalDistanceChange, err2 := measurement.DifferenceSI(ms.Totals.TotalCompletedDistance, ms.PreviousTotals.TotalCompletedDistance)
	if err2 != nil {
		return err2
	}
	completedTotalDistancePercentChange, _ := measurement.PercentChange(ms.Totals.TotalCompletedDistance, ms.PreviousTotals.TotalCompletedDistance)
	completedTotalDurationChange := ms.Totals.TotalCompletedDuration - ms.PreviousTotals.TotalCompletedDuration
	completedTotalDurationPercentChange := utils.PercentChange(ms.Totals.TotalCompletedDuration, ms.PreviousTotals.TotalCompletedDuration)
	completedTotalVerticalChange, err3 := measurement.DifferenceSI(ms.Totals.TotalCompletedVertical, ms.PreviousTotals.TotalCompletedVertical)
	if err3 != nil {
		return err3
	}
	completedTotalVerticalPercentChange, _ := measurement.PercentChange(ms.Totals.TotalCompletedVertical, ms.PreviousTotals.TotalCompletedVertical)

	totalChanges := &TotalsDifferences{
		DifferencePlannedDistance: plannedTotalDistanceChange,
		PercentPlannedDistance:    utils.JsonFloat(plannedTotalDistancePercentChange),
		DifferencePlannedDuration: plannedTotalDurationChange,
		PercentPlannedDuration:    utils.JsonFloat(plannedTotalDurationPercentChange),
		DifferencePlannedVertical: plannedTotalVerticalChange,
		PercentPlannedVertical:    utils.JsonFloat(plannedTotalVerticalPercentChange),

		DifferenceCompletedDistance: completedTotalDistanceChange,
		PercentCompletedDistance:    utils.JsonFloat(completedTotalDistancePercentChange),
		DifferenceCompletedDuration: completedTotalDurationChange,
		PercentCompletedDuration:    utils.JsonFloat(completedTotalDurationPercentChange),
		DifferenceCompletedVertical: completedTotalVerticalChange,
		PercentCompletedVertical:    utils.JsonFloat(completedTotalVerticalPercentChange),
	}

	// weighted total changes
	weightedPlannedTotalDistanceChange, err := measurement.DifferenceSI(ms.Totals.TotalPlannedDistance, ms.AveragePreviousTotals.TotalPlannedDistance)
	if err != nil {
		return err
	}
	weightedPlannedTotalDistancePercentChange, _ := measurement.PercentChange(ms.Totals.TotalPlannedDistance, ms.AveragePreviousTotals.TotalPlannedDistance) // we can ignore error b/c it would have triggered in the previous function
	weightedPlannedTotalDurationChange := ms.Totals.TotalPlannedDuration - ms.AveragePreviousTotals.TotalPlannedDuration
	weightedPlannedTotalDurationPercentChange := utils.PercentChange(ms.Totals.TotalPlannedDuration, ms.AveragePreviousTotals.TotalPlannedDuration)
	weightedPlannedTotalVerticalChange, err1 := measurement.DifferenceSI(ms.Totals.TotalPlannedVertical, ms.AveragePreviousTotals.TotalPlannedVertical)
	if err1 != nil {
		return err1
	}
	weightedPlannedTotalVerticalPercentChange, _ := measurement.PercentChange(ms.Totals.TotalPlannedVertical, ms.AveragePreviousTotals.TotalPlannedVertical)

	weightedCompletedTotalDistanceChange, err2 := measurement.DifferenceSI(ms.Totals.TotalCompletedDistance, ms.AveragePreviousTotals.TotalCompletedDistance)
	if err2 != nil {
		return err2
	}
	weightedCompletedTotalDistancePercentChange, _ := measurement.PercentChange(ms.Totals.TotalCompletedDistance, ms.AveragePreviousTotals.TotalCompletedDistance)
	weightedCompletedTotalDurationChange := ms.Totals.TotalCompletedDuration - ms.AveragePreviousTotals.TotalCompletedDuration
	weightedCompletedTotalDurationPercentChange := utils.PercentChange(ms.Totals.TotalCompletedDuration, ms.AveragePreviousTotals.TotalCompletedDuration)
	weightedCompletedTotalVerticalChange, err3 := measurement.DifferenceSI(ms.Totals.TotalCompletedVertical, ms.AveragePreviousTotals.TotalCompletedVertical)
	if err3 != nil {
		return err3
	}
	weightedCompletedTotalVerticalPercentChange, _ := measurement.PercentChange(ms.Totals.TotalCompletedVertical, ms.AveragePreviousTotals.TotalCompletedVertical)

	weightedTotalChanges := &TotalsDifferences{
		DifferencePlannedDistance: weightedPlannedTotalDistanceChange,
		PercentPlannedDistance:    utils.JsonFloat(weightedPlannedTotalDistancePercentChange),
		DifferencePlannedDuration: weightedPlannedTotalDurationChange,
		PercentPlannedDuration:    utils.JsonFloat(weightedPlannedTotalDurationPercentChange),
		DifferencePlannedVertical: weightedPlannedTotalVerticalChange,
		PercentPlannedVertical:    utils.JsonFloat(weightedPlannedTotalVerticalPercentChange),

		DifferenceCompletedDistance: weightedCompletedTotalDistanceChange,
		PercentCompletedDistance:    utils.JsonFloat(weightedCompletedTotalDistancePercentChange),
		DifferenceCompletedDuration: weightedCompletedTotalDurationChange,
		PercentCompletedDuration:    utils.JsonFloat(weightedCompletedTotalDurationPercentChange),
		DifferenceCompletedVertical: weightedCompletedTotalVerticalChange,
		PercentCompletedVertical:    utils.JsonFloat(weightedCompletedTotalVerticalPercentChange),
	}

	// changes by activity type
	totalsByActivityTypeChanges := []*TotalByActivityTypeDifferences{}
	for _, tat := range ms.TotalsByActivityType {
		// in the case of the activity type not existing the "increase" is just the total of the current
		tmpTotalDifferencesByActivityType := &TotalByActivityTypeDifferences{
			ActivityType:                tat.ActivityType,
			DifferencePlannedDistance:   tat.TotalPlannedDistance,
			PercentPlannedDistance:      utils.JsonFloat(math.NaN()),
			DifferencePlannedDuration:   tat.TotalPlannedDuration,
			PercentPlannedDuration:      utils.JsonFloat(math.NaN()),
			DifferencePlannedVertical:   tat.TotalPlannedVertical,
			PercentPlannedVertical:      utils.JsonFloat(math.NaN()),
			DifferenceCompletedDistance: tat.TotalCompletedDistance,
			PercentCompletedDistance:    utils.JsonFloat(math.NaN()),
			DifferenceCompletedDuration: tat.TotalCompletedDuration,
			PercentCompletedDuration:    utils.JsonFloat(math.NaN()),
			DifferenceCompletedVertical: tat.TotalCompletedVertical,
			PercentCompletedVertical:    utils.JsonFloat(math.NaN()),
		}
		for _, ptat := range ms.PreviousTotalsByActivityType {
			// only interested in calculating differences for same activity types
			// each activity type can only appear once
			if tat.ActivityType == ptat.ActivityType {
				tmpTotalPlannedDistanceActTypeChange, err := measurement.DifferenceSI(tat.TotalPlannedDistance, ptat.TotalPlannedDistance)
				if err != nil {
					return err
				}
				tmpTotalPlannedDistanceActTypePercentChange, _ := measurement.PercentChange(tat.TotalPlannedDistance, ptat.TotalPlannedDistance)
				tmpTotalPlannedDurationActTypeChange := tat.TotalPlannedDuration - ptat.TotalPlannedDuration
				tmpTotalPlannedDurationActTypePercentChange := utils.PercentChange(tat.TotalPlannedDuration, ptat.TotalPlannedDuration)
				tmpTotalPlannedVerticalActTypeChange, err1 := measurement.DifferenceSI(tat.TotalPlannedVertical, ptat.TotalPlannedVertical)
				if err1 != nil {
					return err1
				}
				tmpTotalPlannedVerticalActTypePercentChange, _ := measurement.PercentChange(tat.TotalPlannedVertical, ptat.TotalPlannedVertical)

				tmpTotalCompletedDistanceActTypeChange, err2 := measurement.DifferenceSI(tat.TotalCompletedDistance, ptat.TotalCompletedDistance)
				if err2 != nil {
					return err2
				}
				tmpTotalCompletedDistanceActTypePercentChange, _ := measurement.PercentChange(tat.TotalCompletedDistance, ptat.TotalCompletedDistance)
				tmpTotalCompletedDurationActTypeChange := tat.TotalCompletedDuration - ptat.TotalCompletedDuration
				tmpTotalCompletedDurationActTypePercentChange := utils.PercentChange(tat.TotalCompletedDuration, ptat.TotalCompletedDuration)
				tmpTotalCompletedVerticalActTypeChange, err3 := measurement.DifferenceSI(tat.TotalCompletedVertical, ptat.TotalCompletedVertical)
				if err3 != nil {
					return err3
				}
				tmpTotalCompletedVerticalActTypePercentChange, _ := measurement.PercentChange(tat.TotalCompletedVertical, ptat.TotalCompletedVertical)

				tmpTotalDifferencesByActivityType.DifferencePlannedDistance = tmpTotalPlannedDistanceActTypeChange
				tmpTotalDifferencesByActivityType.PercentPlannedDistance = utils.JsonFloat(tmpTotalPlannedDistanceActTypePercentChange)

				tmpTotalDifferencesByActivityType.DifferencePlannedDuration = tmpTotalPlannedDurationActTypeChange
				tmpTotalDifferencesByActivityType.PercentPlannedDuration = utils.JsonFloat(tmpTotalPlannedDurationActTypePercentChange)

				tmpTotalDifferencesByActivityType.DifferencePlannedVertical = tmpTotalPlannedVerticalActTypeChange
				tmpTotalDifferencesByActivityType.PercentPlannedVertical = utils.JsonFloat(tmpTotalPlannedVerticalActTypePercentChange)

				tmpTotalDifferencesByActivityType.DifferenceCompletedDistance = tmpTotalCompletedDistanceActTypeChange
				tmpTotalDifferencesByActivityType.PercentCompletedDistance = utils.JsonFloat(tmpTotalCompletedDistanceActTypePercentChange)

				tmpTotalDifferencesByActivityType.DifferenceCompletedDuration = tmpTotalCompletedDurationActTypeChange
				tmpTotalDifferencesByActivityType.PercentCompletedDuration = utils.JsonFloat(tmpTotalCompletedDurationActTypePercentChange)

				tmpTotalDifferencesByActivityType.DifferenceCompletedVertical = tmpTotalCompletedVerticalActTypeChange
				tmpTotalDifferencesByActivityType.PercentCompletedVertical = utils.JsonFloat(tmpTotalCompletedVerticalActTypePercentChange)
			}
		}
		totalsByActivityTypeChanges = append(totalsByActivityTypeChanges, tmpTotalDifferencesByActivityType)
	}

	ms.TotalsDifferences = totalChanges
	ms.WeightedTotalsDifferences = weightedTotalChanges
	ms.TotalsByActivityTypeDifferences = totalsByActivityTypeChanges
	return nil
}

// Represents a set number of days (typically between 1 week and 10 days).
type Microcycle struct {
	StartDate       time.Time          `json:"start_date" ts_type:"Date" ts_transform:"new Date(__VALUE__)"`
	EndDate         time.Time          `json:"end_date" ts_type:"Date" ts_transform:"new Date(__VALUE__)"`
	CycleActivities *Cycle             `json:"cycle_activities"`
	Summary         *MicrocycleSummary `json:"summary"`
}
