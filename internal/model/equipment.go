package model

import (
	"fmt"

	"github.com/jcocozza/cassidy-wails/internal/utils/measurement"
)

// Represents a type of equipment.
//
// e.g. Shoes, Bike, etc
type EquipmentType struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// An empty equipment type has no name and an id of -1.
func EmptyEquipmentType() *EquipmentType {
	return &EquipmentType{
		Id:   -1,
		Name: "",
	}
}

// Validate an equipment type.
//
// Validation ensures that the id is not -1 and the name is not empty
func (et *EquipmentType) Validate() error {
	if et.Id == -1 {
		return fmt.Errorf("equipment type id is invalid")
	}
	if et.Name == "" {
		return fmt.Errorf("equipment type name is invalid")
	}
	return nil
}

// Represents a specific piece of equipment.
//
// e.g. My running shoes that have 212 miles on them.
type Equipment struct {
	Id            int                      `json:"id"`
	UserUuid      string                   `json:"user_uuid"`
	EquipmentType *EquipmentType           `json:"equipment_type"`
	Name          string                   `json:"name"`
	Brand         string                   `json:"brand"`
	Model         string                   `json:"model"`
	Cost          float64                  `json:"cost"`
	Size          string                   `json:"size"`
	PurchaseDate  string                   `json:"purchase_date"`
	Notes         string                   `json:"notes"`
	Mileage       *measurement.Measurement `json:"mileage"`
	IsRetired     bool                     `json:"is_retired"`
}

// An empty equipment has an id of -1, not user uuid, an empty equipment type,
// empty name, empty brand, empty model, cost of 0, size of -1, empty purchase date, empty notes, default mileage(0 meters).
func EmptyEquipment() *Equipment {
	return &Equipment{
		Id:            -1,
		UserUuid:      "",
		EquipmentType: EmptyEquipmentType(),
		Name:          "",
		Brand:         "",
		Model:         "",
		Cost:          0,
		Size:          "",
		PurchaseDate:  "",
		Notes:         "",
		Mileage:       measurement.EmptyMeasurement(),
	}
}

// Validate a equipment object
//
// Validation ensures that user uuid, equipment type id and mileage are valid
func (e *Equipment) Validate() error {
	//if e.UserUuid == "" {
	//	return fmt.Errorf("equipment user uuid is invalid")
	//}
	if e.EquipmentType.Id < 1 {
		return fmt.Errorf("equipment equipment type id is invalid")
	}
	if e.Mileage.Length < 0 {
		return fmt.Errorf("mileage cannot be negative")
	}
	return nil
}

// Set the id of an equipment
func (e *Equipment) SetId(id int) {
	e.Id = id
}

// Represents the assignment of a piece of equipment to an activity.
//
// e.g. On activity A, I ran 10 miles in piece of equipment X.
type ActivityEquipment struct {
	Id              int                      `json:"id"`
	ActivityUuid    string                   `json:"activity_uuid"`
	Equipment       *Equipment               `json:"equipment"`
	AssignedMileage *measurement.Measurement `json:"assigned_mileage"`
}

// An empty activity equipment has an id of -1, no activity uuid, an empty equipment, and an empty mileage.
func EmptyActivityEquipment() *ActivityEquipment {
	return &ActivityEquipment{
		Id:              -1,
		ActivityUuid:    "",
		Equipment:       EmptyEquipment(),
		AssignedMileage: measurement.EmptyMeasurement(),
	}
}

// Validate an activity equipment object.
//
// Validation ensures that the object has a uuid, equipment id and non-negative assigned mileage.
//
// If this is a first time insertion, ignore the Id.
func (ae *ActivityEquipment) Validate(insert bool) error {
	if !insert {
		if ae.Id == -1 {
			return fmt.Errorf("activity equipment id is invalid")
		}
	}
	if ae.ActivityUuid == "" {
		return fmt.Errorf("activity equipment activity uuid is invalid")
	}
	if ae.Equipment.Id == -1 {
		return fmt.Errorf("activity equipment equipment id is invalid")
	}
	if ae.AssignedMileage.Length < 0 {
		return fmt.Errorf("activity equipment assigned mileage id cannot be negative")
	}
	return nil
}

// Set the id of an activity equipment
func (ae *ActivityEquipment) SetId(id int) {
	ae.Id = id
}
