package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jcocozza/cassidy-wails/internal/database"
	conversionrepo "github.com/jcocozza/cassidy-wails/internal/repository/conversionRepo"
	miscrepo "github.com/jcocozza/cassidy-wails/internal/repository/miscRepo"
	ctxutil "github.com/jcocozza/cassidy-wails/internal/utils/ctxUtil"
	"github.com/jcocozza/cassidy-wails/internal/utils/dateutil"
	"github.com/jcocozza/cassidy-wails/internal/utils/measurement"
)

// Handle misc data processes
type MiscHandler struct {
	MiscRepository       miscrepo.MiscRepository
	ConversionRepository conversionrepo.MeasurementRepository
}

func NewMiscHandler(db database.DbOperations) *MiscHandler {
	return &MiscHandler{
		MiscRepository:       miscrepo.NewIMiscRepository(db),
		ConversionRepository: conversionrepo.NewIMeasurementRepository(),
	}
}

// Get an n cycle summary
//
// @param: start_date
// @param: end_date
func (mh *MiscHandler) GetNCycleSummary(ctx *gin.Context) {
	startDate := ctx.Param("start_date")
	endDate := ctx.Param("end_date")
	userUuid := ctxutil.GetUserUuidFromHeader(ctx)
	userUnitClass := ctxutil.GetUserUnitsFromHeader(ctx)

	ncycleSummary, err := mh.MiscRepository.ReadNCycleSummary(startDate, endDate, userUuid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err1 := mh.ConversionRepository.ConvertNCycleSummary(conversionrepo.Outgoing, ncycleSummary, measurement.UnitClass(userUnitClass))
	if err1 != nil {
		ctx.JSON(http.StatusInternalServerError, "microcycle failed to convert")
	}
	ctx.JSON(http.StatusOK, ncycleSummary)
}

// Get the next/previous microcycle start and end dates
//
// @param start_date
// @param end_date
func (mh *MiscHandler) GetNextPrevious(ctx *gin.Context) {
	startDate := ctx.Param("start_date")
	endDate := ctx.Param("end_date")

	nextStart, nextEnd, previousStart, previousEnd, err := dateutil.GetNextPrevious(startDate, endDate)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid start date/end date:"+fmt.Sprint(err))
		return
	}

	data := gin.H{
		"next_start_date":     nextStart,
		"next_end_date":       nextEnd,
		"previous_start_date": previousStart,
		"previous_end_date":   previousEnd,
	}

	ctx.JSON(http.StatusOK, data)
}
