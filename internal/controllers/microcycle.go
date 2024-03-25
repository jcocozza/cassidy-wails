package controllers

import (
	"github.com/jcocozza/cassidy-wails/internal/database"
	"github.com/jcocozza/cassidy-wails/internal/model"
	conversionrepo "github.com/jcocozza/cassidy-wails/internal/repository/conversionRepo"
	microcyclerepo "github.com/jcocozza/cassidy-wails/internal/repository/microcycleRepo"
	"github.com/jcocozza/cassidy-wails/internal/utils/dateutil"
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
func (mh *MicrocycleHandler) GetMicrocycle(startDate, endDate string, user *model.User) (*model.Microcycle, error) {
	microcycle, err := mh.MicrocycleRepository.ReadMicrocycle(startDate, endDate, user.Uuid, user.Units)
	if err != nil {
		return nil, err
	}

	err1 := mh.ConversionRepository.ConvertMicrocycle(conversionrepo.Outgoing, microcycle, user.Units)
	if err1 != nil {
		return nil, err1
	}
	return microcycle, nil
}
// Get a list of microcycles
//
// @param: start_date
// @param: end_date
func (mh *MicrocycleHandler) GetCalendar(startDate, endDate string, user *model.User) ([]*model.Microcycle, error) {
	centerMicrocycle, err := mh.MicrocycleRepository.ReadMicrocycle(startDate, endDate, user.Uuid, user.Units)
	if err != nil {
		return nil, err
	}

	var tmpStartFwd = startDate
	var tmpEndFwd = endDate
	var forwardCycles []*model.Microcycle
	for i := 0; i < forward; i++ {
		tmpStartFwd, tmpEndFwd = dateutil.GetNextCycle(tmpStartFwd, tmpEndFwd)
		mc, err := mh.MicrocycleRepository.ReadMicrocycle(tmpStartFwd, tmpEndFwd, user.Uuid, user.Units)
		if err != nil {
			return nil, err
		}
		forwardCycles = append(forwardCycles, mc)
	}

	var tmpStartBwd = startDate
	var tmpEndBwd = endDate
	var backwardCycles []*model.Microcycle
	for i := 0; i < backward; i++ {
		tmpStartBwd, tmpEndBwd = dateutil.GetPreviousCycle(tmpStartBwd, tmpEndBwd)
		mc, err := mh.MicrocycleRepository.ReadMicrocycle(tmpStartBwd, tmpEndBwd, user.Uuid, user.Units)
		if err != nil {
			return nil, err
		}
		backwardCycles = prependMicrocycle(backwardCycles, mc)
	}

	var calendarMicrocycleList []*model.Microcycle

	calendarMicrocycleList = append(calendarMicrocycleList, backwardCycles...)
	calendarMicrocycleList = append(calendarMicrocycleList, centerMicrocycle)
	calendarMicrocycleList = append(calendarMicrocycleList, forwardCycles...)

	for _, mc := range calendarMicrocycleList {
		err1 := mh.ConversionRepository.ConvertMicrocycle(conversionrepo.Outgoing, mc, user.Units)
		if err1 != nil {
			return nil, err1
		}
	}

	return calendarMicrocycleList, nil
}
// Get the next number of microcycles
//
// @param: start_date
// @param: end_date
// @param: number
func (mh *MicrocycleHandler) GetNextNMicrocycles(startDate, endDate string, numberCycles int, user *model.User) ([]*model.Microcycle, error) {
	var tmpStart = startDate
	var tmpEnd = endDate
	var microcycleList = []*model.Microcycle{}
	for i := 0; i < numberCycles; i++ {
		tmpStart, tmpEnd = dateutil.GetNextCycle(tmpStart, tmpEnd)
		mc, err := mh.MicrocycleRepository.ReadMicrocycle(tmpStart, tmpEnd, user.Uuid, user.Units)
		if err != nil {
			return nil, err
		}
		microcycleList = append(microcycleList, mc)
	}

	for _, mc := range microcycleList {
		err1 := mh.ConversionRepository.ConvertMicrocycle(conversionrepo.Outgoing, mc, user.Units)
		if err1 != nil {
			return nil, err1
		}
	}

	return microcycleList, nil
}

// Get the previous number of microcycles
//
// @param: start_date
// @param: end_date
// @param: number
func (mh *MicrocycleHandler) GetPreviousNMicrocycles(startDate, endDate string, numberCycles int, user *model.User) ([]*model.Microcycle, error) {
	var tmpStart = startDate
	var tmpEnd = endDate
	var microcycleList = []*model.Microcycle{}
	for i := 0; i < numberCycles; i++ {
		tmpStart, tmpEnd = dateutil.GetPreviousCycle(tmpStart, tmpEnd)
		mc, err := mh.MicrocycleRepository.ReadMicrocycle(tmpStart, tmpEnd, user.Uuid, user.Units)
		if err != nil {
			return nil, err
		}
		microcycleList = prependMicrocycle(microcycleList, mc)
	}

	for _, mc := range microcycleList {
		err1 := mh.ConversionRepository.ConvertMicrocycle(conversionrepo.Outgoing, mc, user.Units)
		if err1 != nil {
			return nil, err1
		}
	}

	return microcycleList, nil
}
