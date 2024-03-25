package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jcocozza/cassidy-wails/internal/database"
	"github.com/jcocozza/cassidy-wails/internal/model"
	activityrepo "github.com/jcocozza/cassidy-wails/internal/repository/activityRepo"
	conversionrepo "github.com/jcocozza/cassidy-wails/internal/repository/conversionRepo"
	ctxutil "github.com/jcocozza/cassidy-wails/internal/utils/ctxUtil"
	"github.com/jcocozza/cassidy-wails/internal/utils/measurement"
	"github.com/jcocozza/cassidy-wails/internal/utils/uuidgen"
)

type ActivityHandler struct {
	ActivityRepository   activityrepo.ActivityRepository
	ConversionRepository conversionrepo.MeasurementRepository
}

func NewActivityHandler(db database.DbOperations) *ActivityHandler {
	return &ActivityHandler{
		ActivityRepository:   activityrepo.NewIActivityRepository(db),
		ConversionRepository: conversionrepo.NewIMeasurementRepository(),
	}
}

// Create an activity
func (ah *ActivityHandler) CreateActivity(ctx *gin.Context) {
	userUnitClass := measurement.UnitClass(ctxutil.GetUserUnitsFromHeader(ctx))
	var createRequest model.Activity
	if err := ctx.ShouldBindJSON(&createRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err0 := ah.ConversionRepository.ConvertActivity(conversionrepo.Incoming, &createRequest, userUnitClass)
	if err0 != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "activity failed to convert" + err0.Error()})
		return
	}

	userUuid := ctxutil.GetUserUuidFromHeader(ctx)

	newActUuid := uuidgen.GenerateUUID()
	createRequest.SetUuid(newActUuid)

	err := createRequest.Validate()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "activity failed to validate")
		return
	}

	err2 := ah.ActivityRepository.Create(userUuid, &createRequest)
	if err2 != nil {
		ctx.JSON(http.StatusInternalServerError, "failed to create activity")
		return
	} else {
		err3 := ah.ConversionRepository.ConvertActivity(conversionrepo.Outgoing, &createRequest, userUnitClass)
		if err3 != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "activity failed to convert" + err3.Error()})
			return
		}
		ctx.JSON(http.StatusCreated, createRequest)
	}
}

// Update an activity
func (ah *ActivityHandler) UpdateActivity(ctx *gin.Context) {
	userUnitClass := measurement.UnitClass(ctxutil.GetUserUnitsFromHeader(ctx))
	var updateRequest model.Activity
	if err := ctx.ShouldBindJSON(&updateRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err0 := ah.ConversionRepository.ConvertActivity(conversionrepo.Incoming, &updateRequest, userUnitClass)
	if err0 != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "activity failed to convert incoming: " + err0.Error()})
		return
	}

	err := ah.ActivityRepository.Update(&updateRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update: " + err.Error()})
		return
	}
	// return updated object
	err3 := ah.ConversionRepository.ConvertActivity(conversionrepo.Outgoing, &updateRequest, userUnitClass)
	if err3 != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "activity failed to convert outgoing: " + err3.Error()})
		return
	}
	ctx.JSON(http.StatusAccepted, updateRequest)
}

// Delete an activity
//
// @param: uuid
func (ah *ActivityHandler) DeleteActivity(ctx *gin.Context) {
	activityUuid := ctx.Param("uuid")

	err := ah.ActivityRepository.Delete(activityUuid)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete activity: " + err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, "delete")
}
