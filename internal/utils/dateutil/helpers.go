package dateutil

import "time"

// take a start date and end date and produce a list of dates in between those two dates (inclusive of those two)
func GenerateDateRange(startDate, endDate time.Time) ([]time.Time) {
	dateRange := []time.Time{}
	currentDate := startDate
	for !currentDate.After(endDate) {
		dateRange = append(dateRange, currentDate)
		currentDate = currentDate.AddDate(0, 0, 1)
	}
	return dateRange
}
// return a startDate, endDate range that represents the numberPriors of the cycle
func GeneratePriorsRange(startDate time.Time, endDate time.Time, numberPriors int) (time.Time, time.Time) {
	var tmpStart = startDate
	var tmpEnd = endDate
	finalEnd := time.Time{}
	for i := 0; i < numberPriors; i ++ {
		tmpStart, tmpEnd = GetPreviousCycle(tmpStart, tmpEnd)
		if i == 0 {
			finalEnd = tmpEnd
		}
	}
	return tmpStart, finalEnd
}
// Take in a date of form YYYY-MM-DD
// return the day of the week
func GetDayOfWeek(date string) (string, error) {
	// Parse the input date
	parsedDate, err := time.Parse(Layout, date)
	if err != nil {
		return "", err
	}

	// Get the day of the week
	dayOfWeek := parsedDate.Weekday().String()

	return dayOfWeek, nil
}
// calculate # of days between two dates
func daysDifference(date1, date2 time.Time) int {
    date1 = date1.Truncate(24 * time.Hour)
    date2 = date2.Truncate(24 * time.Hour)
    duration := date2.Sub(date1)
    days := int(duration.Hours() / 24)
    return days
}
// check if a date is future or not
func IsFuture(date string) (bool, error) {
	t1, err := time.Parse(Layout, date)
	if err != nil {
		return false, err
	}

	now := time.Now()

	return t1.After(now), nil
}
// Check if two time.Time structs occur on the same date
func SameDate(t1, t2 time.Time) bool {
	return t1.Year() == t2.Year() && t1.Month() == t2.Month() && t1.Day() == t2.Day()
}