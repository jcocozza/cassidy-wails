package controllers

import (
	"fmt"

	"github.com/jcocozza/cassidy-wails/internal/database"
	"github.com/jcocozza/cassidy-wails/internal/model"
	activityrepo "github.com/jcocozza/cassidy-wails/internal/repository/activityRepo"
	conversionrepo "github.com/jcocozza/cassidy-wails/internal/repository/conversionRepo"
	"github.com/jcocozza/cassidy-wails/internal/utils/uuidgen"
)

type ActivityHandler struct {
	ActivityRepository   activityrepo.ActivityRepository
	ConversionRepository conversionrepo.MeasurementRepository
	User *model.User
}
func NewActivityHandler(db database.DbOperations, user *model.User) *ActivityHandler {
	return &ActivityHandler{
		ActivityRepository:   activityrepo.NewIActivityRepository(db),
		ConversionRepository: conversionrepo.NewIMeasurementRepository(),
		User: user,
	}
}
// Create an activity
func (ah *ActivityHandler) CreateActivity(createRequest *model.Activity) (*model.Activity, error) {
	err0 := ah.ConversionRepository.ConvertActivity(conversionrepo.Incoming, createRequest, ah.User.Units)
	if err0 != nil {
		return nil, err0
	}

	newActUuid := uuidgen.GenerateUUID()
	createRequest.SetUuid(newActUuid)

	err := createRequest.Validate()
	if err != nil {
		return nil, fmt.Errorf("activity failed to validate: %w", err)
	}

	err2 := ah.ActivityRepository.Create(ah.User.Uuid, createRequest)
	if err2 != nil {
		return nil, fmt.Errorf("failed to create activity: %w", err2)
	} else {
		err3 := ah.ConversionRepository.ConvertActivity(conversionrepo.Outgoing, createRequest, ah.User.Units)
		if err3 != nil {
			return nil, fmt.Errorf("activity failed to convert: %w", err3)
		}
		return createRequest, nil
	}
}
// Update an activity
func (ah *ActivityHandler) UpdateActivity(updateRequest *model.Activity) (*model.Activity, error) {
	err0 := ah.ConversionRepository.ConvertActivity(conversionrepo.Incoming, updateRequest, ah.User.Units)
	if err0 != nil {
		return nil, fmt.Errorf("activity failed to convert incoming: %w", err0)
	}

	err := ah.ActivityRepository.Update(updateRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to update: %w", err)
	}
	// return updated object
	err3 := ah.ConversionRepository.ConvertActivity(conversionrepo.Outgoing, updateRequest, ah.User.Units)
	if err3 != nil {
		return nil, fmt.Errorf("activity failed to convert outgoing: %w", err3)
	}
	return updateRequest, nil
}
// Delete an activity
//
// @param: uuid
func (ah *ActivityHandler) DeleteActivity(activityUuid string) error {
	err := ah.ActivityRepository.Delete(activityUuid)
	if err != nil {
		return fmt.Errorf("failed to delete activity: %w", err)
	}
	return nil
}
