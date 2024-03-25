package conversionrepo

import (
	"fmt"

	"github.com/jcocozza/cassidy-wails/internal/model"
	"github.com/jcocozza/cassidy-wails/internal/utils/measurement"
)

type direction string

const (
	Incoming direction = "incoming"
	Outgoing direction = "outgoing"
)

type MeasurementRepository interface {
	ConvertActivity(direction, *model.Activity, measurement.UnitClass) error

	ConvertPlanned(direction, *model.Planned) error
	ConvertCompleted(direction, *model.Completed) error

	ConvertEquipment(direction, *model.Equipment) error
	ConvertActivityEquipment(direction, *model.ActivityEquipment) error

	ConvertMicrocycle(direction, *model.Microcycle, measurement.UnitClass) error
	ConvertMicrocycleSummary(direction, *model.MicrocycleSummary, measurement.UnitClass) error
	ConvertTotals(direction, *model.Totals, measurement.UnitClass) error
	ConvertTotalsDifference(direction, *model.TotalsDifferences, measurement.UnitClass) error
	ConvertTotalByActivityType(direction, *model.TotalByActivityType, measurement.UnitClass) error
	ConvertTotalByActivityTypeDifference(direction, *model.TotalByActivityTypeDifferences, measurement.UnitClass) error
	ConvertTotalByActivityTypeAndDate(direction, *model.TotalByActivityTypeAndDate, measurement.UnitClass) error

	ConvertNCycleSummary(direction, *model.NCycleSummary, measurement.UnitClass) error
}

// This repository handles conversions between incoming and outgoing data.
//
// The general principle is that data should be converted into SI units when it arrives
// and converted back to its proper unit when it leaves.
type IMeasurementRepository struct{}

func NewIMeasurementRepository() *IMeasurementRepository {
	return &IMeasurementRepository{}
}

// Convert an activity
func (mr *IMeasurementRepository) ConvertActivity(direction direction, activity *model.Activity, userUnitClass measurement.UnitClass) error {

	// before conversion, we calculate activity paces
	activity.CalculatePace(userUnitClass)
	color, err := activity.CompletionColor()
	if err != nil {
		return fmt.Errorf("failed to get completiong color: %w", err)
	}
	activity.Color = color

	err1 := mr.ConvertPlanned(direction, activity.Planned)
	if err1 != nil {
		return fmt.Errorf("failed to convert planned during activity conversion: %w", err1)
	}
	err2 := mr.ConvertCompleted(direction, activity.Completed)
	if err2 != nil {
		return fmt.Errorf("failed to convert completed during activity conversion: %w", err2)
	}
	for _, equipment := range activity.EquipmentList {
		err3 := mr.ConvertActivityEquipment(direction, equipment)
		if err3 != nil {
			return fmt.Errorf("failed to convert activity equipment during activity conversion: %w", err3)
		}
	}
	return nil
}

// Convert a planned object
func (mr *IMeasurementRepository) ConvertPlanned(direction direction, planned *model.Planned) error {
	if direction == Incoming {
		err1 := measurement.IntoDatabaseConversion(planned.Distance)
		if err1 != nil {
			return fmt.Errorf("failed to convert incoming planned distance: %w", err1)
		}
		err2 := measurement.IntoDatabaseConversion(planned.Vertical)
		if err2 != nil {
			return fmt.Errorf("failed to convert incoming planned vertical: %w", err2)
		}
	} else if direction == Outgoing {
		err1 := measurement.LeavingDatabaseConversion(planned.Distance)
		if err1 != nil {
			return fmt.Errorf("failed to convert outgoing planned distance: %w", err1)
		}
		err2 := measurement.LeavingDatabaseConversion(planned.Vertical)
		if err2 != nil {
			return fmt.Errorf("failed to convert outgoing planned vertical: %w", err2)
		}
	}
	return nil
}

