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
// Get the token during the initial authentication process
func (s *Strava) StartListener() {
	// note this will also set the token in the app struct
	token, err := s.App.AwaitInitialToken(-1)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(token)
}
// Open the authentication link in the user's browser
func (s *Strava) OpenStravaAuth() {
	s.App.OpenAuthorizationGrant()
}