package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jcocozza/cassidy-wails/internal/database"
	activitytyperepo "github.com/jcocozza/cassidy-wails/internal/repository/activityTypeRepo"
)

type ActivityTypeHandler struct {
	ActivityTypeRepository activitytyperepo.ActivityTypeRepository
}

func NewActivityTypeHandler(db database.DbOperations) *ActivityTypeHandler {
	return &ActivityTypeHandler{
		ActivityTypeRepository: activitytyperepo.NewIActivityTypeRepository(db),
	}
}

// list all activity types
func (ath *ActivityTypeHandler) ListActivityTypes(ctx *gin.Context) {
	activityTypeWithSubtypeList, err := ath.ActivityTypeRepository.List()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "something went wrong creating activity type subtype list: "+err.Error())
	}

	ctx.JSON(http.StatusOK, activityTypeWithSubtypeList)
}
