package main

import (
	"embed"
	"os"

	"github.com/jcocozza/cassidy-wails/internal/controllers"
	"github.com/jcocozza/cassidy-wails/internal/database"
	"github.com/jcocozza/cassidy-wails/internal/model"
	"github.com/jcocozza/cassidy-wails/internal/strava"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/build
var assets embed.FS

func writeErrToFile(e error) {
    file, err := os.OpenFile("err_log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
    if err != nil {
        return
    }
    defer file.Close()

    _, err1 := file.WriteString(e.Error())
    if err1 != nil {
        return
    }
}

func main() {
	// Create an instance of the app structure
	app := NewApp()

	//DB := database.InitTestDB()
	DB, err0 := database.ConnectToCassidyDB()
	if err0 != nil { // if we can't find the application database, something has gone very wrong
        writeErrToFile(err0)
		panic("app database not found" + err0.Error())
	}
	handlers := controllers.NewControllers(DB, model.EmptyUser())

	app.Handlers = handlers
	app.DB = DB

	stravaApp := strava.NewStravaApp(handlers)
	// Create application with options
	err := wails.Run(&options.App{
		Title:  "cassidy",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:  app.startup,
		OnShutdown: app.shutdown,
		Bind: []interface{}{
			app,
			app.Handlers.UserHandler,
			app.Handlers.ActivityHandler,
			app.Handlers.ActivityTypeHandler,
			app.Handlers.EquipmentHandler,
			app.Handlers.MicrocycleHandler,
			app.Handlers.MiscHandler,
			stravaApp,
		},
	})

	if err != nil {
		println("Error:", err.Error())
        writeErrToFile(err)
	}
}
