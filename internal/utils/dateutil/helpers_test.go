package dateutil

import (
	"reflect"
	"testing"
)

func TestGenerateDateRange(t *testing.T) {
	type args struct {
		startDateStr string
		endDateStr   string
	}
	tests := []struct {
		name    string
		args    args
		want    []*DateObject
		wantErr bool
	}{
		{"case1", args{"2024-01-01", "2024-01-02"}, []*DateObject{{Date: "2024-01-01", DayOfWeek: Monday}, {Date: "2024-01-02", DayOfWeek: Tuesday}}, false},
		{"case2", args{"2024-01-01", "2024-01-07"}, []*DateObject{
			{Date: "2024-01-01", DayOfWeek: Monday},
			{Date: "2024-01-02", DayOfWeek: Tuesday},
			{Date: "2024-01-03", DayOfWeek: Wednesday},
			{Date: "2024-01-04", DayOfWeek: Thursday},
			{Date: "2024-01-05", DayOfWeek: Friday},
			{Date: "2024-01-06", DayOfWeek: Saturday},
			{Date: "2024-01-07", DayOfWeek: Sunday}},
			false},
		{"case3", args{"asdf1asd", "2024-01-02"}, nil, true},
		{"case4", args{"2024-01-01", "asdf"}, nil, true},
		{"case5", args{"2024-0101", "2024-01-02"}, nil, true},
		{"case6", args{"2024-02-26", "2024-03-03"}, []*DateObject{
			{DayOfWeek: Monday, Date: "2024-02-26"},
			{DayOfWeek: Tuesday, Date: "2024-02-27"},
			{DayOfWeek: Wednesday, Date: "2024-02-28"},
			{DayOfWeek: Thursday, Date: "2024-02-29"},
			{DayOfWeek: Friday, Date: "2024-03-01"},
			{DayOfWeek: Saturday, Date: "2024-03-02"},
			{DayOfWeek: Sunday, Date: "2024-03-03"},
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateDateRange(tt.args.startDateStr, tt.args.endDateStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateDateRange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateDateRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestGetDayOfWeek(t *testing.T) {
	tests := []struct {
		name    string
		date    string
		want    string
		wantErr bool
	}{
		{"monday", "2024-01-01", Monday, false},
		{"tuesday", "2024-01-02", Tuesday, false},
		{"wednesday", "2024-01-03", Wednesday, false},
		{"thursday", "2024-01-04", Thursday, false},
		{"friday", "2024-01-05", Friday, false},
		{"saturday", "2024-01-06", Saturday, false},
		{"sunday", "2024-01-07", Sunday, false},
		{"throw err", "", "", true},
		{"throw err", "123", "", true},
		{"throw err", "asdf", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetDayOfWeek(tt.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDayOfWeek() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetDayOfWeek() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_daysDifference(t *testing.T) {
	type args struct {
		date1 string
		date2 string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"case1", args{"2024-01-01", "2024-01-01"}, 0, false},
		{"case2", args{"2024-01-01", "2024-01-07"}, 6, false},
		{"case3", args{"202-01-01", "2024-01-01"}, 0, true},
		{"case4", args{"2024-01-01", "2asdf1"}, 0, true},
		{"case 5", args{"2024-02-26", "2024-03-03"}, 6, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := daysDifference(tt.args.date1, tt.args.date2)
			if (err != nil) != tt.wantErr {
				t.Errorf("daysDifference() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("daysDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsFuture(t *testing.T) {
	type args struct {
		date string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{"test", args{date: "2024-01-01"}, false, false},
		{"test fail", args{date: "2024-01-0"}, false, true},
		{"test future", args{date: "2100-01-01"}, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsFuture(tt.args.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsFuture() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsFuture() = %v, want %v", got, tt.want)
			}
		})
	}
}
