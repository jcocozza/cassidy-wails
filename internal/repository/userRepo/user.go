package userrepo

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/jcocozza/cassidy-wails/internal/database"
	"github.com/jcocozza/cassidy-wails/internal/model"
	"github.com/jcocozza/cassidy-wails/internal/sqlcode"
	"github.com/jcocozza/cassidy-wails/internal/utils"
)

// The methods for interacting with user objects
type UserRepository interface {
	Create(user *model.User) error
	Read(username string) (*model.User, error)
	ReadPreferences(userUuid string) (string, int, string, error)
	Update(user *model.User) error
	//Delete(user *User) error
}

// Represents a database connection
type IUserRepository struct {
	DB database.DbOperations
}

func NewIUserRespository(db database.DbOperations) *IUserRepository {
	return &IUserRepository{
		DB: db,
	}
}

// Insert the user object into the database.
func (db *IUserRepository) Create(user *model.User) error {
	sql := utils.SQLReader(sqlcode.User_create)

	err1 := user.Validate()
	if err1 != nil {
		return fmt.Errorf("planned creation failed to validate: %w", err1)
	}

	err := db.DB.Execute(sql, user.Uuid, user.Username, user.Password, user.Units, user.CycleStart, user.CycleDays, user.InitialCycleStart)
	if err != nil {
		return fmt.Errorf("user creation failed: %w", err)
	}

	return nil
}

// Read user object from the database for a given username
func (db *IUserRepository) Read(username string) (*model.User, error) {
	query := utils.SQLReader(sqlcode.User_read)
	row := db.DB.QueryRow(query, username)

	usr := model.EmptyUser()

	err := row.Scan(&usr.Uuid, &usr.Username, &usr.Password, &usr.Units, &usr.CycleStart, &usr.CycleDays, &usr.InitialCycleStart)
	if err != nil {
		// in the case that no results are returned, we want to know that b/c that means there are no users with that username
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		return nil, fmt.Errorf("user read failed: %w", err)
	}
	err2 := usr.Validate()
	if err2 != nil {
		return nil, fmt.Errorf("user validation failed: %w", err2)
	}
	return usr, nil
}

// Read the user start date, number of cycle day preferences and initial start date
func (db *IUserRepository) ReadPreferences(userUuid string) (string, int, string, error) {
	sql := utils.SQLReader(sqlcode.User_preferences)
	row := db.DB.QueryRow(sql, userUuid)

	var cycleStart string
	var cycleDays int
	var initialStartDate string
	err := row.Scan(&cycleStart, &cycleDays, &initialStartDate)

	if err != nil {
		return "", -1, "", fmt.Errorf("error scanning user preferences: %w", err)
	}

	return cycleStart, cycleDays, initialStartDate, nil
}

// Update the user preferences
func (db *IUserRepository) Update(user *model.User) error {
	sql := utils.SQLReader(sqlcode.User_update)
	err := db.DB.Execute(sql, user.Units, user.CycleStart, user.CycleDays, user.InitialCycleStart, user.Uuid)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}
