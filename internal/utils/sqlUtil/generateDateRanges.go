package sqlutil

import (
	"time"

	"github.com/jcocozza/cassidy-wails/internal/utils/dateutil"
)

type DateRange struct {
	Start time.Time
	End   time.Time
}
// Create a new date range
func NewDateRange(start, end time.Time) *DateRange {
	return &DateRange{
		Start: start,
		End:   end,
	}
}
/*
Generate a table of this form:

WITH DateRanges AS (

		SELECT '2024-01-01' AS start_date, '2024-01-08' AS end_date
	    UNION ALL
	    SELECT '2024-01-09', '2024-01-16'
	    UNION ALL
	    SELECT '2024-01-17', '2024-01-24'
	    UNION ALL
	    SELECT '2024-01-25', '2024-02-01'

)

The idea is to attach this to the front of quries that need to check lots of date ranges
*/
func generateDateRangesCTE(dateRanges []*DateRange) string {
	// initial
	sql := "WITH DateRanges AS (\n"
	sql += "SELECT '" + dateRanges[0].Start.Format(dateutil.Layout) + "' AS start_date, '" + dateRanges[0].End.Format(dateutil.Layout) + "' AS end_date\n"
	sql += "UNION ALL\n"

	for i := 1; i <= len(dateRanges)-1; i++ {
		sql += "SELECT '" + dateRanges[i].Start.Format(dateutil.Layout) + "', '" + dateRanges[i].End.Format(dateutil.Layout) + "'\n"
		if i != len(dateRanges)-1 {
			sql += "UNION ALL\n"
		}
	}
	sql += ")"
	return sql
}
// Generate a date range Common Table expression for a give start and end date and number of totals in the previous direction
//
// Will be INCLUSIVE of the passed startDate, endDate
func GenerateDateRangesPreviousCTE(startDate, endDate time.Time, numTotals int) string {
	drList := []*DateRange{}
	drList = append(drList, NewDateRange(startDate, endDate))

	var tmpStart = startDate
	var tmpEnd = endDate
	for i := 0; i < numTotals; i++ {
		tmpStart, tmpEnd = dateutil.GetPreviousCycle(tmpStart, tmpEnd)
		drList = append(drList, NewDateRange(tmpStart, tmpEnd))
	}

	return generateDateRangesCTE(drList)
}

// Generate a date range Common Table expression for a give start and end date and number of totals in the next direction
//
// Will be INCLUSIVE of the passed startDate, endDate
func GenerateDateRangesNextCTE(startDate, endDate time.Time, numTotals int) string {
	drList := []*DateRange{}
	drList = append(drList, NewDateRange(startDate, endDate))

	var tmpStart = startDate
	var tmpEnd = endDate
	for i := 0; i < numTotals; i++ {
		tmpStart, tmpEnd = dateutil.GetNextCycle(tmpStart, tmpEnd)
		drList = append(drList, NewDateRange(tmpStart, tmpEnd))
	}

	return generateDateRangesCTE(drList)
}
