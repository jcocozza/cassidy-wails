package main

import (
	"context"
	"fmt"

	"github.com/jcocozza/cassidy-wails/internal/controllers"
	"github.com/jcocozza/cassidy-wails/internal/database"
	"github.com/jcocozza/cassidy-wails/internal/model"
)

// App struct
type App struct {
	ctx context.Context
	UserSettings *model.User
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
	usr, err := a.Handlers.UserHandler.UserRepository.Read(a.UserSettings.Username)
	fmt.Println("GOT USER:", usr)
	if err != nil {
		fmt.Println("warning: failed to load user")
		return nil
	}
	a.UserSettings = usr
	a.Handlers.SetUser(usr)
	return usr
}
func (a *App) SetUser(usr *model.User) {
	a.UserSettings = usr
	a.DB.AppUser = usr
	a.Handlers.SetUser(usr)
}
func (a *App) HasUser() bool {
	return a.UserSettings != nil
}
func (a *App) Logout() {
	a.UserSettings = nil
	a.DB.AppUser = nil
	a.Handlers.SetUser(nil)
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	/*
	userSettings, err := a.Handlers.UserHandler.UserRepository.Read("test1@test.com")
	if err != nil {
		panic("undefined user:" + err.Error())
	}

	a.SetUser(userSettings)
	/*
	a.UserSettings = *userSettings
	a.DB.AppUser = userSettings
	a.Handlers.SetUser(userSettings)
	*/
}