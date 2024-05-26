package strava

import (
	"context"
	"fmt"
	"time"

	"github.com/jcocozza/cassidy-connector/strava/app"
	"github.com/jcocozza/cassidy-connector/strava/swagger"
	"github.com/jcocozza/cassidy-wails/internal/controllers"
	"github.com/jcocozza/cassidy-wails/internal/model"
	"github.com/jcocozza/cassidy-wails/internal/utils/measurement"
	"golang.org/x/oauth2"
)

type Strava struct {
	App *app.App
	Handlers *controllers.Controllers
}
func NewStravaApp(handlers *controllers.Controllers) *Strava {
	return &Strava{
		App: app.NewApp(ClientId, ClientSecret, RedirectUri, Scopes),
		Handlers: handlers,
	}
}
// Get the token during the initial authentication process
func (s *Strava) StartListener() (*oauth2.Token, error) {
	// note this will also set the token in the app struct
	token, err := s.App.AwaitInitialToken(-1)
	if err != nil {
		return nil, err
	}
	return token, nil
}
// Open the authentication link in the user's browser
func (s *Strava) OpenStravaAuth() {
	s.App.OpenAuthorizationGrant()
}
// Open the strava app settings so user can revoke access. Then delete their strava token from database.
func (s *Strava) RevokeAccess(user *model.User) error {
	s.App.OpenStravaAppSettings()
	err := s.Handlers.UserHandler.DeleteStravaToken(user)
	if err != nil {
		return err
	}
	return nil
}
func (s *Strava) sportTypeToActivityType(activityId string, sportType swagger.SportType) (*model.ActivityType, []*model.ActivityTypeSubtype) {

	var actType *model.ActivityType
	actTypeLst := []*model.ActivityTypeSubtype{}

	switch sportType {
	case swagger.RUN_SportType:
		actType = &model.Run
	case swagger.TRAIL_RUN_SportType:
		actType = &model.Run
		actTypeLst = append(actTypeLst, model.NewActivityTypeSubtype(activityId, &model.Run, &model.Run_Trails))
	case swagger.VIRTUAL_RUN_SportType:
		actType = &model.Run
		actTypeLst = append(actTypeLst, model.NewActivityTypeSubtype(activityId, &model.Run, &model.Run_Indoor))
	case swagger.RIDE_SportType:
		actType = &model.Bike
	case swagger.VIRTUAL_RIDE_SportType:
		actType = &model.Bike
		actTypeLst = append(actTypeLst, model.NewActivityTypeSubtype(activityId, &model.Bike, &model.Bike_Indoor))
	case swagger.MOUNTAIN_BIKE_RIDE_SportType:
		actType = &model.MountainBike
	case swagger.SWIM_SportType:
		actType = &model.Swim
	case swagger.HIKE_SportType:
		actType = &model.Hike
	case swagger.WEIGHT_TRAINING_SportType:
		actType = &model.Strength
	default:
		actType = &model.Other
	}
	return actType, actTypeLst
}
// Map strava activity to the cassidy activity struct
func (s *Strava) stravaActivityToCassidyActivity(activity swagger.SummaryActivity, user *model.User) *model.Activity {
	uuid := fmt.Sprint(activity.Id)
	activityType, typeSubtypeList := s.sportTypeToActivityType(uuid, *activity.SportType)

	//userDistanceUnit := measurement.UnitClassControl(user.Units, measurement.Distance)
	//userVerticalUnit := measurement.UnitClassControl(user.Units, measurement.Vertical)
	completed := &model.Completed{
		ActivityUuid: uuid,
		Distance: measurement.CreateMeasurement(measurement.Meter, float64(activity.Distance)),
		MovingDuration: float64(activity.MovingTime),
        ElapsedDuration: float64(activity.ElapsedTime),
		Vertical: measurement.CreateMeasurement(measurement.Meter, float64(activity.TotalElevationGain)),
	}

	act := &model.Activity{
		Uuid: uuid,
		Date: activity.StartDate,
		Order: 1,
		Name: activity.Name,
		Description: "",
		Notes: "",
		Type: activityType,
		TypeSubtypeList: typeSubtypeList,
		EquipmentList: []*model.ActivityEquipment{},
		Planned: model.ZeroPlanned(uuid),
		Completed: completed,
		IsRace: false,
		NumStrides: 0,
		Map: activity.Map_.SummaryPolyline,
	}

	return act
}
// Get all strava data and load it into the database
//
// Will not attempt to merge prior activity data
func (s *Strava) BackfillData(user *model.User, token *oauth2.Token) error {
	s.App.LoadTokenDirect(token)
	activitiyPages, err := s.App.Api.GetActivities(context.TODO(), 200, nil, nil)
	if err != nil {
		return err
	}
	for _, page := range activitiyPages {
		for _, activity := range page {
			act := s.stravaActivityToCassidyActivity(activity, user)
			_, err := s.Handlers.ActivityHandler.CreateActivity(act)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (s *Strava) GetNewData(user *model.User, mostRecentActivityDate time.Time, token *oauth2.Token) error {
	s.App.LoadTokenDirect(token)
	activityPages, err := s.App.Api.GetActivities(context.TODO(), 200, nil, &mostRecentActivityDate)
	if err != nil {
		return err
	}

	for _, page := range activityPages {
		for _, activity := range page {
			act := s.stravaActivityToCassidyActivity(activity, user)
			//_, err := s.Handlers.ActivityHandler.CreateActivity(act)
            _, err := s.Handlers.ActivityHandler.CreateOrMergeActivity(act)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
