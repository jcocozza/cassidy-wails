package microcyclerepo

import (
	"fmt"
	"time"

	"github.com/jcocozza/cassidy-wails/internal/database"
	"github.com/jcocozza/cassidy-wails/internal/model"
	"github.com/jcocozza/cassidy-wails/internal/sqlcode"
	"github.com/jcocozza/cassidy-wails/internal/utils/dateutil"
	"github.com/jcocozza/cassidy-wails/internal/utils/measurement"
)

// The microcycle repository contains functions for queries related to microcycles
type MicrocycleRepository interface {
	ReadMicrocycle(startDate, endDate time.Time, userUuid string, userUnitClass measurement.UnitClass) (*model.Microcycle, error)

	ReadActivityEquipmentList(activityUuidList []string) ([]*model.ActivityEquipment, error)
	ReadActivityTypeSubtypeList(activityUuidList []string) ([]*model.ActivityTypeSubtype, error)
	ReadCycle(startDate, endDate time.Time, userUuid string) (*model.Cycle, error)

	//ReadTotals(startDate string, endDate string, userUuid string, userUnitClass measurement.UnitClass) (*model.Totals, error)
	ReadTotalsPreviousCurrent(startDate, endDate time.Time, userUuid string, userUnitClass measurement.UnitClass) (*model.Totals, *model.Totals, error)
	//ReadTotalsByActivityType(startDate string, endDate string, userUuid string, userUnitClass measurement.UnitClass) ([]*model.TotalByActivityType, error)
	ReadTotalsByActivityTypePreviousCurrent(startDate, endDate time.Time, userUuid string, userUnitClass measurement.UnitClass) ([]*model.TotalByActivityType, []*model.TotalByActivityType, error)
	ReadTotalsByActivityTypeAndDate(startDate, endDate time.Time, userUuid string, userUnitClass measurement.UnitClass) ([]*model.TotalByActivityTypeAndDate, error)
}

// Represents a SQLite database connection
type IMicrocycleRepository struct {
	DB database.DbOperations
}

func NewIMicrocycleRepository(db database.DbOperations) *IMicrocycleRepository {
	return &IMicrocycleRepository{
		DB: db,
	}
}
// Get microcycle information from the database
func (db *IMicrocycleRepository) ReadMicrocycle(startDate, endDate time.Time, userUuid string, userUnitClass measurement.UnitClass) (*model.Microcycle, error) {
	c, err := db.ReadCycle(startDate, endDate, userUuid)
	if err != nil {
		return nil, fmt.Errorf("error during cycle read: %w", err)
	}
	//totals, err1 := db.ReadTotals(startDate, endDate, userUuid, userUnitClass)
	//if err1 != nil {
	//	return nil, fmt.Errorf("error during totals read: %w", err1)
	//}
	totalsPrevious, totalsCurrent, err1 := db.ReadTotalsPreviousCurrent(startDate, endDate, userUuid, userUnitClass)
	if err1 != nil {
		return nil, fmt.Errorf("error during previous/current totals read: %w", err1)
	}

	const numPriors = 4
	s, e := dateutil.GeneratePriorsRange(startDate, endDate, numPriors)
	averagePreviousTotals, err1 := db.ReadAveragePriorTotals(s, e, numPriors, userUuid, userUnitClass)
	if err1 != nil {
		return nil, fmt.Errorf("error during prior totals read: %w", err1)
	}
	//totalsByActivityType, err2 := db.ReadTotalsByActivityType(startDate, endDate, userUuid, userUnitClass)
	//if err2 != nil {
	//	return nil, fmt.Errorf("error during totals by activity type read: %w", err2)
	//}
	totalsByActivityTypePrevious, totalsByActivityTypeCurrent, err2 := db.ReadTotalsByActivityTypePreviousCurrent(startDate, endDate, userUuid, userUnitClass)
	if err2 != nil {
		return nil, fmt.Errorf("error during previous/current totals by activity type read: %w", err2)
	}
	totalsByActivityTypeAndDate, err3 := db.ReadTotalsByActivityTypeAndDate(startDate, endDate, userUuid, userUnitClass)
	if err3 != nil {
		return nil, fmt.Errorf("error during totals by activity type and date read: %w", err3)
	}

	summary := &model.MicrocycleSummary{
		Totals:                       totalsCurrent,
		PreviousTotals:               totalsPrevious,
		AveragePreviousTotals:        averagePreviousTotals,
		TotalsByActivityType:         totalsByActivityTypeCurrent,
		PreviousTotalsByActivityType: totalsByActivityTypePrevious,
		TotalsByActivityTypeAndDate:  totalsByActivityTypeAndDate,
	}
	err4 := summary.CalculateCycleChanges()
	if err4 != nil {
		return nil, fmt.Errorf("error while calculating cycle changes: %w", err4)
	}

	mc := &model.Microcycle{
		StartDate:       startDate,
		EndDate:         endDate,
		CycleActivities: c,
		Summary:         summary,
	}
	return mc, nil
}

