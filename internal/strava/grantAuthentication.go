package strava

import (
	"fmt"

	"github.com/jcocozza/cassidy-connector/strava/app"
)

type Strava struct {
	App *app.App
}

func NewStravaApp() *Strava {
	return &Strava{
		App: app.NewApp(ClientId, ClientSecret, RedirectUri, Scopes),
	}
}
func (s *Strava) StartListener() {
	token, err := s.App.AwaitInitialToken()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(token)
}

func (s *Strava) OpenStravaAuth() {
	s.App.OpenAuthorizationGrant()
}