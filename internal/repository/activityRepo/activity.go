package activityrepo

import (
	"fmt"
	"time"

	"github.com/jcocozza/cassidy-wails/internal/database"
	"github.com/jcocozza/cassidy-wails/internal/model"
	"github.com/jcocozza/cassidy-wails/internal/sqlcode"
	"github.com/jcocozza/cassidy-wails/internal/utils/dateutil"
)

// Functions for working with activity objects.
type ActivityRepository interface {
	// Activity

	Create(userUuid string, activity *model.Activity) error
	Read(activityUuid string) (*model.Activity, error)
	Update(completed *model.Activity) error
	Delete(activityUuid string) error
    GetMostRecentDate(userUuid string) (time.Time, error)
	CreateOrMerge(activity *model.Activity, userUuid string) error
}
// Represents a database connection.
type IActivityRepository struct {
	DB database.DbOperations
}
func NewIActivityRepository(db database.DbOperations) *IActivityRepository {
	return &IActivityRepository{
		DB: db,
	}
}
// Create an activity in the database.
//
// The creates a new row in the activity table, planned table, completed table, and optionally in the activity equipment and activity type subtype tables.
func (db *IActivityRepository) Create(userUuid string, activity *model.Activity) error {
	sqlActivity := sqlcode.SQLReader(sqlcode.Activity_create)
	err := db.DB.Execute(sqlActivity, activity.Uuid, userUuid, activity.Date.Format(dateutil.TimeLayout), activity.Order, activity.Name, activity.Description, activity.Notes, activity.Type.Id, activity.IsRace, activity.NumStrides, activity.Map)
	if err != nil {
		return fmt.Errorf("failed to insert into activity table: %w", err)
	}
	// create planned for the new activity
	sqlPlanned := sqlcode.SQLReader(sqlcode.Planned_create)
	err1 := db.DB.Execute(sqlPlanned, activity.Planned.ActivityUuid,
		activity.Planned.Distance.Length, activity.Planned.Distance.Unit,
		activity.Planned.Duration,
		activity.Planned.Vertical.Length, activity.Planned.Vertical.Unit)
	if err1 != nil {
		return fmt.Errorf("failed to insert into planned table: %w", err)
	}
	// create completed for the new activity
	sqlCompleted := sqlcode.SQLReader(sqlcode.Completed_create)
	err2 := db.DB.Execute(sqlCompleted, activity.Completed.ActivityUuid,
		activity.Completed.Distance.Length, activity.Completed.Distance.Unit,
		activity.Completed.Duration,
		activity.Completed.Vertical.Length, activity.Completed.Vertical.Unit)
	if err2 != nil {
		return fmt.Errorf("failed to insert into completed table: %w", err)
	}
	// add activity equipment (if any)
	sqlActivityEquipment := sqlcode.SQLReader(sqlcode.ActivityEquipment_create)
	for _, ae := range activity.EquipmentList {
		err3 := db.DB.Execute(sqlActivityEquipment, ae.ActivityUuid, ae.Equipment.Id, ae.AssignedMileage.Length, ae.AssignedMileage.Unit)
		if err3 != nil {
			return fmt.Errorf("failed to insert into activity equipment table: %w", err)
		}
	}
	// add activity type subtype(s) (if any)
	sqlActivityTypeSubtype := sqlcode.SQLReader(sqlcode.ActivityTypeSubtype_create)
	for _, ats := range activity.TypeSubtypeList {
		id, err4 := db.DB.ExecuteGetLast(sqlActivityTypeSubtype, ats.ActivityUuid, ats.ActivityType.Id, ats.ActivitySubtype.Id)
		if err4 != nil {
			return fmt.Errorf("failed to insert into activity type subtype table: %w", err)
		}
		ats.SetId(id)
	}
	return nil
}
// Read an activity
func (db *IActivityRepository) Read(activityUuid string) (*model.Activity, error) {
	sql := sqlcode.SQLReader(sqlcode.Activity_read)
	row := db.DB.QueryRow(sql, activityUuid)

	activity := model.EmptyActivity()
	dateStr := ""
	err := row.Scan(&activity.Uuid,
		&dateStr, &activity.Order, &activity.Name, &activity.Description, &activity.Notes, &activity.IsRace, &activity.NumStrides, &activity.Map,
		&activity.Type.Id, &activity.Type.Name,
		&activity.Planned.Distance.Length, &activity.Planned.Distance.Unit, &activity.Planned.Duration, &activity.Planned.Vertical.Length, &activity.Planned.Vertical.Unit,
		&activity.Completed.Distance.Length, &activity.Completed.Distance.Unit, &activity.Completed.Duration, &activity.Completed.Vertical.Length, &activity.Completed.Vertical.Unit,
	)
	if err != nil {
		return nil, fmt.Errorf("error scanning row: %w", err)
	} else {
		tmpDate, err := time.Parse(dateutil.TimeLayout, dateStr)
		if err != nil {
			return nil, fmt.Errorf("activity date failed to parse: %w", err)
		}
		activity.Date = tmpDate
		activity.SetUuid(activity.Uuid)
		err2 := activity.Validate()
		if err2 != nil {
			return nil, fmt.Errorf("activity failed to validate: %w", err2)
		}
	}

	return activity, nil
}