// Get all activity equipment for a given list of activity uuids
func (db *IMicrocycleRepository) ReadActivityEquipmentList(activityUuidList []string) ([]*model.ActivityEquipment, error) {
	sql := sqlcode.SQLReader(sqlcode.ActivityEquipment_list)

	query := database.SQLListInsertion(sql, len(activityUuidList))
	args := make([]interface{}, len(activityUuidList))
	for i, v := range activityUuidList {
		args[i] = v
	}

	rows, err := db.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	activityEquipmentList := []*model.ActivityEquipment{}
	for rows.Next() {
		tmpAE := model.EmptyActivityEquipment()
		err := rows.Scan(
			&tmpAE.Id, &tmpAE.ActivityUuid, &tmpAE.AssignedMileage.Length, &tmpAE.AssignedMileage.Unit,
			&tmpAE.Equipment.Id, &tmpAE.Equipment.UserUuid, &tmpAE.Equipment.Name, &tmpAE.Equipment.Brand, &tmpAE.Equipment.Model, &tmpAE.Equipment.Cost, &tmpAE.Equipment.Size, &tmpAE.Equipment.PurchaseDate, &tmpAE.Equipment.Notes, &tmpAE.Equipment.Mileage.Length, &tmpAE.Equipment.Mileage.Unit,
			&tmpAE.Equipment.EquipmentType.Id, &tmpAE.Equipment.EquipmentType.Name,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		} else {
			err2 := tmpAE.Validate(false)
			if err2 != nil {
				return nil, fmt.Errorf("activity equipment failed to validate: %w", err2)
			} else {
				activityEquipmentList = append(activityEquipmentList, tmpAE)
			}
		}
	}
	return activityEquipmentList, nil
}
// Get all activity types for the activity uuid list
func (db *IMicrocycleRepository) ReadActivityTypeSubtypeList(activityUuidList []string) ([]*model.ActivityTypeSubtype, error) {
	sql := sqlcode.SQLReader(sqlcode.ActivityTypeSubtype_list)

	query := database.SQLListInsertion(sql, len(activityUuidList))
	args := make([]interface{}, len(activityUuidList))
	for i, v := range activityUuidList {
		args[i] = v
	}

	rows, err := db.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	activityTypeSubtypeList := []*model.ActivityTypeSubtype{}
	for rows.Next() {
		tmpTS := model.EmptyActivityTypeSubtype()
		err := rows.Scan(&tmpTS.Id, &tmpTS.ActivityUuid,
			&tmpTS.ActivityType.Id, &tmpTS.ActivityType.Name,
			&tmpTS.ActivitySubtype.Id, &tmpTS.ActivitySubtype.SuperTypeId, &tmpTS.ActivitySubtype.Name,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		} else {
			err2 := tmpTS.Validate()
			if err2 != nil {
				return nil, fmt.Errorf("activity type subtype failed to validate: %w", err2)
			} else {
				activityTypeSubtypeList = append(activityTypeSubtypeList, tmpTS)
			}
		}
	}
	return activityTypeSubtypeList, nil
}
// Get a list of activity lists
func (db *IMicrocycleRepository) ReadCycle(startDate, endDate time.Time, userUuid string) (*model.Cycle, error) {
	sql := sqlcode.SQLReader(sqlcode.Microcycle_read_activity_list)
	rows, err := db.DB.Query(sql, startDate.Format(dateutil.Layout), endDate.Format(dateutil.Layout), userUuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cycle := model.NewCycle(startDate, endDate)
	for rows.Next() {
		tmpAct := model.EmptyActivity()
		tmpDateStr := ""
		err := rows.Scan(&tmpAct.Uuid,
			&tmpDateStr, &tmpAct.Order, &tmpAct.Name, &tmpAct.Description, &tmpAct.Notes, &tmpAct.IsRace, &tmpAct.NumStrides, &tmpAct.Map,
			&tmpAct.Type.Id, &tmpAct.Type.Name,
			&tmpAct.Planned.Distance.Length, &tmpAct.Planned.Distance.Unit, &tmpAct.Planned.Duration, &tmpAct.Planned.Vertical.Length, &tmpAct.Planned.Vertical.Unit,
			&tmpAct.Completed.Distance.Length, &tmpAct.Completed.Distance.Unit, &tmpAct.Completed.Duration, &tmpAct.Completed.Vertical.Length, &tmpAct.Completed.Vertical.Unit,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		} else {

			tmpDate, err := time.Parse(dateutil.TimeLayout, tmpDateStr)
			if err != nil {
				return nil, fmt.Errorf("activity date failed to parse: %w", err)
			}
			tmpAct.Date = tmpDate
			tmpAct.SetUuid(tmpAct.Uuid)
			err2 := tmpAct.Validate()
			if err2 != nil {
				return nil, fmt.Errorf("activity failed to validate: %w", err2)
			} else {
				err4 := cycle.AddActivity(tmpAct)
				if err4 != nil {
					return nil, fmt.Errorf("failed to add activity to cycle: %w", err4)
				}
			}
		}
	}
	//fmt.Println(gostructstringify.StructStringify(cycle))
	uuidList := cycle.CreateUuidList()
	aeList, err1 := db.ReadActivityEquipmentList(uuidList)
	if err1 != nil || aeList == nil {
		return nil, fmt.Errorf("failed to read activity equipment list: %w", err1)
	}

	atsList, err2 := db.ReadActivityTypeSubtypeList(uuidList)
	if err2 != nil || atsList == nil {
		return nil, fmt.Errorf("failed to read activity type subtype list: %w", err2)
	}

	for _, ae := range aeList {
		for _, al := range *cycle {
			for _, act := range al.ActivityList {
				if ae.ActivityUuid == act.Uuid {
					act.AddActivityEquipment(ae)
				}
			}
		}
	}
	for _, ats := range atsList {
		for _, al := range *cycle {
			for _, act := range al.ActivityList {
				if ats.ActivityUuid == act.Uuid {
					act.AddActivityTypeSubtype(ats)
				}
			}
		}
	}

	return cycle, nil
}
// Get totals for the current date date and the previous date range
func (db *IMicrocycleRepository) ReadTotalsPreviousCurrent(startDate, endDate time.Time, userUuid string, userUnitClass measurement.UnitClass) (*model.Totals, *model.Totals, error) {

	previousStart, previousEnd := dateutil.GetPreviousCycle(startDate, endDate)

	sql := sqlcode.SQLReader(sqlcode.Microcycle_read_totals_current_previous)
	rows, err := db.DB.Query(sql, previousStart.Format(dateutil.Layout), previousEnd.Format(dateutil.Layout), startDate.Format(dateutil.Layout), endDate.Format(dateutil.Layout), previousStart, endDate.Format(dateutil.Layout), userUuid)
	if err != nil {
		return nil, nil, fmt.Errorf("previous/current totals did not query properly: %w", err)
	}
	defer rows.Close()

	previousTotals := model.EmptyTotals(userUnitClass)
	currentTotals := model.EmptyTotals(userUnitClass)

	for rows.Next() {
		cycleLocation := ""
		totals := model.EmptyTotals(userUnitClass)
		err := rows.Scan(&cycleLocation,
			&totals.TotalPlannedDistance.Length, &totals.TotalPlannedDuration, &totals.TotalPlannedVertical.Length,
			&totals.TotalCompletedDistance.Length, &totals.TotalCompletedDuration, &totals.TotalCompletedVertical.Length)
		if err != nil {
			return nil, nil, fmt.Errorf("previous/current totals did not scan properly: %w", err)
		}
		// there will only be two returned rows, one with previous, the other with current
		if cycleLocation == "previous" {
			previousTotals = totals
		} else if cycleLocation == "current" {
			currentTotals = totals
		}
	}
	return previousTotals, currentTotals, nil
}
// Get the average totals over a given date range where the divisor is the number of cycles in the range
func (db *IMicrocycleRepository) ReadAveragePriorTotals(priorStart, priorEnd time.Time, numPriors int, userUuid string, userUnitClass measurement.UnitClass) (*model.Totals, error) {
	sql := sqlcode.SQLReader(sqlcode.Microcycle_read_totals_date_range)

	row := db.DB.QueryRow(sql, numPriors, numPriors, numPriors, numPriors, numPriors, numPriors, priorStart.Format(dateutil.Layout), priorEnd.Format(dateutil.Layout), userUuid)
	totals := model.EmptyTotals(userUnitClass)
	err := row.Scan(&totals.TotalPlannedDistance.Length, &totals.TotalPlannedDuration, &totals.TotalPlannedVertical.Length,
		&totals.TotalCompletedDistance.Length, &totals.TotalCompletedDuration, &totals.TotalCompletedVertical.Length)
	if err != nil {
		return nil, fmt.Errorf("failed to read average prior totals: %w", err)
	}

	return totals, nil
}
// Get totals by activity type for the current date date and the previous date range
//
// Will return an empty list if errors
func (db *IMicrocycleRepository) ReadTotalsByActivityTypePreviousCurrent(startDate, endDate time.Time, userUuid string, userUnitClass measurement.UnitClass) ([]*model.TotalByActivityType, []*model.TotalByActivityType, error) {
	previousStart, previousEnd := dateutil.GetPreviousCycle(startDate, endDate)

	sql := sqlcode.SQLReader(sqlcode.Microcycle_read_totals_by_activity_type_current_previous)

	rows, err := db.DB.Query(sql,
		previousStart.Format(dateutil.Layout), previousEnd.Format(dateutil.Layout),
		previousStart.Format(dateutil.Layout), previousEnd.Format(dateutil.Layout),
		previousStart.Format(dateutil.Layout), previousEnd.Format(dateutil.Layout),
		previousStart.Format(dateutil.Layout), previousEnd.Format(dateutil.Layout),
		previousStart.Format(dateutil.Layout), previousEnd.Format(dateutil.Layout),
		previousStart.Format(dateutil.Layout), previousEnd.Format(dateutil.Layout),
		startDate.Format(dateutil.Layout), endDate.Format(dateutil.Layout),
		startDate.Format(dateutil.Layout), endDate.Format(dateutil.Layout),
		startDate.Format(dateutil.Layout), endDate.Format(dateutil.Layout),
		startDate.Format(dateutil.Layout), endDate.Format(dateutil.Layout),
		startDate.Format(dateutil.Layout), endDate.Format(dateutil.Layout),
		startDate.Format(dateutil.Layout), endDate.Format(dateutil.Layout),
		previousStart.Format(dateutil.Layout), endDate.Format(dateutil.Layout), userUuid)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to query previous/current totals by activity type: %w", err)
	}
	defer rows.Close()

	allTotalsPrevious := []*model.TotalByActivityType{}
	allTotalsCurrent := []*model.TotalByActivityType{}

	var distUnit measurement.Unit = measurement.Meter
	var vertUnit measurement.Unit = measurement.Meter

	for rows.Next() {

		actType := model.EmptyActivityType()

		tmpTotActTypePrevious := model.EmptyTotalByActivityType(userUnitClass)
		tmpTotActTypeCurrent := model.EmptyTotalByActivityType(userUnitClass)
		err1 := rows.Scan(&actType.Id, &actType.Name,
			&tmpTotActTypePrevious.TotalPlannedDistance.Length, &tmpTotActTypePrevious.TotalPlannedDuration, &tmpTotActTypePrevious.TotalPlannedVertical.Length,
			&tmpTotActTypePrevious.TotalCompletedDistance.Length, &tmpTotActTypePrevious.TotalCompletedDuration, &tmpTotActTypePrevious.TotalCompletedVertical.Length,

			&tmpTotActTypeCurrent.TotalPlannedDistance.Length, &tmpTotActTypeCurrent.TotalPlannedDuration, &tmpTotActTypeCurrent.TotalPlannedVertical.Length,
			&tmpTotActTypeCurrent.TotalCompletedDistance.Length, &tmpTotActTypeCurrent.TotalCompletedDuration, &tmpTotActTypeCurrent.TotalCompletedVertical.Length)

		tmpTotActTypePrevious.ActivityType = actType
		tmpTotActTypeCurrent.ActivityType = actType

		tmpTotActTypePrevious.TotalPlannedDistance.Unit = distUnit
		tmpTotActTypePrevious.TotalPlannedVertical.Unit = vertUnit
		tmpTotActTypePrevious.TotalCompletedDistance.Unit = distUnit
		tmpTotActTypePrevious.TotalCompletedVertical.Unit = vertUnit

		tmpTotActTypeCurrent.TotalPlannedDistance.Unit = distUnit
		tmpTotActTypeCurrent.TotalPlannedVertical.Unit = vertUnit
		tmpTotActTypeCurrent.TotalCompletedDistance.Unit = distUnit
		tmpTotActTypeCurrent.TotalCompletedVertical.Unit = vertUnit
		if err1 != nil {
			return nil, nil, fmt.Errorf("previous/current totals by activity type did not scan properly: %w", err1)
		}
		allTotalsPrevious = append(allTotalsPrevious, tmpTotActTypePrevious)
		allTotalsCurrent = append(allTotalsCurrent, tmpTotActTypeCurrent)
	}
	return allTotalsPrevious, allTotalsCurrent, nil
}
// Get totals by activity type and date for a date range
func (db *IMicrocycleRepository) ReadTotalsByActivityTypeAndDate(startDate, endDate time.Time, userUuid string, userUnitClass measurement.UnitClass) ([]*model.TotalByActivityTypeAndDate, error) {
	sql := sqlcode.SQLReader(sqlcode.Microcycle_read_totals_by_activity_type_and_date)
	rows, err := db.DB.Query(sql, startDate.Format(dateutil.Layout), endDate.Format(dateutil.Layout), userUuid)
	if err != nil {
		return nil, fmt.Errorf("failed to query totals by activity type and date: %w", err)
	}
	defer rows.Close()

	allTotals := []*model.TotalByActivityTypeAndDate{}
	for rows.Next() {
		totActTypeDate := model.EmptyTotalByActivityTypeAndDate(userUnitClass)
		dateStr := ""
		err1 := rows.Scan(&totActTypeDate.ActivityType.Id, &totActTypeDate.ActivityType.Name, &dateStr,
			&totActTypeDate.TotalPlannedDistance.Length, &totActTypeDate.TotalPlannedDuration, &totActTypeDate.TotalPlannedVertical.Length, &totActTypeDate.TotalPlannedDistance.Unit, &totActTypeDate.TotalPlannedVertical.Unit,
			&totActTypeDate.TotalCompletedDistance.Length, &totActTypeDate.TotalCompletedDuration, &totActTypeDate.TotalCompletedVertical.Length, &totActTypeDate.TotalCompletedDistance.Unit, &totActTypeDate.TotalCompletedVertical.Unit,
		)
		if err1 != nil {
			return nil, fmt.Errorf("error scanning row: %w", err1)
		}

		date, err := time.Parse(dateutil.TimeLayout, dateStr)
		if err != nil {
			return nil, fmt.Errorf("date failed to parse when reading totals by type and date: %w", err)
		}
		totActTypeDate.Date = date

		allTotals = append(allTotals, totActTypeDate)
	}
	return allTotals, nil
}
