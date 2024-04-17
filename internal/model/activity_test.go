package model

import (
	"reflect"
	"testing"
	"time"

	colorutil "github.com/jcocozza/cassidy-wails/internal/utils/colorUtil"
	"github.com/jcocozza/cassidy-wails/internal/utils/measurement"
)

func TestActivityType_Validate(t *testing.T) {
	type fields struct {
		Id   int
		Name string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"valid", fields{Id: 1, Name: "Run"}, false},
		{"invalid id", fields{Id: -1, Name: "Run"}, true},
		{"invalid name", fields{Id: 1, Name: ""}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			at := &ActivityType{
				Id:   tt.fields.Id,
				Name: tt.fields.Name,
			}
			if err := at.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("ActivityType.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestActivitySubtype_Validate(t *testing.T) {
	type fields struct {
		Id          int
		SuperTypeId int
		Name        string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"valid", fields{Id: 1, SuperTypeId: 1, Name: "Long"}, false},
		{"invalid id", fields{Id: -1, SuperTypeId: 1, Name: "Long"}, true},
		{"invalid supertype id", fields{Id: 1, SuperTypeId: -1, Name: "Long"}, true},
		{"invalid name", fields{Id: 1, SuperTypeId: -1, Name: ""}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			as := &ActivitySubtype{
				Id:          tt.fields.Id,
				SuperTypeId: tt.fields.SuperTypeId,
				Name:        tt.fields.Name,
			}
			if err := as.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("ActivitySubtype.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestActivityTypeSubtype_Validate(t *testing.T) {
	type fields struct {
		Id              int
		ActivityUuid    string
		ActivityType    *ActivityType
		ActivitySubtype *ActivitySubtype
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"valid", fields{Id: 1, ActivityUuid: "c63fd1dc-cb8a-11ee-90d8-325096b39f47", ActivityType: &ActivityType{Id: 1, Name: "Run"}, ActivitySubtype: &ActivitySubtype{Id: 1, SuperTypeId: 1, Name: "Long"}}, false},
		{"invalid activity id", fields{Id: -1, ActivityUuid: "c63fd1dc-cb8a-11ee-90d8-325096b39f47", ActivityType: &ActivityType{Id: 1, Name: "Run"}, ActivitySubtype: &ActivitySubtype{Id: 1, SuperTypeId: 1, Name: "Long"}}, true},
		{"invalid activity uuid", fields{Id: 1, ActivityUuid: "", ActivityType: &ActivityType{Id: 1, Name: "Run"}, ActivitySubtype: &ActivitySubtype{Id: 1, SuperTypeId: 1, Name: "Long"}}, true},
		{"invalid activity type", fields{Id: 1, ActivityUuid: "", ActivityType: &ActivityType{Id: -1, Name: "Run"}, ActivitySubtype: &ActivitySubtype{Id: 1, SuperTypeId: 1, Name: "Long"}}, true},
		{"invalid activity subtype", fields{Id: 1, ActivityUuid: "", ActivityType: &ActivityType{Id: -1, Name: "Run"}, ActivitySubtype: &ActivitySubtype{Id: -1, SuperTypeId: 1, Name: "Long"}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ats := &ActivityTypeSubtype{
				Id:              tt.fields.Id,
				ActivityUuid:    tt.fields.ActivityUuid,
				ActivityType:    tt.fields.ActivityType,
				ActivitySubtype: tt.fields.ActivitySubtype,
			}
			if err := ats.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("ActivityTypeSubtype.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestActivity_Validate(t *testing.T) {
	type fields struct {
		Uuid            string
		Date            time.Time
		Order           int
		Name            string
		Description     string
		Notes           string
		Type            *ActivityType
		TypeSubtypeList []*ActivityTypeSubtype
		EquipmentList   []*ActivityEquipment
		Planned         *Planned
		Completed       *Completed
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"valid", fields{Uuid: "c63fd240-cb8a-11ee-bb77-325096b39f47", Date: time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC), Order: 1, Name: "", Description: "", Notes: "", Type: &ActivityType{Id: 1, Name: "Run"}, TypeSubtypeList: []*ActivityTypeSubtype{}, EquipmentList: []*ActivityEquipment{}, Planned: &Planned{ActivityUuid: "c63fd240-cb8a-11ee-bb77-325096b39f47", Distance: measurement.CreateMeasurement(measurement.Mile, 10), Duration: 5000, Vertical: measurement.CreateMeasurement(measurement.Feet, 1500)}, Completed: &Completed{ActivityUuid: "c63fd240-cb8a-11ee-bb77-325096b39f47", Distance: measurement.CreateMeasurement(measurement.Mile, 10), Duration: 5000, Vertical: measurement.CreateMeasurement(measurement.Feet, 1500)}}, false},
		{"invalid uuid", fields{Uuid: "", Date: time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC), Order: 1, Name: "", Description: "", Notes: "", Type: &ActivityType{Id: 1, Name: "Run"}, TypeSubtypeList: []*ActivityTypeSubtype{}, EquipmentList: []*ActivityEquipment{}, Planned: &Planned{ActivityUuid: "c63fd240-cb8a-11ee-bb77-325096b39f47", Distance: measurement.CreateMeasurement(measurement.Mile, 10), Duration: 5000, Vertical: measurement.CreateMeasurement(measurement.Feet, 1500)}, Completed: &Completed{ActivityUuid: "c63fd240-cb8a-11ee-bb77-325096b39f47", Distance: measurement.CreateMeasurement(measurement.Mile, 10), Duration: 5000, Vertical: measurement.CreateMeasurement(measurement.Feet, 1500)}}, true},
		{"invalid date", fields{Uuid: "c63fd240-cb8a-11ee-bb77-325096b39f47", Date: time.Time{}, Order: 1, Name: "", Description: "", Notes: "", Type: &ActivityType{Id: 1, Name: "Run"}, TypeSubtypeList: []*ActivityTypeSubtype{}, EquipmentList: []*ActivityEquipment{}, Planned: &Planned{ActivityUuid: "c63fd240-cb8a-11ee-bb77-325096b39f47", Distance: measurement.CreateMeasurement(measurement.Mile, 10), Duration: 5000, Vertical: measurement.CreateMeasurement(measurement.Feet, 1500)}, Completed: &Completed{ActivityUuid: "c63fd240-cb8a-11ee-bb77-325096b39f47", Distance: measurement.CreateMeasurement(measurement.Mile, 10), Duration: 5000, Vertical: measurement.CreateMeasurement(measurement.Feet, 1500)}}, true},
		{"invalid date2", fields{Uuid: "c63fd240-cb8a-11ee-bb77-325096b39f47", Date: time.Time{}, Order: 1, Name: "", Description: "", Notes: "", Type: &ActivityType{Id: 1, Name: "Run"}, TypeSubtypeList: []*ActivityTypeSubtype{}, EquipmentList: []*ActivityEquipment{}, Planned: &Planned{ActivityUuid: "c63fd240-cb8a-11ee-bb77-325096b39f47", Distance: measurement.CreateMeasurement(measurement.Mile, 10), Duration: 5000, Vertical: measurement.CreateMeasurement(measurement.Feet, 1500)}, Completed: &Completed{ActivityUuid: "c63fd240-cb8a-11ee-bb77-325096b39f47", Distance: measurement.CreateMeasurement(measurement.Mile, 10), Duration: 5000, Vertical: measurement.CreateMeasurement(measurement.Feet, 1500)}}, true},
		{"invalid order", fields{Uuid: "c63fd240-cb8a-11ee-bb77-325096b39f47", Date: time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC), Order: -1, Name: "", Description: "", Notes: "", Type: &ActivityType{Id: 1, Name: "Run"}, TypeSubtypeList: []*ActivityTypeSubtype{}, EquipmentList: []*ActivityEquipment{}, Planned: &Planned{ActivityUuid: "c63fd240-cb8a-11ee-bb77-325096b39f47", Distance: measurement.CreateMeasurement(measurement.Mile, 10), Duration: 5000, Vertical: measurement.CreateMeasurement(measurement.Feet, 1500)}, Completed: &Completed{ActivityUuid: "c63fd240-cb8a-11ee-bb77-325096b39f47", Distance: measurement.CreateMeasurement(measurement.Mile, 10), Duration: 5000, Vertical: measurement.CreateMeasurement(measurement.Feet, 1500)}}, true},
		{"invalid type", fields{Uuid: "c63fd240-cb8a-11ee-bb77-325096b39f47", Date: time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC), Order: 1, Name: "", Description: "", Notes: "", Type: &ActivityType{Id: -1, Name: "Run"}, TypeSubtypeList: []*ActivityTypeSubtype{}, EquipmentList: []*ActivityEquipment{}, Planned: &Planned{ActivityUuid: "c63fd240-cb8a-11ee-bb77-325096b39f47", Distance: measurement.CreateMeasurement(measurement.Mile, 10), Duration: 5000, Vertical: measurement.CreateMeasurement(measurement.Feet, 1500)}, Completed: &Completed{ActivityUuid: "c63fd240-cb8a-11ee-bb77-325096b39f47", Distance: measurement.CreateMeasurement(measurement.Mile, 10), Duration: 5000, Vertical: measurement.CreateMeasurement(measurement.Feet, 1500)}}, true},
		{"invalid planned", fields{Uuid: "c63fd240-cb8a-11ee-bb77-325096b39f47", Date: time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC), Order: 1, Name: "", Description: "", Notes: "", Type: &ActivityType{Id: 1, Name: "Run"}, TypeSubtypeList: []*ActivityTypeSubtype{}, EquipmentList: []*ActivityEquipment{}, Planned: &Planned{ActivityUuid: "c63fd240-cb8a-11ee-bb77-325096b39f47", Distance: measurement.CreateMeasurement(measurement.Mile, -10), Duration: 5000, Vertical: measurement.CreateMeasurement(measurement.Feet, 1500)}, Completed: &Completed{ActivityUuid: "c63fd240-cb8a-11ee-bb77-325096b39f47", Distance: measurement.CreateMeasurement(measurement.Mile, 10), Duration: 5000, Vertical: measurement.CreateMeasurement(measurement.Feet, 1500)}}, true},
		{"invalid completed", fields{Uuid: "c63fd240-cb8a-11ee-bb77-325096b39f47", Date: time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC), Order: 1, Name: "", Description: "", Notes: "", Type: &ActivityType{Id: 1, Name: "Run"}, TypeSubtypeList: []*ActivityTypeSubtype{}, EquipmentList: []*ActivityEquipment{}, Planned: &Planned{ActivityUuid: "c63fd240-cb8a-11ee-bb77-325096b39f47", Distance: measurement.CreateMeasurement(measurement.Mile, 10), Duration: 5000, Vertical: measurement.CreateMeasurement(measurement.Feet, 1500)}, Completed: &Completed{ActivityUuid: "c63fd240-cb8a-11ee-bb77-325096b39f47", Distance: measurement.CreateMeasurement(measurement.Mile, -10), Duration: 5000, Vertical: measurement.CreateMeasurement(measurement.Feet, 1500)}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Activity{
				Uuid:            tt.fields.Uuid,
				Date:            tt.fields.Date,
				Order:           tt.fields.Order,
				Name:            tt.fields.Name,
				Description:     tt.fields.Description,
				Notes:           tt.fields.Notes,
				Type:            tt.fields.Type,
				TypeSubtypeList: tt.fields.TypeSubtypeList,
				EquipmentList:   tt.fields.EquipmentList,
				Planned:         tt.fields.Planned,
				Completed:       tt.fields.Completed,
			}
			if err := a.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Activity.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCycle_AddActivity(t *testing.T) {
	type args struct {
		act *Activity
	}
	tests := []struct {
		name    string
		c       *Cycle
		args    args
		wantErr bool
	}{
		{"valid", &Cycle{}, args{act: &Activity{Uuid: "", Date: time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC), Order: -1, Name: "", Description: "", Notes: "", Type: EmptyActivityType(), TypeSubtypeList: []*ActivityTypeSubtype{}, EquipmentList: []*ActivityEquipment{}, Planned: EmptyPlanned(), Completed: EmptyCompleted()}}, false},
		{"invalid", &Cycle{}, args{act: &Activity{Uuid: "", Date: time.Time{}, Order: -1, Name: "", Description: "", Notes: "", Type: EmptyActivityType(), TypeSubtypeList: []*ActivityTypeSubtype{}, EquipmentList: []*ActivityEquipment{}, Planned: EmptyPlanned(), Completed: EmptyCompleted()}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.AddActivity(tt.args.act); (err != nil) != tt.wantErr {
				t.Errorf("Cycle.AddActivity() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCycle_CreateUuidLists(t *testing.T) {
	tests := []struct {
		name string
		c    *Cycle
		want [][]string
	}{
		{"empty cycle", &Cycle{}, [][]string{}},
		{"cycle 2 days", &Cycle{&ActivityList{ActivityList: []*Activity{{Uuid: "c63fce3a-cb8a-11ee-857a-325096b39f47"}, {Uuid: "c63fd0ce-cb8a-11ee-ac1b-325096b39f47"}}}, &ActivityList{ActivityList: []*Activity{{Uuid: "c63fd164-cb8a-11ee-9cc0-325096b39f47"}, {Uuid: "c63fd1dc-cb8a-11ee-90d8-325096b39f47"}}}}, [][]string{{"c63fce3a-cb8a-11ee-857a-325096b39f47", "c63fd0ce-cb8a-11ee-ac1b-325096b39f47"}, {"c63fd164-cb8a-11ee-9cc0-325096b39f47", "c63fd1dc-cb8a-11ee-90d8-325096b39f47"}}},
		{"cycle 1 days", &Cycle{&ActivityList{ActivityList: []*Activity{{Uuid: "c63fce3a-cb8a-11ee-857a-325096b39f47"}, {Uuid: "c63fd0ce-cb8a-11ee-ac1b-325096b39f47"}}}}, [][]string{{"c63fce3a-cb8a-11ee-857a-325096b39f47", "c63fd0ce-cb8a-11ee-ac1b-325096b39f47"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.CreateUuidLists(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cycle.CreateUuidLists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCycle_CreateUuidList(t *testing.T) {
	tests := []struct {
		name string
		c    *Cycle
		want []string
	}{
		{"empty cycle", &Cycle{}, []string{}},
		{"cycle 2 days", &Cycle{&ActivityList{ActivityList: []*Activity{{Uuid: "c63fce3a-cb8a-11ee-857a-325096b39f47"}, {Uuid: "c63fd0ce-cb8a-11ee-ac1b-325096b39f47"}}}, &ActivityList{ActivityList: []*Activity{{Uuid: "c63fd164-cb8a-11ee-9cc0-325096b39f47"}, {Uuid: "c63fd1dc-cb8a-11ee-90d8-325096b39f47"}}}}, []string{"c63fce3a-cb8a-11ee-857a-325096b39f47", "c63fd0ce-cb8a-11ee-ac1b-325096b39f47", "c63fd164-cb8a-11ee-9cc0-325096b39f47", "c63fd1dc-cb8a-11ee-90d8-325096b39f47"}},
		{"cycle 1 days", &Cycle{&ActivityList{ActivityList: []*Activity{{Uuid: "c63fce3a-cb8a-11ee-857a-325096b39f47"}, {Uuid: "c63fd0ce-cb8a-11ee-ac1b-325096b39f47"}}}}, []string{"c63fce3a-cb8a-11ee-857a-325096b39f47", "c63fd0ce-cb8a-11ee-ac1b-325096b39f47"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.CreateUuidList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cycle.CreateUuidList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestActivity_CompletionColor(t *testing.T) {
	type fields struct {
		Uuid      string
		Date      time.Time
		Planned   *Planned
		Completed *Completed
		Color     colorutil.Color
	}
	tests := []struct {
		name    string
		fields  fields
		want    colorutil.Color
		wantErr bool
	}{
		{"test", fields{Uuid: "", Date: time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC), Planned: &Planned{Distance: measurement.CreateMeasurement("m", 10), Duration: 10, Vertical: measurement.CreateMeasurement("m", 10)}, Completed: &Completed{Distance: measurement.CreateMeasurement("m", 10), Duration: 10, Vertical: measurement.CreateMeasurement("m", 10)}}, colorutil.Green, false},
		{"blue no planned", fields{Uuid: "", Date: time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC), Planned: &Planned{Distance: measurement.CreateMeasurement("m", 0), Duration: 0, Vertical: measurement.CreateMeasurement("m", 0)}, Completed: &Completed{Distance: measurement.CreateMeasurement("m", 10), Duration: 10, Vertical: measurement.CreateMeasurement("m", 10)}}, colorutil.Blue, false},
		{"test", fields{Uuid: "", Date: time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC), Planned: &Planned{Distance: measurement.CreateMeasurement("m", 10), Duration: 0, Vertical: measurement.CreateMeasurement("m", 10)}, Completed: &Completed{Distance: measurement.CreateMeasurement("m", 10), Duration: 10, Vertical: measurement.CreateMeasurement("m", 10)}}, colorutil.Green, false},
		{"red warning", fields{Uuid: "", Date: time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC), Planned: &Planned{Distance: measurement.CreateMeasurement("m", 10), Duration: 10, Vertical: measurement.CreateMeasurement("m", 10)}, Completed: &Completed{Distance: measurement.CreateMeasurement("m", 10), Duration: 15.1, Vertical: measurement.CreateMeasurement("m", 10)}}, colorutil.Red, false},
		{"yellow warning", fields{Uuid: "", Date: time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC), Planned: &Planned{Distance: measurement.CreateMeasurement("m", 10), Duration: 10, Vertical: measurement.CreateMeasurement("m", 10)}, Completed: &Completed{Distance: measurement.CreateMeasurement("m", 10), Duration: 15.0, Vertical: measurement.CreateMeasurement("m", 10)}}, colorutil.Yellow, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Activity{
				Uuid:      tt.fields.Uuid,
				Date:      tt.fields.Date,
				Planned:   tt.fields.Planned,
				Completed: tt.fields.Completed,
				Color:     tt.fields.Color,
			}
			got, err := a.CompletionColor()
			if (err != nil) != tt.wantErr {
				t.Errorf("Activity.CompletionColor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Activity.CompletionColor() = %v, want %v", got, tt.want)
			}
		})
	}
}
