package dateutil

import (
	"reflect"
	"testing"
	"time"
)

func TestGenerateDateRange(t *testing.T) {
	type args struct {
		startDateStr time.Time
		endDateStr   time.Time
	}
	tests := []struct {
		name    string
		args    args
		want    []time.Time
	}{
		{"case1", args{time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC), time.Date(2024, time.January, 2, 0, 0, 0, 0, time.UTC)}, []time.Time{time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC), time.Date(2024, time.January, 2, 0, 0, 0, 0, time.UTC)}},
		{"case2", args{time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC), time.Date(2024, time.January, 7, 0, 0, 0, 0, time.UTC)}, []time.Time{
			time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2024, time.January, 2, 0, 0, 0, 0, time.UTC),
			time.Date(2024, time.January, 3, 0, 0, 0, 0, time.UTC),
			time.Date(2024, time.January, 4, 0, 0, 0, 0, time.UTC),
			time.Date(2024, time.January, 5, 0, 0, 0, 0, time.UTC),
			time.Date(2024, time.January, 6, 0, 0, 0, 0, time.UTC),
			time.Date(2024, time.January, 7, 0, 0, 0, 0, time.UTC)}},
		{"case3", args{time.Date(2024, time.February, 26, 0, 0, 0, 0, time.UTC), time.Date(2024, time.March, 3, 0, 0, 0, 0, time.UTC)}, []time.Time{
			time.Date(2024, time.February, 26, 0, 0, 0, 0, time.UTC),
			time.Date(2024, time.February, 27, 0, 0, 0, 0, time.UTC),
			time.Date(2024, time.February, 28, 0, 0, 0, 0, time.UTC),
			time.Date(2024, time.February, 29, 0, 0, 0, 0, time.UTC),
			time.Date(2024, time.March, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2024, time.March, 2, 0, 0, 0, 0, time.UTC),
			time.Date(2024, time.March, 3, 0, 0, 0, 0, time.UTC)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateDateRange(tt.args.startDateStr, tt.args.endDateStr)
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
		date1 time.Time
		date2 time.Time
	}
	tests := []struct {
		name    string
		args    args
		want    int
	}{
		{"case1", args{time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC), time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)}, 0},
		{"case2", args{time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC), time.Date(2024, time.January, 7, 0, 0, 0, 0, time.UTC)}, 6},
		{"case 5", args{time.Date(2024, time.February, 26, 0, 0, 0, 0, time.UTC), time.Date(2024, time.March, 3, 0, 0, 0, 0, time.UTC)}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := daysDifference(tt.args.date1, tt.args.date2)
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
