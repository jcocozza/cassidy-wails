package model

import (
	"testing"

	"github.com/jcocozza/cassidy-wails/internal/utils/measurement"
)

func TestCompleted_Validate(t *testing.T) {
	type fields struct {
		ActivityUuid string
		Distance     *measurement.Measurement
		Duration     float64
		Vertical     *measurement.Measurement
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"valid", fields{ActivityUuid: "c2cb0cd2-cac8-11ee-b5e8-325096b39f47", Distance: measurement.CreateMeasurement(measurement.Mile, 10), Duration: 5000, Vertical: measurement.CreateMeasurement(measurement.Feet, 150)}, false},
		{"invalid uuid", fields{ActivityUuid: "", Distance: measurement.CreateMeasurement(measurement.Mile, 10), Duration: 5000, Vertical: measurement.CreateMeasurement(measurement.Feet, 150)}, true},
		{"invalid distance length", fields{ActivityUuid: "c2cb0cd2-cac8-11ee-b5e8-325096b39f47", Distance: measurement.CreateMeasurement(measurement.Mile, -1), Duration: 5000, Vertical: measurement.CreateMeasurement(measurement.Feet, 150)}, true},
		{"invalid distance unit", fields{ActivityUuid: "c2cb0cd2-cac8-11ee-b5e8-325096b39f47", Distance: measurement.CreateMeasurement("asdf", 10), Duration: 5000, Vertical: measurement.CreateMeasurement(measurement.Feet, 150)}, true},
		{"invalid duration", fields{ActivityUuid: "c2cb0cd2-cac8-11ee-b5e8-325096b39f47", Distance: measurement.CreateMeasurement(measurement.Mile, 10), Duration: -1, Vertical: measurement.CreateMeasurement(measurement.Feet, 150)}, true},
		{"invalid vertical length", fields{ActivityUuid: "c2cb0cd2-cac8-11ee-b5e8-325096b39f47", Distance: measurement.CreateMeasurement(measurement.Mile, 10), Duration: 5000, Vertical: measurement.CreateMeasurement(measurement.Feet, -1)}, true},
		{"invalid vertical unit", fields{ActivityUuid: "c2cb0cd2-cac8-11ee-b5e8-325096b39f47", Distance: measurement.CreateMeasurement(measurement.Mile, 10), Duration: 5000, Vertical: measurement.CreateMeasurement("mph", 150)}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Completed{
				ActivityUuid: tt.fields.ActivityUuid,
				Distance:     tt.fields.Distance,
				Duration:     tt.fields.Duration,
				Vertical:     tt.fields.Vertical,
			}
			if err := c.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Completed.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
