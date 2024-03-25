package api

import (
	"github.com/jcocozza/cassidy-wails/internal/api/middleware"
	"github.com/jcocozza/cassidy-wails/internal/api/route"
	"github.com/jcocozza/cassidy-wails/internal/database"
)

// Set up and run the API
func SetUpAndRunAPI(db database.DbOperations) {
	router := route.CreateRouter()
	middleware.SetCors(router)
	route.InitRoutes(router, db)

	router.Run(":8080")
}
