package measurement

import (
	"math"
	"reflect"
	"testing"
)

func TestDifferenceSI(t *testing.T) {
	type args struct {
		m1 *Measurement
		m2 *Measurement
	}
	tests := []struct {
		name    string
		args    args
		want    *Measurement
		wantErr bool
	}{
		{"t1", args{m1: &Measurement{Length: 1, Unit: Mile}, m2: &Measurement{Length: 2, Unit: Mile}}, &Measurement{Length: -1, Unit: Mile}, false},
		{"t2", args{m1: &Measurement{Length: 1, Unit: Mile}, m2: &Measurement{Length: 2, Unit: Feet}}, nil, true},
		{"t3", args{m1: &Measurement{Length: 10, Unit: Mile}, m2: &Measurement{Length: 4, Unit: Mile}}, &Measurement{Length: 6, Unit: Mile}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DifferenceSI(tt.args.m1, tt.args.m2)
			if (err != nil) != tt.wantErr {
				t.Errorf("DifferenceSI() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DifferenceSI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPercentChange(t *testing.T) {
	type args struct {
		m1 *Measurement
		m2 *Measurement
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{"t1", args{m1: &Measurement{Length: 1, Unit: Mile}, m2: &Measurement{Length: 2, Unit: Mile}}, -50, false},
		{"t2", args{m1: &Measurement{Length: 2, Unit: Mile}, m2: &Measurement{Length: 1, Unit: Mile}}, 100, false},
		{"t3", args{m1: &Measurement{Length: 0, Unit: Mile}, m2: &Measurement{Length: 2, Unit: Mile}}, -100, false},
		{"t4", args{m1: &Measurement{Length: 2, Unit: Feet}, m2: &Measurement{Length: 0, Unit: Mile}}, math.NaN(), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PercentChange(tt.args.m1, tt.args.m2)

			if math.IsNaN(got) && math.IsNaN(tt.want) {
				return
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("PercentChange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PercentChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
