package measurement

import (
	"reflect"
	"testing"
)

func TestCalculatePace(t *testing.T) {
	type args struct {
		distance *Measurement
		duration float64
		paceUnit PaceUnit
	}
	tests := []struct {
		name string
		args args
		want *Pace
	}{
		{"test", args{distance: CreateMeasurement(Mile, 1609.344), duration: 360, paceUnit: MinutesPerMile}, &Pace{Unit: MinutesPerMile, Speed: 6.00}},
		{"test", args{distance: CreateMeasurement(Mile, 1609.344*2), duration: 360 * 2, paceUnit: MinutesPerMile}, &Pace{Unit: MinutesPerMile, Speed: 6.00}},
		{"test", args{distance: CreateMeasurement(Mile, 1609.344), duration: 360, paceUnit: MilesPerHour}, &Pace{Unit: MilesPerHour, Speed: 10.00}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculatePace(tt.args.distance, tt.args.duration, tt.args.paceUnit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalculatePace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPace_String(t *testing.T) {
	type fields struct {
		Unit  PaceUnit
		Speed float64
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"test", fields{Unit: MinutesPerMile, Speed: 6.5}, "6:30 min/mile"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Pace{
				Unit:  tt.fields.Unit,
				Speed: tt.fields.Speed,
			}
			if got := p.String(); got != tt.want {
				t.Errorf("Pace.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
