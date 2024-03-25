package controllers

import (
	"fmt"

	"github.com/jcocozza/cassidy-wails/internal/database"
	"github.com/jcocozza/cassidy-wails/internal/model"
	conversionrepo "github.com/jcocozza/cassidy-wails/internal/repository/conversionRepo"
	equipmentrepo "github.com/jcocozza/cassidy-wails/internal/repository/equipmentRepo"
)

type EquipmentHandler struct {
	EquipmentRepository  equipmentrepo.EquipmentRepository
	ConversionRepository conversionrepo.MeasurementRepository
}
func NewEquipmentHandler(db database.DbOperations) *EquipmentHandler {
	return &EquipmentHandler{
		EquipmentRepository:  equipmentrepo.NewIEquipmentRepository(db),
		ConversionRepository: conversionrepo.NewIMeasurementRepository(),
	}
}
// Create equipment
func (eh *EquipmentHandler) CreateEquipment(user *model.User, createRequest *model.Equipment) (*model.Equipment, error) {
	err0 := eh.ConversionRepository.ConvertEquipment(conversionrepo.Incoming, createRequest)
	if err0 != nil {
		return nil, fmt.Errorf("equipment failed to convert incoming: %w", err0)
	}

	createRequest.UserUuid = user.Uuid
	err1 := createRequest.Validate()
	if err1 != nil {
		return nil, err1
	}
	id, err2 := eh.EquipmentRepository.Create(createRequest)
	if err2 != nil {
		return nil, err2
	}
	createRequest.SetId(id)

	err3 := eh.ConversionRepository.ConvertEquipment(conversionrepo.Outgoing, createRequest)
	if err3 != nil {
		return nil, fmt.Errorf("equipment failed to convert outgoing: %w", err3)
	}
	return createRequest, nil
}
// Update Equipment
//
// @param: id
func (eh *EquipmentHandler) UpdateEquipment(equipmentId int, updateRequest *model.Equipment) (*model.Equipment, error) {
	err0 := eh.ConversionRepository.ConvertEquipment(conversionrepo.Incoming, updateRequest)
	if err0 != nil {
		return nil, err0
	}

	err1 := eh.EquipmentRepository.Update(updateRequest)
	if err1 != nil {
		return nil, err1
	}

	err2 := eh.ConversionRepository.ConvertEquipment(conversionrepo.Outgoing, updateRequest)
	if err2 != nil {
		return nil, err2
	}
	return updateRequest, nil
}
// Delete Equipment
//
// @param: id
func (eh *EquipmentHandler) DeleteEquipment(equipmentId int) error {
	err1 := eh.EquipmentRepository.Delete(equipmentId)
	if err1 != nil {
		return fmt.Errorf("failed to delete equipment: %w", err1)
	}
	return nil
}
// List all equipment
func (eh *EquipmentHandler) List(user *model.User) ([]*model.Equipment, error) {
	equipmentList, err := eh.EquipmentRepository.List(user.Uuid)
	if err != nil {
		return nil, fmt.Errorf("failed to get equipment list: %w", err)
	}
	for _, equipment := range equipmentList {
		err0 := eh.ConversionRepository.ConvertEquipment(conversionrepo.Outgoing, equipment)
		if err0 != nil {
			return nil, fmt.Errorf("equipment failed to convert: %w", err0)
		}
	}
	return equipmentList, nil
}
// List equipment types
func (eh *EquipmentHandler) ListEquipmentTypes() ([]*model.EquipmentType, error) {
	equipmentTypesList, err := eh.EquipmentRepository.ListEquipmentTypes()
	if err != nil {
		return nil, err
	}

	return equipmentTypesList, nil
}
// Create activity equipment
func (eh *EquipmentHandler) CreateActivityEquipment(createRequest *model.ActivityEquipment) (*model.ActivityEquipment, error) {
	err0 := eh.ConversionRepository.ConvertActivityEquipment(conversionrepo.Incoming, createRequest)
	if err0 != nil {
		return nil, err0
	}

	id, err := eh.EquipmentRepository.CreateActivityEquipment(createRequest)
	if err != nil {
		return nil, err
	}
	createRequest.SetId(id)

	err3 := eh.ConversionRepository.ConvertActivityEquipment(conversionrepo.Outgoing, createRequest)
	if err3 != nil {
		return nil, err3
	}
	return createRequest, nil
}

// Update Activity Equipment
// Update the mileage of an activity equipment (a patch for anything else in that doesn't make sense)
//
// @param: id
func (eh *EquipmentHandler) UpdateActivityEquipment(activityEquipmentId int, updateRequest *model.ActivityEquipment) (*model.ActivityEquipment, error) {
	err0 := eh.ConversionRepository.ConvertActivityEquipment(conversionrepo.Incoming, updateRequest)
	if err0 != nil {
		return nil, err0
	}

	err1 := eh.EquipmentRepository.UpdateActivityEquipment(updateRequest)
	if err1 != nil {
		return nil, err1
	}

	err2 := eh.ConversionRepository.ConvertActivityEquipment(conversionrepo.Outgoing, updateRequest)
	if err2 != nil {
		return nil, err2
	}
	return updateRequest, nil
}

// Delete Activity Equipment
//
// @param: id
func (eh *EquipmentHandler) DeleteActivityEquipment(activityEquipmentId int) error {
	err1 := eh.EquipmentRepository.DeleteActivityEquipment(activityEquipmentId)
	if err1 != nil {
		return err1
	}
	return nil
}
