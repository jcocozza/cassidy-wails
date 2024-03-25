package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jcocozza/cassidy-wails/internal/database"
	"github.com/jcocozza/cassidy-wails/internal/model"
	conversionrepo "github.com/jcocozza/cassidy-wails/internal/repository/conversionRepo"
	microcyclerepo "github.com/jcocozza/cassidy-wails/internal/repository/microcycleRepo"
	ctxutil "github.com/jcocozza/cassidy-wails/internal/utils/ctxUtil"
	"github.com/jcocozza/cassidy-wails/internal/utils/dateutil"
	"github.com/jcocozza/cassidy-wails/internal/utils/measurement"
)

const forward = 2
const backward = 2

// prepend the passed microcycle to the passed list of microcycles
func prependMicrocycle(lst []*model.Microcycle, mc *model.Microcycle) []*model.Microcycle {
	newSlice := make([]*model.Microcycle, 0, len(lst)+1)
	newSlice = append(newSlice, mc)

	// Concatenate the rest of the original slice
	newSlice = append(newSlice, lst...)

	return newSlice
}

type MicrocycleHandler struct {
	MicrocycleRepository microcyclerepo.MicrocycleRepository
	ConversionRepository conversionrepo.MeasurementRepository
}

func NewMicrocycleHandler(db database.DbOperations) *MicrocycleHandler {
	return &MicrocycleHandler{
		MicrocycleRepository: microcyclerepo.NewIMicrocycleRepository(db),
		ConversionRepository: conversionrepo.NewIMeasurementRepository(),
	}
}

// Get a microcycle for a given date range
//
// @param start_date
// @param end_date
func (mh *MicrocycleHandler) GetMicrocycle(ctx *gin.Context) {
	startDate := ctx.Param("start_date")
	endDate := ctx.Param("end_date")
	userUuid := ctxutil.GetUserUuidFromHeader(ctx)
	userUnitClass := ctxutil.GetUserUnitsFromHeader(ctx)

	microcycle, err := mh.MicrocycleRepository.ReadMicrocycle(startDate, endDate, userUuid, measurement.UnitClass(userUnitClass))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err1 := mh.ConversionRepository.ConvertMicrocycle(conversionrepo.Outgoing, microcycle, measurement.UnitClass(userUnitClass))
	if err1 != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "microcycle failed to convert: " + err1.Error()})
		return
	}
	ctx.JSON(http.StatusOK, microcycle)
}

// Get a list of microcycles
//
// @param: start_date
// @param: end_date
func (mh *MicrocycleHandler) GetCalendar(ctx *gin.Context) {
	startDate := ctx.Param("start_date")
	endDate := ctx.Param("end_date")
	userUuid := ctxutil.GetUserUuidFromHeader(ctx)
	userUnitClass := ctxutil.GetUserUnitsFromHeader(ctx)

	centerMicrocycle, err := mh.MicrocycleRepository.ReadMicrocycle(startDate, endDate, userUuid, measurement.UnitClass(userUnitClass))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var tmpStartFwd = startDate
	var tmpEndFwd = endDate
	var forwardCycles []*model.Microcycle
	for i := 0; i < forward; i++ {
		tmpStartFwd, tmpEndFwd = dateutil.GetNextCycle(tmpStartFwd, tmpEndFwd)
		mc, err := mh.MicrocycleRepository.ReadMicrocycle(tmpStartFwd, tmpEndFwd, userUuid, measurement.UnitClass(userUnitClass))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		forwardCycles = append(forwardCycles, mc)
	}

	var tmpStartBwd = startDate
	var tmpEndBwd = endDate
	var backwardCycles []*model.Microcycle
	for i := 0; i < backward; i++ {
		tmpStartBwd, tmpEndBwd = dateutil.GetPreviousCycle(tmpStartBwd, tmpEndBwd)
		mc, err := mh.MicrocycleRepository.ReadMicrocycle(tmpStartBwd, tmpEndBwd, userUuid, measurement.UnitClass(userUnitClass))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		backwardCycles = prependMicrocycle(backwardCycles, mc)
	}

	var calendarMicrocycleList []*model.Microcycle

	calendarMicrocycleList = append(calendarMicrocycleList, backwardCycles...)
	calendarMicrocycleList = append(calendarMicrocycleList, centerMicrocycle)
	calendarMicrocycleList = append(calendarMicrocycleList, forwardCycles...)

	for _, mc := range calendarMicrocycleList {
		err1 := mh.ConversionRepository.ConvertMicrocycle(conversionrepo.Outgoing, mc, measurement.UnitClass(userUnitClass))
		if err1 != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "microcycle failed to convert: " + err1.Error()})
			return
		}
	}

	ctx.JSON(http.StatusOK, calendarMicrocycleList)
}

// Get the next number of microcycles
//
// @param: start_date
// @param: end_date
// @param: number
func (mh *MicrocycleHandler) GetNextNMicrocycles(ctx *gin.Context) {
	startDate := ctx.Param("start_date")
	endDate := ctx.Param("end_date")

	userUuid := ctxutil.GetUserUuidFromHeader(ctx)
	userUnitClass := ctxutil.GetUserUnitsFromHeader(ctx)

	numberCylesStr := ctx.Param("number")
	numberCycles, err := strconv.Atoi(numberCylesStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid number of cycles:" + err.Error()})
		return
	}

	var tmpStart = startDate
	var tmpEnd = endDate
	var microcycleList = []*model.Microcycle{}
	for i := 0; i < numberCycles; i++ {
		tmpStart, tmpEnd = dateutil.GetNextCycle(tmpStart, tmpEnd)
		mc, err := mh.MicrocycleRepository.ReadMicrocycle(tmpStart, tmpEnd, userUuid, measurement.UnitClass(userUnitClass))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		microcycleList = append(microcycleList, mc)
	}

	for _, mc := range microcycleList {
		err1 := mh.ConversionRepository.ConvertMicrocycle(conversionrepo.Outgoing, mc, measurement.UnitClass(userUnitClass))
		if err1 != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "microcycle failed to convert: " + err1.Error()})
			return
		}
	}

	ctx.JSON(http.StatusOK, microcycleList)
}

// Get the previous number of microcycles
//
// @param: start_date
// @param: end_date
// @param: number
func (mh *MicrocycleHandler) GetPreviousNMicrocycles(ctx *gin.Context) {
	startDate := ctx.Param("start_date")
	endDate := ctx.Param("end_date")

	userUuid := ctxutil.GetUserUuidFromHeader(ctx)
	userUnitClass := ctxutil.GetUserUnitsFromHeader(ctx)

	numberCylesStr := ctx.Param("number")
	numberCycles, err := strconv.Atoi(numberCylesStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid number of cycles" + err.Error()})
		return
	}

	var tmpStart = startDate
	var tmpEnd = endDate
	var microcycleList = []*model.Microcycle{}
	for i := 0; i < numberCycles; i++ {
		tmpStart, tmpEnd = dateutil.GetPreviousCycle(tmpStart, tmpEnd)
		mc, err := mh.MicrocycleRepository.ReadMicrocycle(tmpStart, tmpEnd, userUuid, measurement.UnitClass(userUnitClass))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		microcycleList = prependMicrocycle(microcycleList, mc)
	}

	for _, mc := range microcycleList {
		err1 := mh.ConversionRepository.ConvertMicrocycle(conversionrepo.Outgoing, mc, measurement.UnitClass(userUnitClass))
		if err1 != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "microcycle failed to convert: " + err1.Error()})
			return
		}
	}

	ctx.JSON(http.StatusOK, microcycleList)
}