// Convert a completed object
func (mr *IMeasurementRepository) ConvertCompleted(direction direction, completed *model.Completed) error {
	if direction == Incoming {
		err1 := measurement.IntoDatabaseConversion(completed.Distance)
		if err1 != nil {
			return fmt.Errorf("failed to convert completed: %w", err1)
		}
		err2 := measurement.IntoDatabaseConversion(completed.Vertical)
		if err2 != nil {
			return fmt.Errorf("failed to convert completed: %w", err2)
		}
	} else if direction == Outgoing {
		err1 := measurement.LeavingDatabaseConversion(completed.Distance)
		if err1 != nil {
			return fmt.Errorf("failed to convert completed: %w", err1)
		}
		err2 := measurement.LeavingDatabaseConversion(completed.Vertical)
		if err2 != nil {
			return fmt.Errorf("failed to convert completed: %w", err2)
		}
	}
	return nil
}

// Convert an equipment object
func (mr *IMeasurementRepository) ConvertEquipment(direction direction, equipment *model.Equipment) error {
	if direction == Incoming {
		err1 := measurement.IntoDatabaseConversion(equipment.Mileage)
		if err1 != nil {
			return fmt.Errorf("failed to convert incoming equipment mileage: %w", err1)
		}
	} else if direction == Outgoing {
		err2 := measurement.LeavingDatabaseConversion(equipment.Mileage)
		if err2 != nil {
			return fmt.Errorf("failed to convert outgoing equipment mileage: %w", err2)
		}
	}
	return nil
}

// Convert an activity equipment object
func (mr *IMeasurementRepository) ConvertActivityEquipment(direction direction, activityEquipment *model.ActivityEquipment) error {
	err1 := mr.ConvertEquipment(direction, activityEquipment.Equipment)
	if err1 != nil {
		return fmt.Errorf("failed to convert equipment during activity equipment conversion: %w", err1)
	}

	if direction == Incoming {
		err2 := measurement.IntoDatabaseConversion(activityEquipment.AssignedMileage)
		if err2 != nil {
			return fmt.Errorf("failed to convert activity equipment assigned mileage: %w", err2)
		}
	} else if direction == Outgoing {
		err3 := measurement.LeavingDatabaseConversion(activityEquipment.AssignedMileage)
		if err3 != nil {
			return fmt.Errorf("failed to convert activity equipment assigned mileage: %w", err3)
		}
	}
	return nil
}

// convert a microcycle object
func (mr *IMeasurementRepository) ConvertMicrocycle(direction direction, microcycle *model.Microcycle, userUnitClass measurement.UnitClass) error {
	err := mr.ConvertMicrocycleSummary(direction, microcycle.Summary, userUnitClass)
	if err != nil {
		return fmt.Errorf("error converting microcycle summary: %w", err)
	}
	for _, actList := range *microcycle.CycleActivities {
		for _, act := range actList.ActivityList {
			err := mr.ConvertActivity(direction, act, userUnitClass)
			if err != nil {
				return fmt.Errorf("error converting activity during microcycle conversion: %w", err)
			}
		}
	}
	return nil
}

// Convert a microcycle summary object
func (mr *IMeasurementRepository) ConvertMicrocycleSummary(direction direction, microcycleSummary *model.MicrocycleSummary, userUnitClass measurement.UnitClass) error {
	err := mr.ConvertTotals(direction, microcycleSummary.Totals, userUnitClass)
	if err != nil {
		return fmt.Errorf("error converting totals for microcycle summary: %w", err)
	}
	err1 := mr.ConvertTotals(direction, microcycleSummary.PreviousTotals, userUnitClass)
	if err1 != nil {
		return fmt.Errorf("error converting previous totals for microcycle summary: %w", err)
	}
	err2 := mr.ConvertTotalsDifference(direction, microcycleSummary.TotalsDifferences, userUnitClass)
	if err2 != nil {
		return fmt.Errorf("error converting totals difference for microcycle summary: %w", err)
	}

	for _, tbyActType := range microcycleSummary.TotalsByActivityType {
		err := mr.ConvertTotalByActivityType(direction, tbyActType, userUnitClass)
		if err != nil {
			return fmt.Errorf("error converting totals by activity type for microcycle summary: %w", err)
		}
	}
	for _, tbyActType := range microcycleSummary.PreviousTotalsByActivityType {
		err := mr.ConvertTotalByActivityType(direction, tbyActType, userUnitClass)
		if err != nil {
			return fmt.Errorf("error converting previous totals by activity type for microcycle summary: %w", err)
		}
	}
	for _, tbyActTypeDiff := range microcycleSummary.TotalsByActivityTypeDifferences {
		err := mr.ConvertTotalByActivityTypeDifference(direction, tbyActTypeDiff, userUnitClass)
		if err != nil {
			return fmt.Errorf("error converting totals by activity type differences for microcycle summary: %w", err)
		}
	}

	for _, tbyActTypeDate := range microcycleSummary.TotalsByActivityTypeAndDate {
		err := mr.ConvertTotalByActivityTypeAndDate(direction, tbyActTypeDate, userUnitClass)
		if err != nil {
			return fmt.Errorf("error converting totals by activity type and date for microcycle summary: %w", err)
		}
	}
	return nil
}

