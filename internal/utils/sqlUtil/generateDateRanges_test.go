package sqlutil

import (
	"testing"
	"time"
)

func TestgenerateDateRangesCTE(t *testing.T) {
	type args struct {
		dateRanges []*DateRange
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test", args{dateRanges: []*DateRange{
			NewDateRange(time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC), time.Date(2024, time.January, 7, 0, 0, 0, 0, time.UTC)),
			NewDateRange(time.Date(2024, time.January, 8, 0, 0, 0, 0, time.UTC), time.Date(2024, time.January, 14, 0, 0, 0, 0, time.UTC)),
			NewDateRange(time.Date(2024, time.January, 15, 0, 0, 0, 0, time.UTC), time.Date(2024, time.January, 21, 0, 0, 0, 0, time.UTC)),
			NewDateRange(time.Date(2024, time.January, 22, 0, 0, 0, 0, time.UTC), time.Date(2024, time.January, 28, 0, 0, 0, 0, time.UTC)),
		}}, "WITH DateRanges AS (\nSELECT '2024-01-01' AS start_date, '2024-01-07' AS end_date\nUNION ALL\nSELECT '2024-01-08' AS start_date, '2024-01-14' AS end_date\nUNION ALL\nSELECT '2024-01-15' AS start_date, '2024-01-21' AS end_date\nUNION ALL\nSELECT '2024-01-22' AS start_date, '2024-01-28' AS end_date\n)"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateDateRangesCTE(tt.args.dateRanges); got != tt.want {
				t.Errorf("GenerateDateRanges() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateDateRangesPreviousCTE(t *testing.T) {
	type args struct {
		startDate time.Time
		endDate   time.Time
		numTotals int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test", args{startDate: time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC), endDate: time.Date(2024, time.January, 7, 0, 0, 0, 0, time.UTC), numTotals: 4}, "WITH DateRanges AS (\nSELECT '2024-01-01' AS start_date, '2024-01-07' AS end_date\nUNION ALL\nSELECT '2023-12-25', '2023-12-31'\nUNION ALL\nSELECT '2023-12-18', '2023-12-24'\nUNION ALL\nSELECT '2023-12-11', '2023-12-17'\nUNION ALL\nSELECT '2023-12-04', '2023-12-10'\n)"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateDateRangesPreviousCTE(tt.args.startDate, tt.args.endDate, tt.args.numTotals); got != tt.want {
				t.Errorf("GenerateDateRangesPreviousCTE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateDateRangesNextCTE(t *testing.T) {
	type args struct {
		startDate time.Time
		endDate   time.Time
		numTotals int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test", args{startDate: time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC), endDate: time.Date(2024, time.January, 7, 0, 0, 0, 0, time.UTC), numTotals: 4}, "WITH DateRanges AS (\nSELECT '2024-01-01' AS start_date, '2024-01-07' AS end_date\nUNION ALL\nSELECT '2024-01-08', '2024-01-14'\nUNION ALL\nSELECT '2024-01-15', '2024-01-21'\nUNION ALL\nSELECT '2024-01-22', '2024-01-28'\nUNION ALL\nSELECT '2024-01-29', '2024-02-04'\n)"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateDateRangesNextCTE(tt.args.startDate, tt.args.endDate, tt.args.numTotals); got != tt.want {
				t.Errorf("GenerateDateRangesNextCTE() = %v, want %v", got, tt.want)
			}
		})
	}
}
