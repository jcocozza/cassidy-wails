package model

import "fmt"

// TODO: These need to be programatically inserted into the database on build
var (
	Run = ActivityType{Id: 1, Name: "Run"}
		Run_Long = ActivitySubtype{SuperTypeId: Run.Id, Name: "Long"}
		Run_Fartlek = ActivitySubtype{SuperTypeId: Run.Id, Name: "Fartlek"}
		Run_Tempo = ActivitySubtype{SuperTypeId: Run.Id, Name: "Tempo"}
		Run_Track = ActivitySubtype{SuperTypeId: Run.Id, Name: "Track"}
		Run_Intervals = ActivitySubtype{SuperTypeId: Run.Id, Name: "Intervals"}
		Run_Recovery = ActivitySubtype{SuperTypeId: Run.Id, Name: "Recovery"}
		Run_Indoor = ActivitySubtype{SuperTypeId: Run.Id, Name: "Indoor"}
		Run_Trails = ActivitySubtype{SuperTypeId: Run.Id, Name: "Trails"}
	Bike = ActivityType{Id: 2, Name: "Bike"}
		Bike_Long = ActivitySubtype{SuperTypeId: Bike.Id, Name: "Long"}
		Bike_Velodrome = ActivitySubtype{SuperTypeId: Bike.Id, Name: "Velodrome"}
		Bike_Recovery = ActivitySubtype{SuperTypeId: Bike.Id, Name: "Recovery"}
		Bike_Indoor = ActivitySubtype{SuperTypeId: Bike.Id, Name: "Indoor"}
	Swim = ActivityType{Id: 3, Name: "Swim"}
		Swim_Drills = ActivitySubtype{SuperTypeId: Swim.Id, Name: "Drills"}
		Swim_OpenWater = ActivitySubtype{SuperTypeId: Swim.Id, Name: "Open Water"}
		Swim_Recovery = ActivitySubtype{SuperTypeId: Swim.Id, Name: "Recovery"}
	Hike = ActivityType{Id: 4, Name: "Hike"}
	RestDay = ActivityType{Id: 5, Name: "Rest Day"}
	Strength = ActivityType{Id: 6, Name: "Strength"}
	Other = ActivityType{Id: 7, Name: "Other"}
	MountainBike = ActivityType{Id: 8, Name: "Mountain Bike"}
)

// Represents an activity type.
//
// e.g. Run, Bike, etc.
type ActivityType struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
// An empty activity type has an id of -1 and an empty name.
func EmptyActivityType() *ActivityType {
	return &ActivityType{
		Id:   -1,
		Name: "",
	}
}

// Validate an activity type object.
//
// An activity type is valid when it has a positive id and non empty name.
func (at *ActivityType) Validate() error {
	if at.Id == -1 {
		return fmt.Errorf("activity type id cannot be -1")
	}
	if at.Name == "" {
		return fmt.Errorf("activity type name cannot be empty")
	}
	return nil
}

// Represents an activity subtype.
//
// e.g. Long, Easy, Fartlek, etc.
type ActivitySubtype struct {
	Id          int    `json:"id"`
	SuperTypeId int    `json:"supertype_id"`
	Name        string `json:"name"`
}

// An empty activity subtype has an id of -1, a super type id of -1 and an empty name.
func EmptyActivitySubtype() *ActivitySubtype {
	return &ActivitySubtype{
		Id:          -1,
		SuperTypeId: -1,
		Name:        "",
	}
}

// Validate an activity subtype object
//
// An activity subtype is valid when it has a postivive id, positive supertype id, and non-empty name.
func (as *ActivitySubtype) Validate() error {
	if as.Id == -1 {
		return fmt.Errorf("activity subtype id cannot be -1")
	}
	if as.SuperTypeId == -1 {
		return fmt.Errorf("activity subtype supertype id cannot be -1")
	}
	if as.Name == "" {
		return fmt.Errorf("activity subtype name cannot be empty")
	}
	return nil
}

// Represents a type-subtype pair for an activity
//
// e.g. Run: Long or Bike: Velodrome
type ActivityTypeSubtype struct {
	Id              int              `json:"id"`
	ActivityUuid    string           `json:"activity_uuid"`
	ActivityType    *ActivityType    `json:"activity_type"`
	ActivitySubtype *ActivitySubtype `json:"activity_subtype"`
}
func NewActivityTypeSubtype(activityUuid string, activityType *ActivityType, activitySubtype *ActivitySubtype) *ActivityTypeSubtype {
	return &ActivityTypeSubtype{
		ActivityUuid: activityUuid,
		ActivityType: activityType,
		ActivitySubtype: activitySubtype,
	}
}
// An empty activity type subtype has an id of -1, no activity uuid, an empty activity type and and empty activity subtype
func EmptyActivityTypeSubtype() *ActivityTypeSubtype {
	return &ActivityTypeSubtype{
		Id:              -1,
		ActivityUuid:    "",
		ActivityType:    EmptyActivityType(),
		ActivitySubtype: EmptyActivitySubtype(),
	}
}
func (ats *ActivityTypeSubtype) Validate() error {
	if ats.Id == -1 {
		return fmt.Errorf("activity subtype id cannot be -1")
	}
	if ats.ActivityUuid == "" {
		return fmt.Errorf("activity subtype supertype id cannot be -1")
	}
	err1 := ats.ActivityType.Validate()
	if err1 != nil {
		return fmt.Errorf("activity type subtype activity type is invalid: %w", err1)
	}
	err2 := ats.ActivitySubtype.Validate()
	if err2 != nil {
		return fmt.Errorf("activity type subtype activity subtype is invalid: %w", err2)
	}
	return nil
}
func (ats *ActivityTypeSubtype) SetId(id int) {
	ats.Id = id
}

// Is an activity type with a list of its corresponding subtypes
type ActivityTypeWithSubtypes struct {
	ActivityType *ActivityType      `json:"activity_type"`
	SubtypeList  []*ActivitySubtype `json:"subtype_list"`
}

// Add a subtype to the list of subtypes
func (atws *ActivityTypeWithSubtypes) AddSubtype(sub *ActivitySubtype) {
	atws.SubtypeList = append(atws.SubtypeList, sub)
}