// Convert a totals object
func (mr *IMeasurementRepository) ConvertTotals(direction direction, totals *model.Totals, userUnitClass measurement.UnitClass) error {
	if direction == Incoming {
		err := measurement.IntoDatabaseConversion(totals.TotalPlannedDistance)
		if err != nil {
			return fmt.Errorf("error converting total planned distance: %w", err)
		}
		err1 := measurement.IntoDatabaseConversion(totals.TotalPlannedVertical)
		if err1 != nil {
			return fmt.Errorf("error converting total planned distance: %w", err1)
		}

		err3 := measurement.IntoDatabaseConversion(totals.TotalCompletedDistance)
		if err3 != nil {
			return fmt.Errorf("error converting total completed distance: %w", err3)
		}
		err4 := measurement.IntoDatabaseConversion(totals.TotalCompletedVertical)
		if err4 != nil {
			return fmt.Errorf("error converting total completed distance: %w", err4)
		}
	} else if direction == Outgoing {
		err := measurement.LeavingDatabaseConversionUnitClass(totals.TotalPlannedDistance, userUnitClass, measurement.Distance)
		if err != nil {
			return fmt.Errorf("error converting total planned distance: %w", err)
		}
		err1 := measurement.LeavingDatabaseConversionUnitClass(totals.TotalPlannedVertical, userUnitClass, measurement.Vertical)
		if err1 != nil {
			return fmt.Errorf("error converting total planned distance: %w", err1)
		}

		err3 := measurement.LeavingDatabaseConversionUnitClass(totals.TotalCompletedDistance, userUnitClass, measurement.Distance)
		if err3 != nil {
			return fmt.Errorf("error converting total completed distance: %w", err3)
		}
		err4 := measurement.LeavingDatabaseConversionUnitClass(totals.TotalCompletedVertical, userUnitClass, measurement.Vertical)
		if err4 != nil {
			return fmt.Errorf("error converting total completed distance: %w", err4)
		}
	}
	return nil
}

// Convert a totals difference object
func (mr *IMeasurementRepository) ConvertTotalsDifference(direction direction, totalsDifference *model.TotalsDifferences, userUnitClass measurement.UnitClass) error {
	if direction == Incoming {
		err := measurement.IntoDatabaseConversion(totalsDifference.DifferencePlannedDistance)
		if err != nil {
			return fmt.Errorf("error converting total difference planned distance: %w", err)
		}
		err1 := measurement.IntoDatabaseConversion(totalsDifference.DifferencePlannedVertical)
		if err1 != nil {
			return fmt.Errorf("error converting total difference planned distance: %w", err1)
		}

		err3 := measurement.IntoDatabaseConversion(totalsDifference.DifferenceCompletedDistance)
		if err3 != nil {
			return fmt.Errorf("error converting total difference completed distance: %w", err3)
		}
		err4 := measurement.IntoDatabaseConversion(totalsDifference.DifferenceCompletedVertical)
		if err4 != nil {
			return fmt.Errorf("error converting total difference completed distance: %w", err4)
		}
	} else if direction == Outgoing {
		err := measurement.LeavingDatabaseConversionUnitClass(totalsDifference.DifferencePlannedDistance, userUnitClass, measurement.Distance)
		if err != nil {
			return fmt.Errorf("error converting total difference planned distance: %w", err)
		}
		err1 := measurement.LeavingDatabaseConversionUnitClass(totalsDifference.DifferencePlannedVertical, userUnitClass, measurement.Vertical)
		if err1 != nil {
			return fmt.Errorf("error converting total difference planned distance: %w", err1)
		}

		err3 := measurement.LeavingDatabaseConversionUnitClass(totalsDifference.DifferenceCompletedDistance, userUnitClass, measurement.Distance)
		if err3 != nil {
			return fmt.Errorf("error converting total difference completed distance: %w", err3)
		}
		err4 := measurement.LeavingDatabaseConversionUnitClass(totalsDifference.DifferenceCompletedVertical, userUnitClass, measurement.Vertical)
		if err4 != nil {
			return fmt.Errorf("error converting total difference completed distance: %w", err4)
		}
	}
	return nil
}

