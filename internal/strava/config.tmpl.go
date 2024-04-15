package strava

const (
	ClientId_tmpl     string = "client_id"
	ClientSecret_tmpl string = "client_secret"
	RefreshToken_tmpl string = "refresh_token"
	RedirectUri_tmpl  string = "http://localhost:9999/strava/callback" //"http://localhost:5173/user"//"http://localhost/exchange_token" //
)
var (
	Scopes_tmpl = []string{"activity:read_all"}
)