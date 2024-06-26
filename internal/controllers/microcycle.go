package controllers

import (
	"fmt"
	"time"

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
	User *model.User
}

func NewMicrocycleHandler(db database.DbOperations, user *model.User) *MicrocycleHandler {
	return &MicrocycleHandler{
		MicrocycleRepository: microcyclerepo.NewIMicrocycleRepository(db),
		ConversionRepository: conversionrepo.NewIMeasurementRepository(),
		User: user,
	}
}

// Get a microcycle for a given date range
//
// @param start_date
// @param end_date
func (mh *MicrocycleHandler) GetMicrocycle(startDate, endDate time.Time) (*model.Microcycle, error) {
	microcycle, err := mh.MicrocycleRepository.ReadMicrocycle(startDate, endDate, mh.User.Uuid, mh.User.Units)
	if err != nil {
		return nil, err
	}

	err1 := mh.ConversionRepository.ConvertMicrocycle(conversionrepo.Outgoing, microcycle, mh.User.Units)
	if err1 != nil {
		return nil, err1
	}
	return microcycle, nil
}

func (mh *MicrocycleHandler) GetCalendar() ([]*model.Microcycle, error) {
	const numberCyclesForward = 7
	const numberCyclesBackward = 8
	now := time.Now()
	fmt.Println("it is currently: ", now)
	var current []time.Time
	if (mh.User.CycleDays) == 7 {
		current = dateutil.GetDateMicrocycle(now, mh.User.CycleStart, mh.User.CycleDays)
	} else {
		current = dateutil.GetCurrentCycleFromInitialDate(mh.User.InitialCycleStart, mh.User.CycleDays)
	}

	currentStart := current[0]
	currentEnd := current[len(current) - 1]

	centerMicrocycle, err := mh.MicrocycleRepository.ReadMicrocycle(currentStart, currentEnd, mh.User.Uuid, mh.User.Units)
	if err != nil {
		return nil, err
	}

	var tmpStartFwd = currentStart
	var tmpEndFwd = currentEnd
	var forwardCycles []*model.Microcycle
	for i := 0; i < numberCyclesForward; i++ {
		tmpStartFwd, tmpEndFwd = dateutil.GetNextCycle(tmpStartFwd, tmpEndFwd)
		mc, err := mh.MicrocycleRepository.ReadMicrocycle(tmpStartFwd, tmpEndFwd, mh.User.Uuid, mh.User.Units)
		if err != nil {
			return nil, err
		}
		forwardCycles = append(forwardCycles, mc)
	}

	var tmpStartBwd = currentStart
	var tmpEndBwd = currentEnd
	var backwardCycles []*model.Microcycle
	for i := 0; i <  numberCyclesBackward; i++ {
		fmt.Printf("getting previous cycle for %s, %s \n", tmpStartBwd, tmpEndBwd)
		tmpStartBwd, tmpEndBwd = dateutil.GetPreviousCycle(tmpStartBwd, tmpEndBwd)
		fmt.Printf("new cycle date: %s, %s \n\n", tmpStartBwd, tmpEndBwd)
		mc, err := mh.MicrocycleRepository.ReadMicrocycle(tmpStartBwd, tmpEndBwd, mh.User.Uuid, mh.User.Units)
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
		err1 := mh.ConversionRepository.ConvertMicrocycle(conversionrepo.Outgoing, mc, mh.User.Units)
		if err1 != nil {
			return nil, err1
		}
	}

	return calendarMicrocycleList, nil
}
/*
// Get a list of microcycles
//
// @param: start_date
// @param: end_date
func (mh *MicrocycleHandler) GetCalendar(startDate, endDate time.Time) ([]*model.Microcycle, error) {
	centerMicrocycle, err := mh.MicrocycleRepository.ReadMicrocycle(startDate, endDate, mh.User.Uuid, mh.User.Units)
	if err != nil {
		return nil, err
	}

	var tmpStartFwd = startDate
	var tmpEndFwd = endDate
	var forwardCycles []*model.Microcycle
	for i := 0; i < forward; i++ {
		tmpStartFwd, tmpEndFwd = dateutil.GetNextCycle(tmpStartFwd, tmpEndFwd)
		mc, err := mh.MicrocycleRepository.ReadMicrocycle(tmpStartFwd, tmpEndFwd, mh.User.Uuid, mh.User.Units)
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
		mc, err := mh.MicrocycleRepository.ReadMicrocycle(tmpStartBwd, tmpEndBwd, mh.User.Uuid, mh.User.Units)
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
		err1 := mh.ConversionRepository.ConvertMicrocycle(conversionrepo.Outgoing, mc, mh.User.Units)
		if err1 != nil {
			return nil, err1
		}
	}

	return calendarMicrocycleList, nil
}
*/

// Get the next number of microcycles
//
// @param: start_date
// @param: end_date
// @param: number
func (mh *MicrocycleHandler) GetNextNMicrocycles(startDate, endDate time.Time, numberCycles int) ([]*model.Microcycle, error) {
	var tmpStart = startDate
	var tmpEnd = endDate
	var microcycleList = []*model.Microcycle{}
	for i := 0; i < numberCycles; i++ {
		tmpStart, tmpEnd = dateutil.GetNextCycle(tmpStart, tmpEnd)
		mc, err := mh.MicrocycleRepository.ReadMicrocycle(tmpStart, tmpEnd, mh.User.Uuid, mh.User.Units)
		if err != nil {
			return nil, err
		}
		microcycleList = append(microcycleList, mc)
	}

	for _, mc := range microcycleList {
		err1 := mh.ConversionRepository.ConvertMicrocycle(conversionrepo.Outgoing, mc, mh.User.Units)
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
func (mh *MicrocycleHandler) GetPreviousNMicrocycles(startDate, endDate time.Time, numberCycles int) ([]*model.Microcycle, error) {
	var tmpStart = startDate
	var tmpEnd = endDate
	var microcycleList = []*model.Microcycle{}
	for i := 0; i < numberCycles; i++ {
		tmpStart, tmpEnd = dateutil.GetPreviousCycle(tmpStart, tmpEnd)
		mc, err := mh.MicrocycleRepository.ReadMicrocycle(tmpStart, tmpEnd, mh.User.Uuid, mh.User.Units)
		if err != nil {
			return nil, err
		}
		microcycleList = prependMicrocycle(microcycleList, mc)
	}

	for _, mc := range microcycleList {
		err1 := mh.ConversionRepository.ConvertMicrocycle(conversionrepo.Outgoing, mc, mh.User.Units)
		if err1 != nil {
			return nil, err1
		}
	}

	return microcycleList, nil
}
