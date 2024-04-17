package controllers

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/jcocozza/cassidy-wails/internal/database"
	"github.com/jcocozza/cassidy-wails/internal/model"
	userrepo "github.com/jcocozza/cassidy-wails/internal/repository/userRepo"
	"github.com/jcocozza/cassidy-wails/internal/utils/dateutil"
	"github.com/jcocozza/cassidy-wails/internal/utils/uuidgen"
	"golang.org/x/oauth2"
)

type UserHandler struct {
	UserRepository userrepo.UserRepository
	User           *model.User
}
func NewUserHandler(db database.DbOperations, user *model.User) *UserHandler {
	return &UserHandler{
		UserRepository: userrepo.NewIUserRespository(db),
		User:           user,
	}
}
// Create a user
func (uh *UserHandler) CreateUser(createRequest *model.User) (*model.User, error) {
	_, err := uh.UserRepository.Read(createRequest.Username)

	if err != nil {
		// if there is no user
		if errors.Is(err, sql.ErrNoRows) {
			uuid := uuidgen.GenerateUUID()
			createRequest.SetUuid(uuid)
			err2 := createRequest.Validate()
			if err2 != nil {
				return nil, fmt.Errorf("user is invalid")
			}
			err3 := uh.UserRepository.Create(createRequest)
			if err3 != nil {
				fmt.Println("error creating user: " + err3.Error())
				return nil, err3
			}
			return createRequest, nil
		} else {
			return nil, fmt.Errorf("something went wrong creating user: %w", err)
		}
	} else {
		return nil, fmt.Errorf("user already exists")
	}
}
// the frontend will send this struct to us
type authRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
// authenticate a user
func (uh *UserHandler) AuthenticateUser(authRequest authRequest) (*model.User, error) {
	usr, err1 := uh.UserRepository.Read(authRequest.Username)

	if err1 != nil {
		return nil, err1
	}

	if usr.Password != authRequest.Password {
		return nil, fmt.Errorf("incorrect password")
	} else if usr.Password == authRequest.Password {
		return usr, nil
	} else {
		return nil, fmt.Errorf("unknown authorization failure")
	}
}

type MCCurrentDate struct {
	StartDate time.Time `json:"start_date" ts_type:"Date" ts_transform:"new Date(__VALUE__)"`
	EndDate   time.Time `json:"end_date" ts_type:"Date" ts_transform:"new Date(__VALUE__)"`
}
// Return the start date and end date of the current microcycle
func (uh *UserHandler) GetMicrocycleCurrentDates() *MCCurrentDate {
	currentDate := time.Now()
	var mc []time.Time
	if uh.User.CycleDays == 7 {
		mc = dateutil.GetDateMicrocycle(currentDate, uh.User.CycleStart, uh.User.CycleDays)
	} else {
		mc = dateutil.GetCurrentCycleFromInitialDate(uh.User.InitialCycleStart, uh.User.CycleDays)
	}

	mp := &MCCurrentDate{
		StartDate: mc[0],
		EndDate:   mc[len(mc)-1],
	}

	fmt.Println("CURRENT MICROCYCLE DATES:", mp)

	return mp
}
// Update the user preferences
func (uh *UserHandler) UpdateUser(updateRequest *model.User) (*model.User, error) {
	err := uh.UserRepository.Update(updateRequest)
	if err != nil {
		return nil, err
	}

	return updateRequest, nil
}
// Strava functions
func (uh *UserHandler) CreateStravaToken(user *model.User, token *oauth2.Token) error {
	err := uh.UserRepository.CreateStravaToken(user, token)
	if err != nil {
		return err
	}
	return nil
}
// if the token does not exist for the user, create it
func (uh *UserHandler) UpdateorCreateStravaToken(user *model.User, token *oauth2.Token) error {
	err := uh.UserRepository.UpdateStravaToken(user, token)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err := uh.UserRepository.CreateStravaToken(user, token)
			if err != nil {
				return fmt.Errorf("failed to create strava token after attempted update: %w", err)
			}
		} else {
			return err
		}
	}
	return nil
}
func (uh *UserHandler) GetStravaToken(user *model.User) (*oauth2.Token, error) {
	token, err := uh.UserRepository.ReadStravaToken(user)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("no user token exists: %w", err)
		} else {
			return nil, err
		}
	}
	return token, nil
}