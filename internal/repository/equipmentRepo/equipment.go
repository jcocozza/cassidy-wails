package equipmentrepo

import (
	"fmt"

	"github.com/jcocozza/cassidy-wails/internal/database"
	"github.com/jcocozza/cassidy-wails/internal/model"
	"github.com/jcocozza/cassidy-wails/internal/sqlcode"
	"github.com/jcocozza/cassidy-wails/internal/utils"
)

// The methods for interacting with equipment objects
type EquipmentRepository interface {
	Create(equipment *model.Equipment) (int, error)
	Update(equipment *model.Equipment) error
	Delete(id int) error
	List(userUuid string) ([]*model.Equipment, error)

	// Equipment Types
	ListEquipmentTypes() ([]*model.EquipmentType, error)

	// Activity Equipment
	CreateActivityEquipment(activityEquipment *model.ActivityEquipment) (int, error)
	UpdateActivityEquipment(activityEquipment *model.ActivityEquipment) error
	DeleteActivityEquipment(id int) error
}

// Represent a SQLite database connection
type IEquipmentRepository struct {
	DB database.DbOperations
}

func NewIEquipmentRepository(db database.DbOperations) *IEquipmentRepository {
	return &IEquipmentRepository{
		DB: db,
	}
}

// Insert the equipment object into the database.
func (db *IEquipmentRepository) Create(equipment *model.Equipment) (int, error) {
	sql := utils.SQLReader(sqlcode.Equipment_create)

	err1 := equipment.Validate()
	if err1 != nil {
		return -1, fmt.Errorf("equipment creation failed to validate: %w", err1)
	}

	id, err := db.DB.ExecuteGetLast(sql, equipment.EquipmentType.Id, equipment.UserUuid, equipment.Name, equipment.Brand, equipment.Model, equipment.Cost, equipment.Size, equipment.PurchaseDate, equipment.Name, equipment.Mileage.Length, equipment.Mileage.Unit, equipment.IsRetired)
	if err != nil {
		return -1, fmt.Errorf("equipment creation failed: %w", err)
	}

	return id, nil
}

// Update the equipment object in the database
func (db *IEquipmentRepository) Update(equipment *model.Equipment) error {
	err1 := equipment.Validate()
	if err1 != nil {
		return fmt.Errorf("equipment update failed to validate: %w", err1)
	}

	sql := utils.SQLReader(sqlcode.Equipment_update)
	err := db.DB.Execute(sql, equipment.Name, equipment.Brand, equipment.Model, equipment.Cost, equipment.Size, equipment.PurchaseDate, equipment.Notes, equipment.Mileage.Length, equipment.Mileage.Unit, equipment.IsRetired, equipment.Id)
	if err != nil {
		return fmt.Errorf("equipment update failed: %w", err)
	}
	return nil
}

// Delete the equipment object in the database
func (db *IEquipmentRepository) Delete(id int) error {
	sql := utils.SQLReader(sqlcode.Equipment_delete)
	err := db.DB.Execute(sql, id)
	if err != nil {
		return fmt.Errorf("equipment delete failed: %w", err)
	}
	return nil
}

// List all equipment for a given user uuid
func (db *IEquipmentRepository) List(userUuid string) ([]*model.Equipment, error) {
	sql := utils.SQLReader(sqlcode.Equipment_list)
	rows, err := db.DB.Query(sql, userUuid)
	if err != nil {
		return nil, fmt.Errorf("error querying equipment list: %w", err)
	}
	defer rows.Close()
	equipmentList := []*model.Equipment{}
	for rows.Next() {
		equipment := model.EmptyEquipment()
		err1 := rows.Scan(&equipment.Id,
			&equipment.EquipmentType.Id,
			&equipment.EquipmentType.Name,
			&equipment.Name,
			&equipment.Brand,
			&equipment.Model,
			&equipment.Cost,
			&equipment.Size,
			&equipment.PurchaseDate,
			&equipment.Notes,
			&equipment.Mileage.Length,
			&equipment.Mileage.Unit,
			&equipment.IsRetired)

		if err1 != nil {
			return nil, fmt.Errorf("error scanning equipment list: %w", err1)
		}
		equipmentList = append(equipmentList, equipment)
	}
	return equipmentList, nil
}

// List all equipment types in the database
//
// Return an empty or possibly incomplete list if an error is thrown
func (db *IEquipmentRepository) ListEquipmentTypes() ([]*model.EquipmentType, error) {
	equipmentTypeList := []*model.EquipmentType{}

	sql := utils.SQLReader(sqlcode.EquipmentType_list)
	rows, err := db.DB.Query(sql)
	if err != nil {
		return equipmentTypeList, fmt.Errorf("failed to get equipment type list from database: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		equipmentType := model.EmptyEquipmentType()
		err := rows.Scan(&equipmentType.Id, &equipmentType.Name)
		if err != nil {
			return equipmentTypeList, fmt.Errorf("error scanning equipment type: %w", err)
		}

		err2 := equipmentType.Validate()
		if err2 != nil {
			return equipmentTypeList, fmt.Errorf("equipment type failed to validate while scanning: %w", err)
		}

		equipmentTypeList = append(equipmentTypeList, equipmentType)
	}
	return equipmentTypeList, nil
}

// Insert the activity equipment object into the database.
//
// Returns the id of the inserted row or -1 if the insert fails.
func (db *IEquipmentRepository) CreateActivityEquipment(activityEquipment *model.ActivityEquipment) (int, error) {
	sql := utils.SQLReader(sqlcode.ActivityEquipment_create)

	err := activityEquipment.Validate(true)
	if err != nil {
		return -1, fmt.Errorf("activity equipment creation failed to validate: %w", err)
	}

	id, err := db.DB.ExecuteGetLast(sql, activityEquipment.ActivityUuid, activityEquipment.Equipment.Id, activityEquipment.AssignedMileage.Length, activityEquipment.AssignedMileage.Unit)

	if err != nil || id == -1 {
		return -1, fmt.Errorf("failed to create activity equipment: %w", err)
	}
	return id, nil
}

// Update activity equipment object in the database.
//
// This really amounts to setting the assigned mileage.
func (db *IEquipmentRepository) UpdateActivityEquipment(activityEquipment *model.ActivityEquipment) error {
	err1 := activityEquipment.Validate(false)
	if err1 != nil {
		return fmt.Errorf("activity equipment update failed to validate: %w", err1)
	}

	sql := utils.SQLReader(sqlcode.ActivityEquipment_update)
	err := db.DB.Execute(sql, activityEquipment.AssignedMileage.Length, activityEquipment.AssignedMileage.Unit, activityEquipment.Id)
	if err != nil {
		return fmt.Errorf("activity equipment update failed: %w", err)
	}
	return nil
}

// Delete activity equipment object in the database.
func (db *IEquipmentRepository) DeleteActivityEquipment(id int) error {
	sql := utils.SQLReader(sqlcode.ActivityEquipment_delete)
	err := db.DB.Execute(sql, id)
	if err != nil {
		return fmt.Errorf("activity equipment delete failed: %w", err)
	}
	return nil
}
