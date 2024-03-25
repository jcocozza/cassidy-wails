package activitytyperepo

import (
	"fmt"

	"github.com/jcocozza/cassidy-wails/internal/database"
	"github.com/jcocozza/cassidy-wails/internal/model"
	"github.com/jcocozza/cassidy-wails/internal/sqlcode"
	"github.com/jcocozza/cassidy-wails/internal/utils"
)

// Methods for Working with activity types
type ActivityTypeRepository interface {
	List() ([]*model.ActivityTypeWithSubtypes, error)
}

// Represents a database connection
type IActivityTypeRepository struct {
	DB database.DbOperations
}

func NewIActivityTypeRepository(db database.DbOperations) *IActivityTypeRepository {
	return &IActivityTypeRepository{
		DB: db,
	}
}

// Check if an activity type is in a list of activity types with subtypes.
func activityTypeInList(activityType *model.ActivityType, lst []*model.ActivityTypeWithSubtypes) bool {
	for _, elm := range lst {
		if elm.ActivityType.Id == activityType.Id {
			return true
		}
	}
	return false
}

// For a given activity type id, find the index for that activity type in the list of activity tytpe with subtypes.
func findATWS(actId int, lst []*model.ActivityTypeWithSubtypes) (int, error) {
	for i, atws := range lst {
		if atws.ActivityType.Id == actId {
			return i, nil
		}
	}
	return -1, fmt.Errorf("failed to find activity type with subtypes")
}

// List all activity types with their subtypes.
func (db *IActivityTypeRepository) List() ([]*model.ActivityTypeWithSubtypes, error) {
	sql := utils.SQLReader(sqlcode.ActivityType_list_with_subtype)

	rows, err := db.DB.Query(sql)
	if err != nil {
		return nil, fmt.Errorf("failed to get activity type with subtypes list: %w", err)
	}

	activityTypeWithSubtypeList := []*model.ActivityTypeWithSubtypes{}
	for rows.Next() {
		actType := model.EmptyActivityType()
		actSubtype := model.EmptyActivitySubtype()

		err2 := rows.Scan(&actType.Id, &actType.Name, &actSubtype.Id, &actSubtype.Name)
		actSubtype.SuperTypeId = actType.Id
		if err2 != nil {
			//Scan Error is okay here because some rows should be null
			// TODO: handle this properly
			// in the case of no subtypes
			actSubtype.Id = 0
			actSubtype.Name = ""
		}

		if !activityTypeInList(actType, activityTypeWithSubtypeList) {
			var newATWS *model.ActivityTypeWithSubtypes
			if actSubtype.Id != 0 {
				newATWS = &model.ActivityTypeWithSubtypes{ActivityType: actType, SubtypeList: []*model.ActivitySubtype{actSubtype}}
			} else {
				newATWS = &model.ActivityTypeWithSubtypes{ActivityType: actType, SubtypeList: []*model.ActivitySubtype{}}
			}
			activityTypeWithSubtypeList = append(activityTypeWithSubtypeList, newATWS)
		} else {
			if actSubtype.Id != 0 {
				idx, err3 := findATWS(actType.Id, activityTypeWithSubtypeList)

				if err3 != nil {
					return nil, fmt.Errorf("unexpected error creating activity type with subtypes list")
				}
				activityTypeWithSubtypeList[idx].AddSubtype(actSubtype)
			}
		}
	}
	return activityTypeWithSubtypeList, nil
}
