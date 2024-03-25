package measurement

import (
	"testing"
)

func TestRoundToHundredth(t *testing.T) {
	type args struct {
		num float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"test 1", args{10.5}, 10.5},
		{"test 2", args{10.4999}, 10.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RoundToHundredth(tt.args.num); got != tt.want {
				t.Errorf("RoundToHundredth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestModifyMeasurement(t *testing.T) {
	type args struct {
		measurement  *Measurement
		incomingUnit Unit
		outgoingUnit Unit
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"t 1", args{measurement: &Measurement{Unit: Feet, Length: 10}, incomingUnit: Feet, outgoingUnit: Meter}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ModifyMeasurement(tt.args.measurement, tt.args.incomingUnit, tt.args.outgoingUnit); (err != nil) != tt.wantErr {
				t.Errorf("ModifyMeasurement() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
