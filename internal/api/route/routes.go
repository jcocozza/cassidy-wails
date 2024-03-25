package route

import (
	"github.com/gin-gonic/gin"
	"github.com/jcocozza/cassidy-wails/internal/api/controllers"
	"github.com/jcocozza/cassidy-wails/internal/database"
)

type Handlers struct {
	UserHandler *controllers.UserHandler
	ActivityHandler *controllers.ActivityHandler
	ActivityTypeHandler *controllers.ActivityTypeHandler
	EquipmentHandler *controllers.EquipmentHandler
	MicrocycleHandler *controllers.MicrocycleHandler
	MiscHandler *controllers.MiscHandler
}

func InitHandlers(db database.DbOperations) *Handlers {
	return &Handlers{
		UserHandler: controllers.NewUserHandler(db),
		ActivityHandler: controllers.NewActivityHandler(db),
		ActivityTypeHandler: controllers.NewActivityTypeHandler(db),
		EquipmentHandler: controllers.NewEquipmentHandler(db),
		MicrocycleHandler: controllers.NewMicrocycleHandler(db),
		MiscHandler: controllers.NewMiscHandler(db),
	}
}

func InitRoutes(r *gin.Engine, db database.DbOperations) {
	h := InitHandlers(db)

	userGroup := r.Group("/user")
	// create user
	userGroup.PUT("/create", h.UserHandler.CreateUser)
	// authenticate user
	userGroup.POST("/auth", h.UserHandler.AuthenticateUser)
	// update user
	userGroup.PATCH("/update", h.UserHandler.UpdateUser)

	activityGroup := r.Group("/activity")
	// create an activity
	activityGroup.PUT("/create", h.ActivityHandler.CreateActivity)
	// update an activity at a given id
	activityGroup.PATCH("/:uuid", h.ActivityHandler.UpdateActivity)
	// delete an activity at a given id
	activityGroup.DELETE("/:uuid", h.ActivityHandler.DeleteActivity)

	activityTypesGroup := r.Group("/activity/type")
	// list all activity types
	activityTypesGroup.GET("/list", h.ActivityTypeHandler.ListActivityTypes)

	dataGroup := r.Group("/data")
	// get the n cycle summary
	dataGroup.GET("n-cycle-summary/:start_date/:end_date", h.MiscHandler.GetNCycleSummary)

	dateUtilGroup := r.Group("/date")
	// get the current microcycle dates
	dateUtilGroup.GET("/microcycle/current", h.UserHandler.GetMicrocycleCurrentDates)
	// get the next and previous microcycle dates
	dateUtilGroup.GET("/cycle/:start_date/:end_date/next-previous", h.MiscHandler.GetNextPrevious)

	equipmentGroup := r.Group("/equipment")
	// create equipment
	equipmentGroup.PUT("/create", h.EquipmentHandler.CreateEquipment)
	// update equipment by id
	equipmentGroup.PATCH("/:id", h.EquipmentHandler.UpdateEquipment)
	// delete equipment by id
	equipmentGroup.DELETE("/:id", h.EquipmentHandler.DeleteEquipment)
	// list all equipment
	equipmentGroup.GET("/list", h.EquipmentHandler.List)
	// list all equipment types
	equipmentGroup.GET("/type/list", h.EquipmentHandler.ListEquipmentTypes)

	activityEquipmentGroup := r.Group("/activity/equipment")
	// create activity equipment
	activityEquipmentGroup.PUT("/create", h.EquipmentHandler.CreateActivityEquipment)
	// update activity equipment by id
	activityEquipmentGroup.PATCH("/:id", h.EquipmentHandler.UpdateActivityEquipment)
	// delete activity equipment by id
	activityEquipmentGroup.DELETE("/:id", h.EquipmentHandler.DeleteActivityEquipment)


	microcycleGroup := r.Group("/microcycle")
	// get a microcycle for a start and end date pair
	microcycleGroup.GET("/:start_date/:end_date", h.MicrocycleHandler.GetMicrocycle)
	// get a calendar around a given microcycle start and end date pair
	microcycleGroup.GET("/calendar/:start_date/:end_date", h.MicrocycleHandler.GetCalendar)
	// get the next N microcycles for a given start/end date pair
	microcycleGroup.GET("/:start_date/:end_date/next/:number", h.MicrocycleHandler.GetNextNMicrocycles)
	// get the previous N microcycles for a given start/end date pair
	microcycleGroup.GET("/:start_date/:end_date/previous/:number", h.MicrocycleHandler.GetPreviousNMicrocycles)
}