// Convert a total by activity type object
func (mr *IMeasurementRepository) ConvertTotalByActivityType(direction direction, totalByActivityType *model.TotalByActivityType, userUnitClass measurement.UnitClass) error {
	if direction == Incoming {
		err := measurement.IntoDatabaseConversion(totalByActivityType.TotalPlannedDistance)
		if err != nil {
			return fmt.Errorf("error converting activity type total planned distance: %w", err)
		}
		err1 := measurement.IntoDatabaseConversion(totalByActivityType.TotalPlannedVertical)
		if err1 != nil {
			return fmt.Errorf("error converting activity type total planned distance: %w", err1)
		}

		err3 := measurement.IntoDatabaseConversion(totalByActivityType.TotalCompletedDistance)
		if err3 != nil {
			return fmt.Errorf("error converting activity type total completed distance: %w", err3)
		}
		err4 := measurement.IntoDatabaseConversion(totalByActivityType.TotalCompletedVertical)
		if err4 != nil {
			return fmt.Errorf("error converting activity type total completed distance: %w", err4)
		}
	} else if direction == Outgoing {

		// calculate paces before conversion
		totalByActivityType.CalculatePaces(userUnitClass)

		err := measurement.LeavingDatabaseConversionUnitClass(totalByActivityType.TotalPlannedDistance, userUnitClass, measurement.Distance)
		if err != nil {
			return fmt.Errorf("error converting activity type total planned distance: %w", err)
		}
		totalByActivityType.TotalPlannedDistance.Unit = measurement.UnitClassControl(userUnitClass, measurement.Distance)

		err1 := measurement.LeavingDatabaseConversionUnitClass(totalByActivityType.TotalPlannedVertical, userUnitClass, measurement.Vertical)
		if err1 != nil {
			return fmt.Errorf("error converting activity type total planned distance: %w", err1)
		}
		totalByActivityType.TotalPlannedVertical.Unit = measurement.UnitClassControl(userUnitClass, measurement.Vertical)

		err3 := measurement.LeavingDatabaseConversionUnitClass(totalByActivityType.TotalCompletedDistance, userUnitClass, measurement.Distance)
		if err3 != nil {
			return fmt.Errorf("error converting activity type total completed distance: %w", err3)
		}
		totalByActivityType.TotalCompletedDistance.Unit = measurement.UnitClassControl(userUnitClass, measurement.Distance)

		err4 := measurement.LeavingDatabaseConversionUnitClass(totalByActivityType.TotalCompletedVertical, userUnitClass, measurement.Vertical)
		if err4 != nil {
			return fmt.Errorf("error converting activity type total completed distance: %w", err4)
		}
		totalByActivityType.TotalCompletedVertical.Unit = measurement.UnitClassControl(userUnitClass, measurement.Vertical)
	}
	return nil
}
func (mr *IMeasurementRepository) ConvertTotalByActivityTypeDifference(direction direction, totalByActivityTypeDifference *model.TotalByActivityTypeDifferences, userUnitClass measurement.UnitClass) error {
	if direction == Incoming {
		err := measurement.IntoDatabaseConversion(totalByActivityTypeDifference.DifferencePlannedDistance)
		if err != nil {
			return fmt.Errorf("error converting activity type difference total planned distance: %w", err)
		}
		err1 := measurement.IntoDatabaseConversion(totalByActivityTypeDifference.DifferencePlannedVertical)
		if err1 != nil {
			return fmt.Errorf("error converting activity type difference total planned distance: %w", err1)
		}

		err3 := measurement.IntoDatabaseConversion(totalByActivityTypeDifference.DifferenceCompletedDistance)
		if err3 != nil {
			return fmt.Errorf("error converting activity type difference total completed distance: %w", err3)
		}
		err4 := measurement.IntoDatabaseConversion(totalByActivityTypeDifference.DifferenceCompletedVertical)
		if err4 != nil {
			return fmt.Errorf("error converting activity type difference total completed distance: %w", err4)
		}
	} else if direction == Outgoing {
		err := measurement.LeavingDatabaseConversionUnitClass(totalByActivityTypeDifference.DifferencePlannedDistance, userUnitClass, measurement.Distance)
		if err != nil {
			return fmt.Errorf("error converting activity type difference total planned distance: %w", err)
		}
		totalByActivityTypeDifference.DifferencePlannedDistance.Unit = measurement.UnitClassControl(userUnitClass, measurement.Distance)

		err1 := measurement.LeavingDatabaseConversionUnitClass(totalByActivityTypeDifference.DifferencePlannedVertical, userUnitClass, measurement.Vertical)
		if err1 != nil {
			return fmt.Errorf("error converting activity type difference total planned distance: %w", err1)
		}
		totalByActivityTypeDifference.DifferencePlannedVertical.Unit = measurement.UnitClassControl(userUnitClass, measurement.Vertical)

		err3 := measurement.LeavingDatabaseConversionUnitClass(totalByActivityTypeDifference.DifferenceCompletedDistance, userUnitClass, measurement.Distance)
		if err3 != nil {
			return fmt.Errorf("error converting activity type difference total completed distance: %w", err3)
		}
		totalByActivityTypeDifference.DifferenceCompletedDistance.Unit = measurement.UnitClassControl(userUnitClass, measurement.Distance)

		err4 := measurement.LeavingDatabaseConversionUnitClass(totalByActivityTypeDifference.DifferenceCompletedVertical, userUnitClass, measurement.Vertical)
		if err4 != nil {
			return fmt.Errorf("error converting activity type difference total completed distance: %w", err4)
		}
		totalByActivityTypeDifference.DifferenceCompletedVertical.Unit = measurement.UnitClassControl(userUnitClass, measurement.Vertical)
	}
	return nil
}