// Update an activity
//
// Note: a change in the activity_type_id will trigger a delete of all activity type subtypes
//
// Updates the activity and its planned/completed
func (db *IActivityRepository) Update(activity *model.Activity) error {
	sqlActivity := sqlcode.SQLReader(sqlcode.Activity_update)
	err := db.DB.Execute(sqlActivity, activity.Date.Format(dateutil.TimeLayout), activity.Order, activity.Name, activity.Description, activity.Notes, activity.Type.Id, activity.IsRace, activity.NumStrides, activity.Uuid)

	if err != nil {
		return fmt.Errorf("error updating activity in database: %w", err)
	}

	sqlPlanned := sqlcode.SQLReader(sqlcode.Planned_update)
	err1 := db.DB.Execute(sqlPlanned, activity.Planned.Distance.Length, activity.Planned.Distance.Unit, activity.Planned.Duration, activity.Planned.Vertical.Length, activity.Planned.Vertical.Unit, activity.Planned.ActivityUuid)

	if err1 != nil {
		return fmt.Errorf("error updating planned in database: %w", err1)
	}

	sqlCompleted := sqlcode.SQLReader(sqlcode.Completed_update)
	err2 := db.DB.Execute(sqlCompleted, activity.Completed.Distance.Length, activity.Completed.Distance.Unit, activity.Completed.Duration, activity.Completed.Vertical.Length, activity.Completed.Vertical.Unit, activity.Completed.ActivityUuid)

	if err2 != nil {
		return fmt.Errorf("error updating completed in database: %w", err2)
	}

	// TO update activity type subtypes, we delete them all and then recreate them.
	sqlDeleteActivityTypes := sqlcode.SQLReader(sqlcode.ActivityTypeSubtype_delete_all_by_activity_uuid)
	err3 := db.DB.Execute(sqlDeleteActivityTypes, activity.Uuid)
	if err3 != nil {
		return fmt.Errorf("error deleting type subtype for update of activity: %w", err3)
	}

	// add activity type subtype(s) (if any)
	sqlActivityTypeSubtype := sqlcode.SQLReader(sqlcode.ActivityTypeSubtype_create)
	for _, ats := range activity.TypeSubtypeList {
		id, err4 := db.DB.ExecuteGetLast(sqlActivityTypeSubtype, ats.ActivityUuid, ats.ActivityType.Id, ats.ActivitySubtype.Id)
		if err4 != nil {
			return fmt.Errorf("failed to insert into activity type subtype table: %w", err)
		}
		ats.SetId(id)
	}
	return nil
}
// Delete an activity from the database
//
// Note: A delete will trigger a procedure in the database that deletes other things
// so we only need to worry about deleting the activity in the activity table.
func (db *IActivityRepository) Delete(activityUuid string) error {
	sql := sqlcode.SQLReader(sqlcode.Activity_delete)
	err := db.DB.Execute(sql, activityUuid)
	if err != nil {
		return fmt.Errorf("error deleting activity: %w", err)
	}
	return nil
}
// Get the most recent activity that has completed data
func (db *IActivityRepository) GetMostRecentDate(userUuid string) (time.Time, error) {
	sql := sqlcode.SQLReader(sqlcode.Activity_GetMostRecent)
	row := db.DB.QueryRow(sql, userUuid)

	var mostRecentDateStr string
	err := row.Scan(&mostRecentDateStr)
	if err != nil {
		return time.Time{}, err
	}

	mostRecentDate, err := time.Parse(dateutil.Layout, mostRecentDateStr)
	if err != nil {
		return time.Time{}, err
	}
	return mostRecentDate, nil
}
// Create an activity with the possibility of a merge
//
// Check for activities with the following criteria:
// 	1. same date as passed activity
// 	2. same activity type
// 	3. no completed data
//
// After that, if there is one activity, do the merge.
func (db *IActivityRepository) CreateOrMerge(activity *model.Activity, userUuid string) error {
    sql := sqlcode.SQLReader(sqlcode.Activity_CheckDate)
    rows, err := db.DB.Query(sql, activity.Date.Format(dateutil.Layout), userUuid, activity.Type.Id)
    if err != nil {
        return fmt.Errorf("error create/merge activity: %w", err)
    }
    dateActivityList := []*model.Activity{}
    for rows.Next() {
        // do stuff here
        tmpAct := model.EmptyActivity()
        tmpDateStr := ""

        err := rows.Scan(&tmpAct.Uuid,
			&tmpDateStr, &tmpAct.Order, &tmpAct.Name, &tmpAct.Description, &tmpAct.Notes, &tmpAct.IsRace, &tmpAct.NumStrides, &tmpAct.Map,
			&tmpAct.Type.Id, &tmpAct.Type.Name,
			&tmpAct.Planned.Distance.Length, &tmpAct.Planned.Distance.Unit, &tmpAct.Planned.Duration, &tmpAct.Planned.Vertical.Length, &tmpAct.Planned.Vertical.Unit,
			&tmpAct.Completed.Distance.Length, &tmpAct.Completed.Distance.Unit, &tmpAct.Completed.Duration, &tmpAct.Completed.Vertical.Length, &tmpAct.Completed.Vertical.Unit,
		)
		if err != nil {
			return fmt.Errorf("error scanning row: %w", err)
		} else {

			tmpDate, err := time.Parse(dateutil.TimeLayout, tmpDateStr)
			if err != nil {
				return fmt.Errorf("activity date failed to parse: %w", err)
			}
			tmpAct.Date = tmpDate
			tmpAct.SetUuid(tmpAct.Uuid)
			err2 := tmpAct.Validate()
			if err2 != nil {
				return fmt.Errorf("activity failed to validate: %w", err2)
			}
			dateActivityList = append(dateActivityList, tmpAct)
		}
    }
    if len(dateActivityList) != 0 {
        fmt.Println("activity list is long enough to consider activity merging")
    }
    // once we have all the existing activities for the day, we can check if any of them are eligible for merging.
    for _, act := range dateActivityList {
        // TODO: unsure if this logic is entirely correct.
        // probably need to be more sophisticated about the merging logic
        // two activities could be eligible for merger, but the first one gets chosen just b/c is appears first
        // it would make sense to extent the logic to merge based on which activity appears more similar in the planned data
        // e.g. 2 activities can be merged into an activity that has 5 miles planned
        // if one of the activities has 10 miles completed and the other has 4 miles completed
        // we should merge the 4 mile one with the 5 mile planned one
        // (all else being equal)
        fmt.Println("activity can merge value: ", act.CanMerge(activity))
        if act.CanMerge(activity) {
            oldUuid := act.Uuid
            fmt.Println("************* ACTIVITY MERGE *****************")
            fmt.Println("merging activity: " + act.Uuid + " -> " + activity.Uuid)
            act.Merge(activity)
            // replace the old uuid with the one that comes from the activity
            updateUuidSQL := sqlcode.SQLReader(sqlcode.Activity_UpdateUuid)
            err := db.DB.Execute(updateUuidSQL, activity.Uuid, oldUuid, activity.Uuid, oldUuid, activity.Uuid, oldUuid)
            if err != nil {
                return fmt.Errorf("failed to update uuid during merge: %w", err)
            }

            return db.Update(act)
        }
    }
    // if there are no merges, or the activity cannot be merged, then we need to create a new activity
    return db.Create(userUuid, activity)
}
