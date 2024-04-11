package plannedrepo

import (
	"fmt"

	"github.com/jcocozza/cassidy-wails/internal/database"
	"github.com/jcocozza/cassidy-wails/internal/model"
	"github.com/jcocozza/cassidy-wails/internal/sqlcode"
)

// The methods for interacting with planned objects
type PlannedRepository interface {
	Create(planned *model.Planned) error
	Read(activityUuid string) (*model.Planned, error)
	Update(planned *model.Planned) error
	Delete(planned *model.Planned) error
}

// Represents a SQLite database connection
type IPlannedRepository struct {
	DB database.DbOperations
}

// Insert the planned object into the database.
func (db *IPlannedRepository) Create(planned *model.Planned) error {
	sql := sqlcode.SQLReader(sqlcode.Planned_create)

	err1 := planned.Validate()
	if err1 != nil {
		return fmt.Errorf("planned creation failed to validate: %w", err1)
	}

	err := db.DB.Execute(sql, planned.ActivityUuid, planned.Distance.Length, planned.Distance.Unit, planned.Duration, planned.Vertical.Length, planned.Vertical.Unit)
	if err != nil {
		return fmt.Errorf("planned creation failed: %w", err)
	}

	return nil
}

// Read planned object from the database for a given activity_uuid.
func (db *IPlannedRepository) Read(activityUuid string) (*model.Planned, error) {
	sql := sqlcode.SQLReader(sqlcode.Planned_read)
	row := db.DB.QueryRow(sql, activityUuid)

	planned := model.EmptyPlanned()

	err := row.Scan(&planned.ActivityUuid, &planned.Distance.Length, &planned.Distance.Unit, &planned.Duration, &planned.Vertical.Length, &planned.Vertical.Unit)
	if err != nil {
		return nil, fmt.Errorf("planned read failed: %w", err)
	}
	err2 := planned.Validate()
	if err2 != nil {
		return nil, fmt.Errorf("planned validation failed: %w", err2)
	}
	return planned, nil
}

// Update planned object in the database(for a given activity uuid).
func (db *IPlannedRepository) Update(planned *model.Planned) error {
	err1 := planned.Validate()
	if err1 != nil {
		return fmt.Errorf("planned update failed to validate: %w", err1)
	}

	sql := sqlcode.SQLReader(sqlcode.Planned_update)
	err := db.DB.Execute(sql, planned.Distance, planned.Distance.Length, planned.Duration, planned.Vertical, planned.Vertical.Length, planned.ActivityUuid)
	if err != nil {
		return fmt.Errorf("planned update failed: %w", err)
	}
	return nil
}

// Delete a planned object in the database(by uuid)
//
// ! Warning: Not implemented: (planned should not be manually deleted from a command, the database handles deletions)
func (db *IPlannedRepository) Delete(planned *model.Planned) error {
	return fmt.Errorf("planned delete not implemented. are you sure the planned object should be deleted in this way?")
}
