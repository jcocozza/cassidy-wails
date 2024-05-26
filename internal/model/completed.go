package model

import (
	"fmt"

	"github.com/jcocozza/cassidy-wails/internal/utils/measurement"
)

// Represents completed data for an activity.
type Completed struct {
	ActivityUuid    string                   `json:"activity_uuid"`
	Distance        *measurement.Measurement `json:"distance"`
	MovingDuration  float64                  `json:"moving_duration"`
    ElapsedDuration float64                  `json:"elapsed_duration"`
	Vertical        *measurement.Measurement `json:"vertical"`
	Pace            string                   `json:"pace"`
}

// An empty completed has no uuid, units and has all lengths set to -1.
func EmptyCompleted() *Completed {
	return &Completed{
		ActivityUuid: "",
		Distance:     measurement.EmptyMeasurement(),
        MovingDuration: -1,
		ElapsedDuration:     -1,
		Vertical:     measurement.EmptyMeasurement(),
	}
}

// Zero completed has a passed uuid, lengths of 0 and a duration of 0
func ZeroCompleted(activityUuid string) *Completed {
	return &Completed{
		ActivityUuid: activityUuid,
		Distance:     measurement.DefaultMeasurement(),
        MovingDuration: 0,
		ElapsedDuration:     0,
		Vertical:     measurement.DefaultMeasurement(),
	}
}

// Validate a completed object.
//
// Validation ensures that uuid, distance, vertical, duration and units are valid.
func (c *Completed) Validate() error {
	if c.ActivityUuid == "" {
		return fmt.Errorf("completed activity uuid is invalid")
	}
	err := c.Distance.Validate()
	if err != nil {
		return fmt.Errorf("completed distance is invalid: %w", err)
	}
	if c.ElapsedDuration == -1 {
		return fmt.Errorf("completed elapsed duration is invalid")
	}
    if c.MovingDuration == -1 {
		return fmt.Errorf("completed moving duration is invalid")
	}
	err2 := c.Vertical.Validate()
	if err2 != nil {
		return fmt.Errorf("completed vertical is invalid: %w", err2)
	}
	return nil
}

// Calculate the pace of the completed
func (c *Completed) CalculatePace(paceUnit measurement.PaceUnit) {
	pace := measurement.CalculatePace(c.Distance, c.MovingDuration, paceUnit)
	c.Pace = pace.String()
}

// Return true if all values of the completed are zero
func (c *Completed) IsZero() bool {
	return c.Distance.Length == 0 && c.MovingDuration == 0 && c.Vertical.Length == 0 && c.ElapsedDuration == 0
}
