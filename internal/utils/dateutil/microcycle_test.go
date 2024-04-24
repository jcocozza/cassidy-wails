package dateutil

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	gostructstringify "github.com/jcocozza/go_struct_stringify"
)

func TestGetDateMicrocycle(t *testing.T) {
	type args struct {
		date                     time.Time
		microcycleStartDayOfWeek string
		microcycleLength         int
	}
	tests := []struct {
		name string
		args args
		want []time.Time
	}{
		{"case1", args{time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC), Monday, 1}, []time.Time{
			time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)},
		},
		{"case2", args{time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC), Monday, 2}, []time.Time{
			time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2024, time.January, 2, 0, 0, 0, 0, time.UTC),
		}},
		{"case3", args{time.Date(2024, time.January, 5, 0, 0, 0, 0, time.UTC), Sunday, 7}, []time.Time{
			time.Date(2023, time.December, 31, 0, 0, 0, 0, time.UTC),
			time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2024, time.January, 2, 0, 0, 0, 0, time.UTC),
			time.Date(2024, time.January, 3, 0, 0, 0, 0, time.UTC),
			time.Date(2024, time.January, 4, 0, 0, 0, 0, time.UTC),
			time.Date(2024, time.January, 5, 0, 0, 0, 0, time.UTC),
			time.Date(2024, time.January, 6, 0, 0, 0, 0, time.UTC),
		}},
		{"case4", args{time.Date(2024, time.January, 6, 0, 0, 0, 0, time.UTC), Sunday, 7}, []time.Time{
			time.Date(2023, time.December, 31, 0, 0, 0, 0, time.UTC),
			time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2024, time.January, 2, 0, 0, 0, 0, time.UTC),
			time.Date(2024, time.January, 3, 0, 0, 0, 0, time.UTC),
			time.Date(2024, time.January, 4, 0, 0, 0, 0, time.UTC),
			time.Date(2024, time.January, 5, 0, 0, 0, 0, time.UTC),
			time.Date(2024, time.January, 6, 0, 0, 0, 0, time.UTC),
		}},
		{"case5", args{time.Date(2024, time.February, 26, 0, 0, 0, 0, time.UTC), Monday, 7}, []time.Time{
			time.Date(2024, time.February, 26, 0, 0, 0, 0, time.UTC),
			time.Date(2024, time.February, 27, 0, 0, 0, 0, time.UTC),
			time.Date(2024, time.February, 28, 0, 0, 0, 0, time.UTC),
			time.Date(2024, time.February, 29, 0, 0, 0, 0, time.UTC),
			time.Date(2024, time.March, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2024, time.March, 2, 0, 0, 0, 0, time.UTC),
			time.Date(2024, time.March, 3, 0, 0, 0, 0, time.UTC),
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDateMicrocycle(tt.args.date, tt.args.microcycleStartDayOfWeek, tt.args.microcycleLength); !reflect.DeepEqual(got, tt.want) {
				fmt.Println(gostructstringify.StructStringify(got))
				t.Errorf("GetDateMicrocycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestGetNextCycle(t *testing.T) {
	type args struct {
		startDate time.Time
		endDate   time.Time
	}
	tests := []struct {
		name  string
		args  args
		want  time.Time
		want1 time.Time
	}{
		{"case1", args{time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC), time.Date(2024, time.January, 2, 0, 0, 0, 0, time.UTC)}, time.Date(2024, time.January, 3, 0, 0, 0, 0, time.UTC), time.Date(2024, time.January, 4, 0, 0, 0, 0, time.UTC)},
		{"case2", args{time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC), time.Date(2024, time.January, 6, 0, 0, 0, 0, time.UTC)}, time.Date(2024, time.January, 7, 0, 0, 0, 0, time.UTC), time.Date(2024, time.January, 12, 0, 0, 0, 0, time.UTC)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GetNextCycle(tt.args.startDate, tt.args.endDate)
			if got != tt.want {
				t.Errorf("GetNextCycle() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetNextCycle() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

// TODO
func TestGetNextNCycles(t *testing.T) {
	type args struct {
		startDate      time.Time
		endDate        time.Time
		numberOfCycles int
	}
	tests := []struct {
		name    string
		args    args
		want    [][]time.Time
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetNextNCycles(tt.args.startDate, tt.args.endDate, tt.args.numberOfCycles)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetNextNCycles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPreviousCycle(t *testing.T) {
	type args struct {
		startDate time.Time
		endDate   time.Time
	}
	tests := []struct {
		name  string
		args  args
		want  time.Time
		want1 time.Time
	}{
		{"case1", args{time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC), time.Date(2024, time.January, 2, 0, 0, 0, 0, time.UTC)}, time.Date(2023, time.December, 30, 0, 0, 0, 0, time.UTC), time.Date(2023, time.December, 31, 0, 0, 0, 0, time.UTC)},
		{"case2", args{time.Date(2024, time.January, 7, 0, 0, 0, 0, time.UTC), time.Date(2024, time.January, 12, 0, 0, 0, 0, time.UTC)}, time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC), time.Date(2024, time.January, 6, 0, 0, 0, 0, time.UTC)},
		{"case3", args{time.Date(2024, time.March, 4, 0, 0, 0, 0, time.UTC), time.Date(2024, time.March, 10, 0, 0, 0, 0, time.UTC)}, time.Date(2024, time.February, 26, 0, 0, 0, 0, time.UTC), time.Date(2024, time.March, 3, 0, 0, 0, 0, time.UTC)},
		{"case3", args{time.Date(2024, time.March, 4, 21, 34, 0, 0, time.Local), time.Date(2024, time.March, 10, 21, 34, 0, 0, time.Local)}, time.Date(2024, time.February, 26, 21, 34, 0, 0, time.Local), time.Date(2024, time.March, 3, 21, 34, 0, 0, time.Local)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GetPreviousCycle(tt.args.startDate, tt.args.endDate)
			if got != tt.want {
				t.Errorf("GetPreviousCycle() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetPreviousCycle() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

// TODO
func TestGetPreviousNCycles(t *testing.T) {
	type args struct {
		startDate      time.Time
		endDate        time.Time
		numberOfCycles int
	}
	tests := []struct {
		name    string
		args    args
		want    [][]time.Time
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetPreviousNCycles(tt.args.startDate, tt.args.endDate, tt.args.numberOfCycles)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPreviousNCycles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetNextPrevious(t *testing.T) {
	type args struct {
		startDate time.Time
		endDate   time.Time
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		want1   time.Time
		want2   time.Time
		want3   time.Time
	}{
		{"case1", args{time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC), time.Date(2024, time.January, 2, 0, 0, 0, 0, time.UTC)}, time.Date(2024, time.January, 3, 0, 0, 0, 0, time.UTC), time.Date(2024, time.January, 4, 0, 0, 0, 0, time.UTC), time.Date(2023, time.December, 30, 0, 0, 0, 0, time.UTC), time.Date(2023, time.December, 31, 0, 0, 0, 0, time.UTC)},
		{"case1", args{time.Date(2024, time.March, 4, 0, 0, 0, 0, time.UTC), time.Date(2024, time.March, 10, 0, 0, 0, 0, time.UTC)}, time.Date(2024, time.March, 11, 0, 0, 0, 0, time.UTC), time.Date(2024, time.March, 17, 0, 0, 0, 0, time.UTC), time.Date(2024, time.February, 26, 0, 0, 0, 0, time.UTC), time.Date(2024, time.March, 3, 0, 0, 0, 0, time.UTC)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, got3 := GetNextPrevious(tt.args.startDate, tt.args.endDate)
			if got != tt.want {
				t.Errorf("GetNextPrevious() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetNextPrevious() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("GetNextPrevious() got2 = %v, want %v", got2, tt.want2)
			}
			if got3 != tt.want3 {
				t.Errorf("GetNextPrevious() got3 = %v, want %v", got3, tt.want3)
			}
		})
	}
}
// the 'correct' answer to this function's output will change since it is based on the current date
// so I'm not writing any tests for now.
func TestGetCurrentCycleFromInitialDate(t *testing.T) {
	type args struct {
		initialDate      time.Time
		microcycleLength int
	}
	tests := []struct {
		name string
		args args
		want []time.Time
	}{
		//{"test", args{initialDate: "2024-02-26", microcycleLength: 10}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			g := GetCurrentCycleFromInitialDate(tt.args.initialDate, tt.args.microcycleLength);

			println(gostructstringify.StructStringify(g))

			if got := GetCurrentCycleFromInitialDate(tt.args.initialDate, tt.args.microcycleLength); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCurrentCycleFromInitialDate() = %v, want %v", got, tt.want)
			}
		})
	}
}
