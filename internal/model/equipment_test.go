package model

import (
	"testing"

	"github.com/jcocozza/cassidy-wails/internal/utils/measurement"
)

func TestEquipmentType_Validate(t *testing.T) {
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
			et := &EquipmentType{
				Id:   tt.fields.Id,
				Name: tt.fields.Name,
			}
			if err := et.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("EquipmentType.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEquipment_Validate(t *testing.T) {
	type fields struct {
		Id            int
		UserUuid      string
		EquipmentType *EquipmentType
		Name          string
		Brand         string
		Model         string
		Cost          float64
		Size          string
		PurchaseDate  string
		Notes         string
		Mileage       *measurement.Measurement
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"valid", fields{Id: 1, UserUuid: "c2cb0cd2-cac8-11ee-b5e8-325096b39f47", EquipmentType: &EquipmentType{Id: 1, Name: "Shoes"}, Name: "name", Brand: "brand", Model: "model", Cost: 150, Size: "14", PurchaseDate: "2020-01-01", Notes: "", Mileage: measurement.CreateMeasurement(measurement.Mile, 0)}, false},
		//{"invalid user uuid", fields{Id: 1, UserUuid: "", EquipmentType: &EquipmentType{Id: 1, Name: "Shoes"}, Name: "name", Brand: "brand", Model: "model", Cost: 150, Size: "14", PurchaseDate: "2020-01-01", Notes: "", Mileage: measurement.CreateMeasurement(measurement.Mile, 100)}, true},
		{"invalid equipment type id", fields{Id: 1, UserUuid: "c2cb0cd2-cac8-11ee-b5e8-325096b39f47", EquipmentType: &EquipmentType{Id: -1, Name: "Shoes"}, Name: "name", Brand: "brand", Model: "model", Cost: 150, Size: "14", PurchaseDate: "2020-01-01", Notes: "", Mileage: measurement.CreateMeasurement(measurement.Mile, 10)}, true},
		{"invalid mileage", fields{Id: 1, UserUuid: "c2cb0cd2-cac8-11ee-b5e8-325096b39f47", EquipmentType: &EquipmentType{Id: 1, Name: "Shoes"}, Name: "name", Brand: "brand", Model: "model", Cost: 150, Size: "14", PurchaseDate: "2020-01-01", Notes: "", Mileage: measurement.CreateMeasurement(measurement.Mile, -1)}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Equipment{
				Id:            tt.fields.Id,
				UserUuid:      tt.fields.UserUuid,
				EquipmentType: tt.fields.EquipmentType,
				Name:          tt.fields.Name,
				Brand:         tt.fields.Brand,
				Model:         tt.fields.Model,
				Cost:          tt.fields.Cost,
				Size:          tt.fields.Size,
				PurchaseDate:  tt.fields.PurchaseDate,
				Notes:         tt.fields.Notes,
				Mileage:       tt.fields.Mileage,
			}
			if err := e.Validate(); (err != nil) != tt.wantErr {
			t.Errorf("Equipment.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestActivityEquipment_Validate(t *testing.T) {
	type fields struct {
		Id              int
		ActivityUuid    string
		Equipment       *Equipment
		AssignedMileage *measurement.Measurement
	}
	type args struct {
		insert bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"valid1", fields{Id: 1, ActivityUuid: "c2cb13d0-cac8-11ee-b7fb-325096b39f47", Equipment: &Equipment{Id: 1}, AssignedMileage: measurement.CreateMeasurement(measurement.Mile, 0)}, args{false}, false},
		{"valid2", fields{Id: 1, ActivityUuid: "c2cb13d0-cac8-11ee-b7fb-325096b39f47", Equipment: &Equipment{Id: 1}, AssignedMileage: measurement.CreateMeasurement(measurement.Mile, 10)}, args{false}, false},
		{"valid3", fields{Id: 1, ActivityUuid: "c2cb13d0-cac8-11ee-b7fb-325096b39f47", Equipment: &Equipment{Id: 1}, AssignedMileage: measurement.CreateMeasurement(measurement.Mile, 10)}, args{true}, false},
		{"invalid id", fields{Id: -1, ActivityUuid: "c2cb13d0-cac8-11ee-b7fb-325096b39f47", Equipment: &Equipment{Id: 1}, AssignedMileage: measurement.CreateMeasurement(measurement.Mile, 10)}, args{false}, true},
		{"invalid activity uuid", fields{Id: -1, ActivityUuid: "", Equipment: &Equipment{Id: 1}, AssignedMileage: measurement.CreateMeasurement(measurement.Mile, 10)}, args{false}, true},
		{"invalid equipment id", fields{Id: -1, ActivityUuid: "c2cb13d0-cac8-11ee-b7fb-325096b39f47", Equipment: &Equipment{Id: -1}, AssignedMileage: measurement.CreateMeasurement(measurement.Mile, 10)}, args{false}, true},
		{"invalid equipment assigned mileage", fields{Id: -1, ActivityUuid: "c2cb13d0-cac8-11ee-b7fb-325096b39f47", Equipment: &Equipment{Id: -1}, AssignedMileage: measurement.CreateMeasurement(measurement.Mile, -5)}, args{false}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ae := &ActivityEquipment{
				Id:              tt.fields.Id,
				ActivityUuid:    tt.fields.ActivityUuid,
				Equipment:       tt.fields.Equipment,
				AssignedMileage: tt.fields.AssignedMileage,
			}
			if err := ae.Validate(tt.args.insert); (err != nil) != tt.wantErr {
				t.Errorf("ActivityEquipment.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
