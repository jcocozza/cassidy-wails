package model

import (
	"fmt"

	"github.com/jcocozza/cassidy-wails/internal/utils/measurement"
)

// Represents planned data for an activity.
type Planned struct {
	ActivityUuid string                   `json:"activity_uuid"`
	Distance     *measurement.Measurement `json:"distance"`
	Duration     float64                  `json:"duration"`
	Vertical     *measurement.Measurement `json:"vertical"`
	Pace         string                   `json:"pace"`
}

// An empty planned has no uuid, units and has all lengths set to -1.
func EmptyPlanned() *Planned {
	return &Planned{
		ActivityUuid: "",
		Distance:     measurement.EmptyMeasurement(),
		Duration:     -1,
		Vertical:     measurement.EmptyMeasurement(),
	}
}

// Zero planned has a passed uuid, lengths of 0 and a duration of 0
func ZeroPlanned(activityUuid string) *Planned {
	return &Planned{
		ActivityUuid: activityUuid,
		Distance:     measurement.DefaultMeasurement(),
		Duration:     0,
		Vertical:     measurement.DefaultMeasurement(),
	}
}

// Validate a planned object.
//
// Validation ensures that uuid, distance, vertical, duration and units are valid.
func (p *Planned) Validate() error {
	if p.ActivityUuid == "" {
		return fmt.Errorf("planned activity uuid is invalid")
	}
	err := p.Distance.Validate()
	if err != nil {
		return fmt.Errorf("planned distance is invalid: %w", err)
	}
	if p.Duration == -1 {
		return fmt.Errorf("planned duration is invalid")
	}
	err2 := p.Vertical.Validate()
	if err2 != nil {
		return fmt.Errorf("planned vertical is invalid: %w", err2)
	}
	return nil
}

// Calculate the pace of the planned
func (p *Planned) CalculatePace(paceUnit measurement.PaceUnit) {
	pace := measurement.CalculatePace(p.Distance, p.Duration, paceUnit)
	p.Pace = pace.String()
}

// Return true if all values of the planned are zero
func (p *Planned) IsZero() bool {
	return p.Distance.Length == 0 && p.Duration == 0 && p.Vertical.Length == 0
}
