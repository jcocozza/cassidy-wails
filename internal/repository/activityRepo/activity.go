package activityrepo

import (
	"fmt"

	"github.com/jcocozza/cassidy-wails/internal/database"
	"github.com/jcocozza/cassidy-wails/internal/model"
	"github.com/jcocozza/cassidy-wails/internal/sqlcode"
)

// Functions for working with activity objects.
type ActivityRepository interface {
	// Activity

	Create(userUuid string, activity *model.Activity) error
	//Read(activityUuid string) (*model.Activity, error)
	Update(completed *model.Activity) error
	Delete(activityUuid string) error
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
	err := db.DB.Execute(sqlActivity, activity.Uuid, userUuid, activity.Date, activity.Order, activity.Name, activity.Description, activity.Notes, activity.Type.Id, activity.IsRace, activity.NumStrides)
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

// Update an activity
//
// Note: a change in the activity_type_id will trigger a delete of all activity type subtypes
//
// Updates the activity and its planned/completed
func (db *IActivityRepository) Update(activity *model.Activity) error {
	sqlActivity := sqlcode.SQLReader(sqlcode.Activity_update)
	err := db.DB.Execute(sqlActivity, activity.Date, activity.Order, activity.Name, activity.Description, activity.Notes, activity.Type.Id, activity.IsRace, activity.NumStrides, activity.Uuid)

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
