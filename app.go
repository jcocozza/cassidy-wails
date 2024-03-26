package main

import (
	"context"

	"github.com/jcocozza/cassidy-wails/internal/controllers"
	"github.com/jcocozza/cassidy-wails/internal/database"
	"github.com/jcocozza/cassidy-wails/internal/model"
)

// App struct
type App struct {
	ctx context.Context
	UserSettings model.User
	DB *database.Database
	Handlers *controllers.Controllers
}

// NewApp creates a new App application struct
func NewApp() *App {
	DB := database.InitTestDB()
	handlers := controllers.NewControllers(DB, model.EmptyUser())

	return &App{
		DB: DB,
		Handlers: handlers,
	}
}

func (a *App) LoadUser() *model.User {
	return &a.UserSettings
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	userSettings, err := a.Handlers.UserHandler.UserRepository.Read("test1@test.com")

	if err != nil {
		panic("undefined user:" + err.Error())
	}

	a.UserSettings = *userSettings
	a.DB.AppUser = userSettings
	a.Handlers.SetUser(userSettings)
}