// Convert a total by activity type and date object
func (mr *IMeasurementRepository) ConvertTotalByActivityTypeAndDate(direction direction, totalByActivityTypeAndDate *model.TotalByActivityTypeAndDate, userUnitClass measurement.UnitClass) error {
	if direction == Incoming {
		err := measurement.IntoDatabaseConversion(totalByActivityTypeAndDate.TotalPlannedDistance)
		if err != nil {
			return fmt.Errorf("error converting activity type and date total planned distance: %w", err)
		}
		err1 := measurement.IntoDatabaseConversion(totalByActivityTypeAndDate.TotalPlannedVertical)
		if err1 != nil {
			return fmt.Errorf("error converting activity type and date total planned distance: %w", err1)
		}

		err3 := measurement.IntoDatabaseConversion(totalByActivityTypeAndDate.TotalCompletedDistance)
		if err3 != nil {
			return fmt.Errorf("error converting activity type and date total completed distance: %w", err3)
		}
		err4 := measurement.IntoDatabaseConversion(totalByActivityTypeAndDate.TotalCompletedVertical)
		if err4 != nil {
			return fmt.Errorf("error converting activity type and date total completed distance: %w", err4)
		}
	} else if direction == Outgoing {
		err := measurement.LeavingDatabaseConversionUnitClass(totalByActivityTypeAndDate.TotalPlannedDistance, userUnitClass, measurement.Distance)
		if err != nil {
			return fmt.Errorf("error converting activity type activity type and date total planned distance: %w", err)
		}
		totalByActivityTypeAndDate.TotalPlannedDistance.Unit = measurement.UnitClassControl(userUnitClass, measurement.Distance)

		err1 := measurement.LeavingDatabaseConversionUnitClass(totalByActivityTypeAndDate.TotalPlannedVertical, userUnitClass, measurement.Vertical)
		if err1 != nil {
			return fmt.Errorf("error converting activity type activity type and date total planned distance: %w", err1)
		}
		totalByActivityTypeAndDate.TotalPlannedVertical.Unit = measurement.UnitClassControl(userUnitClass, measurement.Vertical)

		err3 := measurement.LeavingDatabaseConversionUnitClass(totalByActivityTypeAndDate.TotalCompletedDistance, userUnitClass, measurement.Distance)
		if err3 != nil {
			return fmt.Errorf("error converting activity type activity type and date total completed distance: %w", err3)
		}
		totalByActivityTypeAndDate.TotalCompletedDistance.Unit = measurement.UnitClassControl(userUnitClass, measurement.Distance)

		err4 := measurement.LeavingDatabaseConversionUnitClass(totalByActivityTypeAndDate.TotalCompletedVertical, userUnitClass, measurement.Vertical)
		if err4 != nil {
			return fmt.Errorf("error converting activity type activity type and date total completed distance: %w", err4)
		}
		totalByActivityTypeAndDate.TotalCompletedVertical.Unit = measurement.UnitClassControl(userUnitClass, measurement.Vertical)
	}
	return nil
}

