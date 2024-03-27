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
			uh.UserRepository.Create(createRequest)
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
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}
// Return the start date and end date of the current microcycle
func (uh *UserHandler) GetMicrocycleCurrentDates() *MCCurrentDate {
	currentDate := time.Now().Format(dateutil.Layout)
	var mc []*dateutil.DateObject
	if uh.User.CycleDays == 7 {
		mc = dateutil.GetDateMicrocycle(currentDate, uh.User.CycleStart, uh.User.CycleDays)
	} else {
		mc = dateutil.GetCurrentCycleFromInitialDate(uh.User.InitialCycleStart, uh.User.CycleDays)
	}

	mp := &MCCurrentDate{
		StartDate: mc[0].Date,
		EndDate:   mc[len(mc)-1].Date,
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
