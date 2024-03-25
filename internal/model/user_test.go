package model

import (
	"testing"

	"github.com/jcocozza/cassidy-wails/internal/utils/dateutil"
	"github.com/jcocozza/cassidy-wails/internal/utils/measurement"
)

func TestUser_Validate(t *testing.T) {
	type fields struct {
		Uuid       string
		Username   string
		Password   string
		Units      measurement.UnitClass
		CycleStart string
		CycleDays  int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"valid", fields{Uuid: "c2cb0fde-cac8-11ee-90da-325096b39f47", Username: "username", Password: "password", Units: measurement.Imperial, CycleStart: dateutil.Monday, CycleDays: 7}, false},
		{"invalid uuid", fields{Uuid: "", Username: "username", Password: "password", Units: measurement.Imperial, CycleStart: dateutil.Monday, CycleDays: 7}, true},
		{"invalid username", fields{Uuid: "c2cb0fde-cac8-11ee-90da-325096b39f47", Username: "", Password: "password", Units: measurement.Imperial, CycleStart: dateutil.Monday, CycleDays: 7}, true},
		{"invalid password", fields{Uuid: "c2cb0fde-cac8-11ee-90da-325096b39f47", Username: "username", Password: "", Units: measurement.Imperial, CycleStart: dateutil.Monday, CycleDays: 7}, true},
		{"invalid units", fields{Uuid: "c2cb0fde-cac8-11ee-90da-325096b39f47", Username: "username", Password: "password", Units: "kg", CycleStart: dateutil.Monday, CycleDays: 7}, true},
		{"invalid cycle start", fields{Uuid: "c2cb0fde-cac8-11ee-90da-325096b39f47", Username: "username", Password: "password", Units: measurement.Imperial, CycleStart: "wed", CycleDays: 7}, true},
		{"invalid cycle days", fields{Uuid: "c2cb0fde-cac8-11ee-90da-325096b39f47", Username: "username", Password: "password", Units: measurement.Imperial, CycleStart: dateutil.Monday, CycleDays: -1}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usr := &User{
				Uuid:       tt.fields.Uuid,
				Username:   tt.fields.Username,
				Password:   tt.fields.Password,
				Units:      tt.fields.Units,
				CycleStart: tt.fields.CycleStart,
				CycleDays:  tt.fields.CycleDays,
			}
			if err := usr.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("User.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
