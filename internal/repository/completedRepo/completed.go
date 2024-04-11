package completedrepo

import (
	"fmt"

	"github.com/jcocozza/cassidy-wails/internal/database"
	"github.com/jcocozza/cassidy-wails/internal/model"
	"github.com/jcocozza/cassidy-wails/internal/sqlcode"
)

// The methods for interacting with completed objects
type CompletedRepository interface {
	Create(completed *model.Completed) error
	Read(activityUuid string) (*model.Completed, error)
	Update(completed *model.Completed) error
	Delete(completed *model.Completed) error
}

// Represents a database connection
type ICompletedRepository struct {
	DB database.DbOperations
}

// Insert the completed object into the database.
func (db *ICompletedRepository) Create(completed *model.Completed) error {
	sql := sqlcode.SQLReader(sqlcode.Completed_create)

	err1 := completed.Validate()
	if err1 != nil {
		return fmt.Errorf("completed creation failed to validate: %w", err1)
	}

	err := db.DB.Execute(sql, completed.ActivityUuid, completed.Distance.Length, completed.Distance.Unit, completed.Duration, completed.Vertical.Length, completed.Vertical.Unit)
	if err != nil {
		return fmt.Errorf("completed creation failed: %w", err)
	}

	return nil
}

// Read completed object from the database for a given activity_uuid.
func (db *ICompletedRepository) Read(activityUuid string) (*model.Completed, error) {
	sql := sqlcode.SQLReader(sqlcode.Completed_read)
	row := db.DB.QueryRow(sql, activityUuid)

	completed := model.EmptyCompleted()

	err := row.Scan(&completed.ActivityUuid, &completed.Distance.Length, &completed.Distance.Unit, &completed.Duration, &completed.Vertical.Length, &completed.Vertical.Unit)
	if err != nil {
		return nil, fmt.Errorf("completed read failed: %w", err)
	}
	err2 := completed.Validate()
	if err2 != nil {
		return nil, fmt.Errorf("completed validation failed: %w", err2)
	}
	return completed, nil
}

// Update completed object in the database(for a given activity uuid).
func (db *ICompletedRepository) Update(completed *model.Completed) error {
	err1 := completed.Validate()
	if err1 != nil {
		return fmt.Errorf("completed update failed to validate: %w", err1)
	}

	sql := sqlcode.SQLReader(sqlcode.Completed_update)
	err := db.DB.Execute(sql, completed.Distance.Length, completed.Distance.Unit, completed.Duration, completed.Vertical.Length, completed.Vertical.Unit, completed.ActivityUuid)
	if err != nil {
		return fmt.Errorf("completed update failed: %w", err)
	}
	return nil
}

// Delete a completed object in the database(by uuid)
//
// ! Warning: Not implemented: (completed should not be manually deleted from a command, the database handles deletions)
func (db *ICompletedRepository) Delete(completed *model.Completed) error {
	return fmt.Errorf("completed delete not implemented. are you sure the completed object should be deleted in this way?")
}
