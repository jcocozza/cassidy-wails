package utils

import (
	"math"
	"testing"
)

func TestPercentChange(t *testing.T) {
	type args struct {
		v1 float64
		v2 float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"t1", args{1,2}, -50},
		{"t2", args{2,1}, 100},
		{"t3", args{2,0}, math.NaN()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := PercentChange(tt.args.v1, tt.args.v2)

			if math.IsNaN(got) && math.IsNaN(tt.want) {
				return
			}

			if got != tt.want {
				t.Errorf("PercentChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
