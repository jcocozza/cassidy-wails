package model

import (
	"fmt"
	"time"

	colorutil "github.com/jcocozza/cassidy-wails/internal/utils/colorUtil"
	"github.com/jcocozza/cassidy-wails/internal/utils/dateutil"
	"github.com/jcocozza/cassidy-wails/internal/utils/measurement"
)

type Activity struct {
	Uuid            string                 `json:"uuid"`
	Date            time.Time              `json:"date" ts_type:"Date" ts_transform:"new Date(__VALUE__)"`
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
	Map				string				   `json:"map"`
}
// An empty activity has no uuid, date, order of -1, no name, no description, no notes, empty type, empty typesubtype list, empty equipment list, empty planned, empty completed
func EmptyActivity() *Activity {
	return &Activity{
		Uuid:            "",
		Date:            time.Now(),
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
		Map: 			 "",
	}
}
// Validate an activity
//
// An activity is valid if it has a uuid, date, non-negative order, Non-empty type and valid, planned and completed.
func (a *Activity) Validate() error {
	if a.Uuid == "" {
		return fmt.Errorf("activity uuid is invalid")
	}
	/*
	if a.Date == "" {
		return fmt.Errorf("activity date is invalid")
	}
	_, err := time.Parse(dateutil.Layout, a.Date)
	if err != nil {
		return fmt.Errorf("activity date is invalid: %w", err)
	}
	*/
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
	userUnits := measurement.UnitClassControl(userUnitClass, measurement.Distance)
	plannedPaceUnit := measurement.PaceUnitByDistanceUnit(a.Type.Id, userUnits)
	completedPaceUnit := measurement.PaceUnitByDistanceUnit(a.Type.Id, userUnits)
	a.Planned.CalculatePace(plannedPaceUnit)
	a.Completed.CalculatePace(completedPaceUnit)
}
// return true if an activity is for a date in the future (i.e. it is planned)
func (a *Activity) IsFuture() bool {
	return a.Date.After(time.Now())
}
// Determine completion color of activity
//
// Based on duration, then distance
func (a *Activity) CompletionColor() (colorutil.Color, error) {
	isFuture := a.IsFuture()

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
// Logic for merging activities
//
// Check if the passed activity can be merged into the activity
//
// These checks should be redundant with proper database querying.
// When merging activities, one can just query the database based on similar conditions.
//
// Check for the following:
//  1. same date
//  2. no completed data
//  3. same activity type
func (a *Activity) CanMerge(activity *Activity) bool {
	// dates need to be the same
	y1, m1, d1 := a.Date.Date()
	y2, m2, d2 := activity.Date.Date()
	if y1 != y2 || m1 != m2 || d1 != d2 {
		return false
	}
	// shouldn't merge if there is already completed data
	if !a.Completed.IsZero() {
		return false
	}
	// only merge activities of the same type
	if a.Type.Id != activity.Type.Id {
		return false
	}
	return true
}
// Merge the passed activity into the activity
//
// This is typically used on import from external sources like Strava
func (a *Activity) Merge(activity *Activity) {
	a.Completed = activity.Completed

	// A merge will not overwrite existing fields unless those fields are already empty
	if a.Description == "" {
		a.Description = activity.Description
	}
	if a.Name == "" {
		a.Name = activity.Name
	}
	if a.Notes == "" {
		a.Notes = activity.Notes
	}
    if a.Map == "" {
        a.Map = activity.Map
    }
	// the new activity uuid will override the old one
	// hopefully this will eventually facilitate easier look back to data sources
    a.SetUuid(activity.Uuid)
}
// Represents a list of activities on a given day
type ActivityList struct {
	Date         time.Time   `json:"date" ts_type:"Date" ts_transform:"new Date(__VALUE__)"`
	ActivityList []*Activity `json:"activity_list"`
}
// An empty activity list has an empty date object and no activities
func EmptyActivityList() *ActivityList {
	return &ActivityList{
		Date: time.Now(),
		ActivityList: []*Activity{},
	}
}
// An empty dated activity list is an activity list with a date, but no activities in its list.
func EmptyDatedActivityList(date time.Time) *ActivityList {
	return &ActivityList{
		Date: date,
		ActivityList: []*Activity{},
	}
}
// Create a new activity list object for the passed date with an empty activity list
func NewActivityList(date time.Time) *ActivityList {
	return &ActivityList{
		Date: date,
		ActivityList: []*Activity{},
	}
}
// add an activity to the activity list
func (al *ActivityList) AddActivity(act *Activity) {
	al.ActivityList = append(al.ActivityList, act)
}

// A cycle is a list of activity lists.
type Cycle []*ActivityList

// Creata a new cycle based on start and end dates
func NewCycle(startDate, endDate time.Time) (*Cycle) {
	dates := dateutil.GenerateDateRange(startDate, endDate)
	cycle := Cycle{}
	for _, date := range dates {
		actList := EmptyDatedActivityList(date)
		cycle = append(cycle, actList)
	}
	return &cycle
}
// Add an activity to a cycle at the proper date.
func (c *Cycle) AddActivity(act *Activity) error {
	for _, actListOb := range *c {
		l := len(actListOb.ActivityList)
		if dateutil.SameDate(actListOb.Date, act.Date) {
			actListOb.AddActivity(act)
			if len(actListOb.ActivityList) != 1+l {
				return fmt.Errorf("failed to add activity to cycle")
			} else {
				return nil
			}
		}
	}
	// if we don't add to existing then we need to create a new activity list
	newAl := NewActivityList(act.Date)
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
