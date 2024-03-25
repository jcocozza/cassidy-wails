package model

import (
	"fmt"

	"github.com/jcocozza/cassidy-wails/internal/utils/dateutil"
	"github.com/jcocozza/cassidy-wails/internal/utils/measurement"
)

type User struct {
	Uuid              string                `json:"uuid"`
	Username          string                `json:"username"`
	Password          string                `json:"password"`
	Units             measurement.UnitClass `json:"units"`
	CycleStart        string                `json:"cycle_start"`
	CycleDays         int                   `json:"cycle_days"`
	InitialCycleStart string                `json:"initial_cycle_start"`
}

// An empty user has no uuid, information and has cycle days of -1
func EmptyUser() *User {
	return &User{
		Uuid:              "",
		Username:          "",
		Password:          "",
		Units:             "",
		CycleStart:        "",
		CycleDays:         -1,
		InitialCycleStart: "",
	}
}

// Validate a user.
//
// Validation ensures that a user has a uuid, username, password, units, cycle start and a non-negative cycle days.
func (usr *User) Validate() error {
	if usr.Uuid == "" {
		return fmt.Errorf("user uuid is invalid")
	}
	if usr.Username == "" {
		return fmt.Errorf("user username is invalid")
	}
	if usr.Password == "" {
		return fmt.Errorf("user password is invalid")
	}
	err := measurement.ValidateUnitClass(usr.Units)
	if err != nil {
		return fmt.Errorf("user units are invalid: %w", err)
	}
	err2 := dateutil.ValidateDayOfWeek(usr.CycleStart)
	if err2 != nil {
		return fmt.Errorf("user day of week is invalid: %w", err2)
	}
	if usr.CycleDays == -1 {
		return fmt.Errorf("user cycle days must be positive")
	}
	return nil
}
func (usr *User) SetUuid(uuid string) {
	usr.Uuid = uuid
}

type NewUser struct {
	Username          string `json:"username"`
	Password          string `json:"password"`
	Units             string `json:"units"`
	CycleStart        string `json:"cycle_start"`
	CycleDays         int    `json:"cycle_days"`
	InitialCycleStart string `json:"initial_cycle_start"`
}
