package dateutil

import "time"

// take a start date and end date and produce a list of dates in between those two dates (inclusive of those two)
func GenerateDateRange(startDateStr, endDateStr string) ([]*DateObject, error) {
	startDate, err := time.Parse(Layout, startDateStr)
	if err != nil {
		return nil, err
	}

	endDate, err := time.Parse(Layout, endDateStr)
	if err != nil {
		return nil, err
	}

	dateRange := []*DateObject{}
	currentDate := startDate
	for !currentDate.After(endDate) {
		dto := &DateObject{DayOfWeek: currentDate.Weekday().String(), Date: currentDate.Format(Layout)}
		dateRange = append(dateRange, dto)
		currentDate = currentDate.AddDate(0, 0, 1)
	}

	return dateRange, nil
}
// return a startDate, endDate range that represents the numberPriors of the cycle
func GeneratePriorsRange(startDate string, endDate string, numberPriors int) (string, string) {
	var tmpStart = startDate
	var tmpEnd = endDate
	finalEnd := ""
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
func daysDifference(date1, date2 string) (int, error) {
	t1, err := time.Parse(Layout, date1)
	if err != nil {
		return 0, err
	}

	t2, err := time.Parse(Layout, date2)
	if err != nil {
		return 0, err
	}

	duration := t2.Sub(t1)
	days := int(duration.Hours() / 24)

	return days, nil
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