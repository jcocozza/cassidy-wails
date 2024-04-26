package model

import (
	"time"

	"github.com/jcocozza/cassidy-wails/internal/utils/measurement"
)

// An n cycle summary represents a summary of n cycles.
//
// It is an ordered date list where each element in the date list has a corresponding element in several other lists which
// provide totals for their respective purposes.
type NCycleSummary struct {
	StartDateList    []time.Time     			`json:"start_date_list" ts_type:"Date[]"`
	PlannedDistances []*measurement.Measurement `json:"planned_distances"`
	PlannedDurations []float64                  `json:"planned_durations"`
	PlannedVerticals []*measurement.Measurement `json:"planned_verticals"`

	CompletedDistances []*measurement.Measurement `json:"completed_distances"`
	CompletedDurations []float64                  `json:"completed_durations"`
	CompletedVerticals []*measurement.Measurement `json:"completed_verticals"`
}
