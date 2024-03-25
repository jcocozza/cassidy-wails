package dateutil

import (
	"fmt"
	"reflect"
	"testing"

	gostructstringify "github.com/jcocozza/go_struct_stringify"
)

func TestGetDateMicrocycle(t *testing.T) {
	type args struct {
		date                     string
		microcycleStartDayOfWeek string
		microcycleLength         int
	}
	tests := []struct {
		name string
		args args
		want []*DateObject
	}{
		{"case1", args{"2024-01-01", Monday, 1}, []*DateObject{{DayOfWeek: Monday, Date: "2024-01-01"}}},
		{"case2", args{"2024-01-01", Monday, 2}, []*DateObject{{DayOfWeek: Monday, Date: "2024-01-01"}, {DayOfWeek: Tuesday, Date: "2024-01-02"}}},
		{"case3", args{"2024-01-05", Sunday, 7}, []*DateObject{
			{DayOfWeek: Sunday, Date: "2023-12-31"},
			{DayOfWeek: Monday, Date: "2024-01-01"},
			{DayOfWeek: Tuesday, Date: "2024-01-02"},
			{DayOfWeek: Wednesday, Date: "2024-01-03"},
			{DayOfWeek: Thursday, Date: "2024-01-04"},
			{DayOfWeek: Friday, Date: "2024-01-05"},
			{DayOfWeek: Saturday, Date: "2024-01-06"},
		}},
		{"case4", args{"2024-01-06", Sunday, 7}, []*DateObject{
			{DayOfWeek: Sunday, Date: "2023-12-31"},
			{DayOfWeek: Monday, Date: "2024-01-01"},
			{DayOfWeek: Tuesday, Date: "2024-01-02"},
			{DayOfWeek: Wednesday, Date: "2024-01-03"},
			{DayOfWeek: Thursday, Date: "2024-01-04"},
			{DayOfWeek: Friday, Date: "2024-01-05"},
			{DayOfWeek: Saturday, Date: "2024-01-06"},
		}},
		{"case5", args{"2024-02-26", Monday, 7}, []*DateObject{
			{DayOfWeek: Monday, Date: "2024-02-26"},
			{DayOfWeek: Tuesday, Date: "2024-02-27"},
			{DayOfWeek: Wednesday, Date: "2024-02-28"},
			{DayOfWeek: Thursday, Date: "2024-02-29"},
			{DayOfWeek: Friday, Date: "2024-03-01"},
			{DayOfWeek: Saturday, Date: "2024-03-02"},
			{DayOfWeek: Sunday, Date: "2024-03-03"},
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
		startDate string
		endDate   string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
	}{
		{"case1", args{"2024-01-01", "2024-01-02"}, "2024-01-03", "2024-01-04"},
		{"case2", args{"2024-01-01", "2024-01-06"}, "2024-01-07", "2024-01-12"},
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
		startDate      string
		endDate        string
		numberOfCycles int
	}
	tests := []struct {
		name    string
		args    args
		want    [][]*DateObject
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetNextNCycles(tt.args.startDate, tt.args.endDate, tt.args.numberOfCycles)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNextNCycles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetNextNCycles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPreviousCycle(t *testing.T) {
	type args struct {
		startDate string
		endDate   string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
	}{
		{"case1", args{"2024-01-01", "2024-01-02"}, "2023-12-30", "2023-12-31"},
		{"case2", args{"2024-01-07", "2024-01-12"}, "2024-01-01", "2024-01-06"},
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
		startDate      string
		endDate        string
		numberOfCycles int
	}
	tests := []struct {
		name    string
		args    args
		want    [][]*DateObject
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetPreviousNCycles(tt.args.startDate, tt.args.endDate, tt.args.numberOfCycles)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPreviousNCycles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPreviousNCycles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetNextPrevious(t *testing.T) {
	type args struct {
		startDate string
		endDate   string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   string
		want2   string
		want3   string
		wantErr bool
	}{
		{"case1", args{"2024-01-01", "2024-01-02"}, "2024-01-03", "2024-01-04", "2023-12-30", "2023-12-31", false},
		{"case2", args{"asdf-01", "2024-01-02"}, "", "", "", "", true},
		{"case3", args{"2024-01-01", "asdf2"}, "", "", "", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, got3, err := GetNextPrevious(tt.args.startDate, tt.args.endDate)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNextPrevious() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
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
		initialDate      string
		microcycleLength int
	}
	tests := []struct {
		name string
		args args
		want []*DateObject
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
