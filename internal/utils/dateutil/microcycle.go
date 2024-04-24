package dateutil

import "time"

// get the microcycle for a given date and start day of week combination.
//
// e.g. Get the microcycle that contains a date, "YYYY-MM-DD" that starts on weekday e.g. Monday
func GetDateMicrocycle(date time.Time, microcycleStartDayOfWeek string, microcycleLength int) []time.Time {
	// Get the day of the week
	dayOfWeek := date.Weekday().String()

	if microcycleStartDayOfWeek == dayOfWeek {
		lastDayMicrocycle := date.AddDate(0, 0, microcycleLength - 1)
		mc := GenerateDateRange(date, lastDayMicrocycle)
		return mc
	} else {
		firstDayMicrocycle := date
		for microcycleStartDayOfWeek != dayOfWeek {
			firstDayMicrocycle = firstDayMicrocycle.AddDate(0, 0, -1) // there will never be a day where the beginning of the microcycle is after that day
			dayOfWeek = firstDayMicrocycle.Weekday().String()
		}

		lastDayMicrocycle := firstDayMicrocycle.AddDate(0, 0, microcycleLength - 1)
		mc := GenerateDateRange(firstDayMicrocycle, lastDayMicrocycle)
		return mc
	}
}
// Get the current microcyle based on a length and start date
//
// e.g. Get the microcyle that contains the current date starting on the passed weekday
//
// (Untested)
func CurrentMicrocycle(microcycleStartDayOfWeek string, microcycleLength int) []time.Time {
	currentDate := time.Now()
	return GetDateMicrocycle(currentDate, microcycleStartDayOfWeek, microcycleLength)
}
// Get the current microcycle for a given length and initial date combination
func GetCurrentCycleFromInitialDate(initialDate time.Time, microcycleLength int) []time.Time {
	// Calculate the difference in days from the initial date to today
	daysDifference := time.Since(initialDate).Hours() / 24

	// Calculate number of microcycles away from today
	numCyclesAway := int(daysDifference) / microcycleLength

	// Calculate the start date of the current microcycle
	startDate := initialDate.Add(time.Duration(numCyclesAway * microcycleLength) * 24 * time.Hour)

	return GetDateMicrocycle(startDate, startDate.Weekday().String(), microcycleLength)
}
// return the next cycle start/end for the passed start and end
func GetNextCycle(startDate time.Time, endDate time.Time) (time.Time, time.Time) {
	cycleLength := daysDifference(startDate, endDate)

	newSd := startDate.AddDate(0, 0, cycleLength+1)
	newEd := endDate.AddDate(0, 0, cycleLength+1)

	return newSd, newEd
}
// ! Currently not used
// from a given (startDate, endDate) cycle, return the next N cycles (from that cycle)
func GetNextNCycles(startDate time.Time, endDate time.Time, numberOfCycles int) [][]time.Time {
	cycleLength := daysDifference(startDate, endDate)
	var cycleList [][]time.Time
	var tmpSd time.Time = startDate
	var tmpEd time.Time = endDate
	for i := 0; i < numberOfCycles; i++ {
		tmpSd = tmpSd.AddDate(0, 0, cycleLength+1)
		tmpEd = tmpEd.AddDate(0, 0, cycleLength+1)
		newCycleRange := GenerateDateRange(tmpSd, tmpEd)
		cycleList = append(cycleList, newCycleRange)
	}
	return cycleList
}
// return the previous cycle start/end for the passed start and end
func GetPreviousCycle(startDate time.Time, endDate time.Time) (time.Time, time.Time) {
	cycleLength := daysDifference(startDate, endDate)
	newSd := startDate.AddDate(0, 0, -cycleLength-1).In(time.UTC)
	newEd := endDate.AddDate(0, 0, -cycleLength-1).In(time.UTC)
	return newSd, newEd
}
// from a given (startDate, endDate) cycle, return the previous N cycles prior to that cycle
func GetPreviousNCycles(startDate time.Time, endDate time.Time, numberOfCycles int) [][]time.Time {
	cycleLength := daysDifference(startDate, endDate)
	var cycleList [][]time.Time
	var tmpSd time.Time = startDate
	var tmpEd time.Time = endDate
	for i := 0; i < numberOfCycles; i++ {
		tmpSd = tmpSd.AddDate(0, 0, -cycleLength-1)
		tmpEd = tmpEd.AddDate(0, 0, -cycleLength-1)
		newCycleRange := GenerateDateRange(tmpSd, tmpEd)
		cycleList = append(cycleList, newCycleRange)
	}
	return cycleList
}

// from a given (startDate, endDate) cycle, return the next cycle start/end pair and the previous start/end pair
func GetNextPrevious(startDate time.Time, endDate time.Time) (time.Time, time.Time, time.Time, time.Time) {
	cycleLength := daysDifference(startDate, endDate)
	if cycleLength == 0 {
		cycleLength = 1
	}
	sdNext := startDate.AddDate(0, 0, cycleLength+1)
	edNext := endDate.AddDate(0, 0, cycleLength+1)
	sdPrevious := startDate.AddDate(0, 0, -cycleLength-1)
	edPrevious := endDate.AddDate(0, 0, -cycleLength-1)
	return sdNext, edNext, sdPrevious, edPrevious
}