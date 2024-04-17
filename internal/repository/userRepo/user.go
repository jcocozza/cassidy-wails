package userrepo

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/jcocozza/cassidy-wails/internal/database"
	"github.com/jcocozza/cassidy-wails/internal/model"
	"github.com/jcocozza/cassidy-wails/internal/sqlcode"
	"github.com/jcocozza/cassidy-wails/internal/utils/dateutil"
	"golang.org/x/oauth2"
)

// The methods for interacting with user objects
type UserRepository interface {
	Create(user *model.User) error
	Read(username string) (*model.User, error)
	//ReadPreferences(userUuid string) (string, int, string, error)
	Update(user *model.User) error
	//Delete(user *User) error
	CreateStravaToken(user *model.User, token *oauth2.Token) error
	ReadStravaToken(user *model.User) (*oauth2.Token, error)
	UpdateStravaToken(user *model.User, token *oauth2.Token) error
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
	sql := sqlcode.SQLReader(sqlcode.User_create)

	err1 := user.Validate()
	if err1 != nil {
		return fmt.Errorf("planned creation failed to validate: %w", err1)
	}

	err := db.DB.Execute(sql, user.Uuid, user.Username, user.Password, user.Units, user.CycleStart, user.CycleDays, user.InitialCycleStart.Format(dateutil.Layout))
	if err != nil {
		return fmt.Errorf("user creation failed: %w", err)
	}
	return nil
}
// Read user object from the database for a given username
func (db *IUserRepository) Read(username string) (*model.User, error) {
	query := sqlcode.SQLReader(sqlcode.User_read)
	row := db.DB.QueryRow(query, username)
	usr := model.EmptyUser()

	var initialCycleStartStr string
	err := row.Scan(&usr.Uuid, &usr.Username, &usr.Password, &usr.Units, &usr.CycleStart, &usr.CycleDays, &initialCycleStartStr)
	if err != nil {
		// in the case that no results are returned, we want to know that b/c that means there are no users with that username
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		return nil, fmt.Errorf("user read failed: %w", err)
	}

	if initialCycleStartStr != "" {
		initialCycleStart, err15 := time.Parse(dateutil.Layout, initialCycleStartStr)
		if err15 != nil {
			return nil, err15
		}
		usr.InitialCycleStart = initialCycleStart
	}

	err2 := usr.Validate()
	if err2 != nil {
		return nil, fmt.Errorf("user validation failed: %w", err2)
	}
	return usr, nil
}
/*
TODO: Depreciate this
// Read the user start date, number of cycle day preferences and initial start date
func (db *IUserRepository) ReadPreferences(userUuid string) (string, int, string, error) {
	sql := sqlcode.SQLReader(sqlcode.User_preferences)
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
*/
// Update the user preferences
func (db *IUserRepository) Update(user *model.User) error {
	sql := sqlcode.SQLReader(sqlcode.User_update)
	err := db.DB.Execute(sql, user.Units, user.CycleStart, user.CycleDays, user.InitialCycleStart.Format(dateutil.Layout), user.Uuid)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}
// Create a strava token for the user
func (db *IUserRepository) CreateStravaToken(user *model.User, token *oauth2.Token) error {
	sql := sqlcode.SQLReader(sqlcode.Strava_token_create)
	err := db.DB.Execute(sql, user.Uuid, token.AccessToken, token.TokenType, token.RefreshToken, token.Expiry.Format(dateutil.TokenLayout))
	if err != nil {
		return fmt.Errorf("failed to create strava token: %w", err)
	}
	return nil
}
// Read the strava token stored in the database for a given user
func (db *IUserRepository) ReadStravaToken(user *model.User) (*oauth2.Token, error) {
	sql := sqlcode.SQLReader(sqlcode.Strava_token_read)
	// there should only ever be 1 user token
	row := db.DB.QueryRow(sql, user.Uuid)
	var accessToken, tokenType, refreshToken, expiryStr string
	err := row.Scan(accessToken, tokenType, refreshToken, expiryStr)
	if err != nil {
		return nil, err
	}
	expiry, err := time.Parse(dateutil.TokenLayout, expiryStr)
	if err != nil {
		return nil, err
	}
	return &oauth2.Token{
		AccessToken: accessToken,
		TokenType: tokenType,
		RefreshToken: refreshToken,
		Expiry: expiry,
	}, nil
}
// Update the user's strava token
func (db *IUserRepository) UpdateStravaToken(user *model.User, token *oauth2.Token) error {
	sql := sqlcode.SQLReader(sqlcode.Strava_token_update)
	err := db.DB.Execute(sql, token.AccessToken, token.TokenType, token.RefreshToken, token.Expiry.Format(dateutil.TokenLayout), user.Uuid)
	if err != nil {
		return fmt.Errorf("failed to update strava token: %w", err)
	}
	return nil
}
