package controllers

import (
	"github.com/jcocozza/cassidy-wails/internal/database"
	"github.com/jcocozza/cassidy-wails/internal/model"
	conversionrepo "github.com/jcocozza/cassidy-wails/internal/repository/conversionRepo"
	miscrepo "github.com/jcocozza/cassidy-wails/internal/repository/miscRepo"
	"github.com/jcocozza/cassidy-wails/internal/utils/dateutil"
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
func (mh *MiscHandler) GetNCycleSummary(startDate, endDate string, user *model.User) (*model.NCycleSummary, error) {
	ncycleSummary, err := mh.MiscRepository.ReadNCycleSummary(startDate, endDate, user.Uuid)
	if err != nil {
		return nil, err
	}

	err1 := mh.ConversionRepository.ConvertNCycleSummary(conversionrepo.Outgoing, ncycleSummary, user.Units)
	if err1 != nil {
		return nil, err1
	}
	return ncycleSummary, nil
}

// Get the next/previous microcycle start and end dates
//
// @param start_date
// @param end_date
type NextPrevious struct {
	NextStartDate string `json:"next_start_date"`
	NextEndDate string `json:"next_end_date"`
	PreviousStartDate string `json:"previous_start_date"`
	PreviousEndDate string `json:"previous_end_date"`
}
func (mh *MiscHandler) GetNextPrevious(startDate, endDate string) (*NextPrevious, error ){
	nextStart, nextEnd, previousStart, previousEnd, err := dateutil.GetNextPrevious(startDate, endDate)

	if err != nil {
		return nil, err
	}

	mp := &NextPrevious{
		NextStartDate:     nextStart,
		NextEndDate:       nextEnd,
		PreviousStartDate: previousStart,
		PreviousEndDate:   previousEnd,
	}

	return mp, nil
}
