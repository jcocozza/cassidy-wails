package model

import (
	"fmt"
	"time"

	colorutil "github.com/jcocozza/cassidy-wails/internal/utils/colorUtil"
	"github.com/jcocozza/cassidy-wails/internal/utils/dateutil"
	"github.com/jcocozza/cassidy-wails/internal/utils/measurement"
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

type Activity struct {
	Uuid            string                 `json:"uuid"`
	Date            string                 `json:"date"`
	Order           int                    `json:"order"`
	Name            string                 `json:"name"`
	Description     string                 `json:"description"`
	Notes           string                 `json:"notes"`
	Type            *ActivityType          `json:"activity_type"`
	TypeSubtypeList []*ActivityTypeSubtype `json:"type_subtype_list"`
	EquipmentList   []*ActivityEquipment   `json:"equipment_list"`
	Planned         *Planned               `json:"planned"`
	Completed       *Completed             `json:"completed"`
	Color           colorutil.Color        `json:"color"`
	IsRace          bool                   `json:"is_race"`
	NumStrides      int                    `json:"num_strides"`
}

// An empty activity has no uuid, date, order of -1, no name, no description, no notes, empty type, empty typesubtype list, empty equipment list, empty planned, empty completed
func EmptyActivity() *Activity {
	return &Activity{
		Uuid:            "",
		Date:            "",
		Order:           -1,
		Name:            "",
		Description:     "",
		Notes:           "",
		Type:            EmptyActivityType(),
		TypeSubtypeList: []*ActivityTypeSubtype{},
		EquipmentList:   []*ActivityEquipment{},
		Planned:         EmptyPlanned(),
		Completed:       EmptyCompleted(),
		IsRace:          false,
		NumStrides:      0,
	}
}

// Validate an activity
//
// An activity is valid if it has a uuid, date, non-negative order, Non-empty type and valid, planned and completed.
func (a *Activity) Validate() error {
	if a.Uuid == "" {
		return fmt.Errorf("activity uuid is invalid")
	}
	if a.Date == "" {
		return fmt.Errorf("activity date is invalid")
	}
	_, err := time.Parse(dateutil.Layout, a.Date)
	if err != nil {
		return fmt.Errorf("activity date is invalid: %w", err)
	}
	if a.Order < 1 {
		return fmt.Errorf("activity order is invalid")
	}
	err1 := a.Type.Validate()
	if err1 != nil {
		return fmt.Errorf("activity type is invalid: %w", err1)
	}
	err2 := a.Planned.Validate()
	if err2 != nil {
		return fmt.Errorf("activity planned is invalid: %w", err2)
	}
	err3 := a.Completed.Validate()
	if err3 != nil {
		return fmt.Errorf("activity completed is invalid: %w", err3)
	}
	return nil
}
// Add activity TypeSubtype object to the activity object
func (a *Activity) AddActivityTypeSubtype(ats *ActivityTypeSubtype) {
	a.TypeSubtypeList = append(a.TypeSubtypeList, ats)
}
// add activity equipment object to the activity object
func (a *Activity) AddActivityEquipment(e *ActivityEquipment) {
	a.EquipmentList = append(a.EquipmentList, e)
}
// Set the activity uuid of sub elements
func (a *Activity) SetUuid(uuid string) {
	a.Uuid = uuid
	a.Planned.ActivityUuid = uuid
	a.Completed.ActivityUuid = uuid

	for _, activityEquipment := range a.EquipmentList {
		activityEquipment.ActivityUuid = uuid
	}
	for _, ats := range a.TypeSubtypeList {
		ats.ActivityUuid = uuid
	}
}
// Calculate pace for planned and completed
func (a *Activity) CalculatePace(userUnitClass measurement.UnitClass) {
	//paceUnit := measurement.PaceUnitByActivityType(a.Type.Id, userUnitClass)
	plannedPaceUnit := measurement.PaceUnitByDistanceUnit(a.Type.Id, a.Planned.Distance.Unit)
	completedPaceUnit := measurement.PaceUnitByDistanceUnit(a.Type.Id, a.Completed.Distance.Unit)
	a.Planned.CalculatePace(plannedPaceUnit)
	a.Completed.CalculatePace(completedPaceUnit)
}
// return true if an activity is for a date in the future (i.e. it is planned)
func (a *Activity) IsFuture() (bool, error) {
	f, err := dateutil.IsFuture(a.Date)
	if err != nil {
		return true, err // in this context, it makes sense to set the is_future to true b/c of the coloring
	}

	return f, nil
}
// Determine completion color of activity
//
// Based on duration, then distance
func (a *Activity) CompletionColor() (colorutil.Color, error) {
	isFuture, err := a.IsFuture()
	if err != nil {
		return colorutil.Grey, err
	}

	if isFuture {
		return colorutil.Grey, nil
	}
	if a.Planned.IsZero() {
		return colorutil.Blue, nil
	}

	if a.Planned.Duration != 0 {
		if (a.Completed.Duration >= a.Planned.Duration*.8) && (a.Completed.Duration <= a.Planned.Duration*1.2) {
			return colorutil.Green, nil
		} else if ((a.Completed.Duration < a.Planned.Duration*.8) && (a.Completed.Duration >= a.Planned.Duration*.5)) || ((a.Completed.Duration > a.Planned.Duration*1.2) && (a.Completed.Duration <= a.Planned.Duration*1.5)) {
			return colorutil.Yellow, nil
		} else {
			return colorutil.Red, nil
		}
	}

	if a.Planned.Distance.Length != 0 {
		if (a.Completed.Distance.Length >= a.Planned.Distance.Length*.8) && (a.Completed.Distance.Length <= a.Planned.Distance.Length*1.2) {
			return colorutil.Green, nil
		} else if ((a.Completed.Distance.Length < a.Planned.Distance.Length*.8) && (a.Completed.Distance.Length >= a.Planned.Distance.Length*.5)) || ((a.Completed.Distance.Length > a.Planned.Distance.Length*1.2) && (a.Completed.Distance.Length <= a.Planned.Distance.Length*1.5)) {
			return colorutil.Yellow, nil
		} else {
			return colorutil.Red, nil
		}
	}
	return colorutil.Grey, nil // default color is grey
}
// Represents a list of activities on a given day
type ActivityList struct {
	DateObject   *dateutil.DateObject `json:"date_object"`
	ActivityList []*Activity          `json:"activity_list"`
}
// An empty activity list has an empty date object and no activities
func EmptyActivityList() *ActivityList {
	return &ActivityList{
		DateObject:   dateutil.EmptyDateObject(),
		ActivityList: []*Activity{},
	}
}
// An empty dated activity list is an activity list with a date, but no activities in its list.
func EmptyDatedActivityList(date string) (*ActivityList, error) {
	d, err := dateutil.CreateFromDate(date)
	if err != nil {
		return nil, err
	}
	return &ActivityList{
		DateObject:   d,
		ActivityList: []*Activity{},
	}, nil
}
// Create a new activity list object for the passed date with an empty activity list
func NewActivityList(date string) (*ActivityList, error) {
	do, err := dateutil.CreateFromDate(date)
	if err != nil {
		return nil, fmt.Errorf("unable to create activity list: %w", err)
	}
	return &ActivityList{
		DateObject:   do,
		ActivityList: []*Activity{},
	}, nil
}
// add an activity to the activity list
func (al *ActivityList) AddActivity(act *Activity) {
	al.ActivityList = append(al.ActivityList, act)
}

// A cycle is a list of activity lists.
type Cycle []*ActivityList
// Creata a new cycle based on start and end dates
func NewCycle(startDate, endDate string) (*Cycle, error) {
	dates, err := dateutil.GenerateDateRange(startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("failed to generate date range for cycle creation: %w", err)
	}

	cycle := Cycle{}
	for _, date := range dates {
		actList, err1 := EmptyDatedActivityList(date.Date)
		if err1 != nil {
			return nil, fmt.Errorf("failed to create activity list: %w", err1)
		}
		cycle = append(cycle, actList)
	}
	return &cycle, nil
}
// Add an activity to a cycle at the proper date.
func (c *Cycle) AddActivity(act *Activity) error {
	for _, actListOb := range *c {
		l := len(actListOb.ActivityList)
		if actListOb.DateObject.Date == act.Date {
			actListOb.AddActivity(act)
			if len(actListOb.ActivityList) != 1+l {
				return fmt.Errorf("failed to add activity to cycle")
			} else {
				return nil
			}
		}
	}
	// if we don't add to existing then we need to create a new activity list
	newAl, err := NewActivityList(act.Date)
	if err != nil {
		return fmt.Errorf("failed to add activity to cycle: %w", err)
	}
	*c = append(*c, newAl)
	return nil
}
// Return a list of uuids lists for each day in the cycle.
func (c *Cycle) CreateUuidLists() [][]string {
	uuidLists := [][]string{}
	for _, actListOb := range *c {
		uuidList := []string{}
		for _, act := range actListOb.ActivityList {
			uuidList = append(uuidList, act.Uuid)
		}
		uuidLists = append(uuidLists, uuidList)
	}
	return uuidLists
}
// Return a list of uuids for all activities in the cycle.
func (c *Cycle) CreateUuidList() []string {
	uuidList := []string{}
	for _, actListObj := range *c {
		for _, act := range actListObj.ActivityList {
			uuidList = append(uuidList, act.Uuid)
		}
	}
	return uuidList
}
// Calculate pace for each activity in the cycle
func (c *Cycle) CalculatePaces(userUnitClass measurement.UnitClass) {
	for _, actListObj := range *c {
		for _, act := range actListObj.ActivityList {
			act.CalculatePace(userUnitClass)
		}
	}
}
