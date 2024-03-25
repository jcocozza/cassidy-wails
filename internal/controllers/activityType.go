package controllers

import (
	"fmt"

	"github.com/jcocozza/cassidy-wails/internal/database"
	"github.com/jcocozza/cassidy-wails/internal/model"
	activitytyperepo "github.com/jcocozza/cassidy-wails/internal/repository/activityTypeRepo"
)

type ActivityTypeHandler struct {
	ActivityTypeRepository activitytyperepo.ActivityTypeRepository
}

func NewActivityTypeHandler(db database.DbOperations) *ActivityTypeHandler {
	return &ActivityTypeHandler{
		ActivityTypeRepository: activitytyperepo.NewIActivityTypeRepository(db),
	}
}

// list all activity types
func (ath *ActivityTypeHandler) ListActivityTypes() ([]*model.ActivityTypeWithSubtypes, error) {
	activityTypeWithSubtypeList, err := ath.ActivityTypeRepository.List()

	if err != nil {
		return nil, fmt.Errorf("something went wrong creating activity type subtype list: %w", err)
	}
	return activityTypeWithSubtypeList, nil
}