// Convert an n cycle summary object
func (mr *IMeasurementRepository) ConvertNCycleSummary(direction direction, nCycleSummary *model.NCycleSummary, userUnitClass measurement.UnitClass) error {
	if direction == Incoming {
		for _, pd := range nCycleSummary.PlannedDistances {
			err := measurement.IntoDatabaseConversion(pd)
			if err != nil {
				return fmt.Errorf("error converting planned distances in n cycle summary: %w", err)
			}
		}
		for _, pv := range nCycleSummary.PlannedVerticals {
			err := measurement.IntoDatabaseConversion(pv)
			if err != nil {
				return fmt.Errorf("error converting planned verticals in n cycle summary: %w", err)
			}
		}
		for _, cd := range nCycleSummary.CompletedDistances {
			err := measurement.IntoDatabaseConversion(cd)
			if err != nil {
				return fmt.Errorf("error converting completed distances in n cycle summary: %w", err)
			}
		}
		for _, cv := range nCycleSummary.CompletedVerticals {
			err := measurement.IntoDatabaseConversion(cv)
			if err != nil {
				return fmt.Errorf("error converting completed verticals in n cycle summary: %w", err)
			}
		}
	} else if direction == Outgoing {
		for _, pd := range nCycleSummary.PlannedDistances {
			err := measurement.LeavingDatabaseConversionUnitClass(pd, userUnitClass, measurement.Distance)
			pd.Unit = measurement.UnitClassControl(userUnitClass, measurement.Distance)
			if err != nil {
				return fmt.Errorf("error converting planned distances in n cycle summary: %w", err)
			}
		}
		for _, pv := range nCycleSummary.PlannedVerticals {
			err := measurement.LeavingDatabaseConversionUnitClass(pv, userUnitClass, measurement.Vertical)
			pv.Unit = measurement.UnitClassControl(userUnitClass, measurement.Vertical)
			if err != nil {
				return fmt.Errorf("error converting planned verticals in n cycle summary: %w", err)
			}
		}
		for _, cd := range nCycleSummary.CompletedDistances {
			err := measurement.LeavingDatabaseConversionUnitClass(cd, userUnitClass, measurement.Distance)
			cd.Unit = measurement.UnitClassControl(userUnitClass, measurement.Distance)
			if err != nil {
				return fmt.Errorf("error converting completed distances in n cycle summary: %w", err)
			}
		}
		for _, cv := range nCycleSummary.CompletedVerticals {
			err := measurement.LeavingDatabaseConversionUnitClass(cv, userUnitClass, measurement.Vertical)
			cv.Unit = measurement.UnitClassControl(userUnitClass, measurement.Vertical)
			if err != nil {
				return fmt.Errorf("error converting completed verticals in n cycle summary: %w", err)
			}
		}
	}
	return nil
}
