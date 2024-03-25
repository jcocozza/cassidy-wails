package measurement

import (
	"testing"
)

func TestValidateUnit(t *testing.T) {
	type args struct {
		unit Unit
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateUnit(tt.args.unit); (err != nil) != tt.wantErr {
				t.Errorf("ValidateUnit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateUnitClass(t *testing.T) {
	type args struct {
		unitClass UnitClass
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateUnitClass(tt.args.unitClass); (err != nil) != tt.wantErr {
				t.Errorf("ValidateUnitClass() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUnitClassControl(t *testing.T) {
	type args struct {
		unitClass UnitClass
		unitType  UnitType
	}
	tests := []struct {
		name string
		args args
		want Unit
	}{
		{"t1", args{unitClass: Imperial, unitType: Vertical}, Feet},
		{"t2", args{unitClass: Imperial, unitType: Distance}, Mile},
		{"t3", args{unitClass: Metric, unitType: Vertical}, Meter},
		{"t4", args{unitClass: Metric, unitType: Distance}, Kilometer},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UnitClassControl(tt.args.unitClass, tt.args.unitType); got != tt.want {
				t.Errorf("UnitClassControl() = %v, want %v", got, tt.want)
			}
		})
	}
}
