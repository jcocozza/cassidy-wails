package dateutil

import "time"

// get the microcycle for a given date and start day of week combination.
//
// e.g. Get the microcycle that contains a date, "YYYY-MM-DD" that starts on weekday e.g. Monday
func GetDateMicrocycle(date string, microcycleStartDayOfWeek string, microcycleLength int) []*DateObject {
	parsedDate, err := time.Parse(Layout, date)
	if err != nil {
		panic(err)
	}
	// Get the day of the week
	dayOfWeek := parsedDate.Weekday().String()

	if microcycleStartDayOfWeek == dayOfWeek {
		lastDayMicrocycle := parsedDate.AddDate(0, 0, microcycleLength - 1)
		mc, _ := GenerateDateRange(date, lastDayMicrocycle.Format(Layout))
		return mc
	} else {
		firstDayMicrocycle := parsedDate
		for microcycleStartDayOfWeek != dayOfWeek {
			firstDayMicrocycle = firstDayMicrocycle.AddDate(0, 0, -1) // there will never be a day where the beginning of the microcycle is after that day
			dayOfWeek = firstDayMicrocycle.Weekday().String()
		}

		lastDayMicrocycle := firstDayMicrocycle.AddDate(0, 0, microcycleLength - 1)
		mc, _ := GenerateDateRange(firstDayMicrocycle.Format(Layout), lastDayMicrocycle.Format(Layout))
		return mc
	}
}
// Get the current microcyle based on a length and start date
//
// e.g. Get the microcyle that contains the current date starting on the passed weekday
//
// (Untested)
func CurrentMicrocycle(microcycleStartDayOfWeek string, microcycleLength int) []*DateObject {
	currentDate := time.Now().Format(Layout)
	return GetDateMicrocycle(currentDate, microcycleStartDayOfWeek, microcycleLength)
}
// Get the current microcycle for a given length and initial date combination
func GetCurrentCycleFromInitialDate(initialDate string, microcycleLength int) []*DateObject {
	parsedDate, err := time.Parse(Layout, initialDate)
	if err != nil {
		panic(err)
	}

	// Calculate the difference in days from the initial date to today
	daysDifference := time.Since(parsedDate).Hours() / 24

	// Calculate number of microcycles away from today
	numCyclesAway := int(daysDifference) / microcycleLength

	// Calculate the start date of the current microcycle
	startDate := parsedDate.Add(time.Duration(numCyclesAway * microcycleLength) * 24 * time.Hour)

	return GetDateMicrocycle(startDate.Format(Layout), startDate.Weekday().String(), microcycleLength)
}

// return the next cycle start/end for the passed start and end
func GetNextCycle(startDate string, endDate string) (string, string) {
	cycleLength, err := daysDifference(startDate, endDate)
	if err != nil {
		panic(err)
	}

	sd, _ := time.Parse(Layout, startDate)
	ed, _ := time.Parse(Layout, endDate)

	newSd := sd.AddDate(0, 0, cycleLength+1)
	newEd := ed.AddDate(0, 0, cycleLength+1)

	return newSd.Format(Layout), newEd.Format(Layout)
}

// ! Currently not used
// from a given (startDate, endDate) cycle, return the next N cycles (from that cycle)
func GetNextNCycles(startDate string, endDate string, numberOfCycles int) ([][]*DateObject, error) {
	cycleLength, err := daysDifference(startDate, endDate)

	if err != nil {
		return [][]*DateObject{}, err
	}

	sd, _ := time.Parse(Layout, startDate)
	ed, _ := time.Parse(Layout, endDate)

	var cycleList [][]*DateObject
	var tmpSd time.Time = sd
	var tmpEd time.Time = ed
	for i := 0; i < numberOfCycles; i++ {
		tmpSd = tmpSd.AddDate(0, 0, cycleLength+1)
		tmpEd = tmpEd.AddDate(0, 0, cycleLength+1)
		newCycleRange, _ := GenerateDateRange(tmpSd.Format(Layout), tmpEd.Format(Layout))
		cycleList = append(cycleList, newCycleRange)
	}
	return cycleList, nil
}

// return the previous cycle start/end for the passed start and end
func GetPreviousCycle(startDate string, endDate string) (string, string) {
	cycleLength, err := daysDifference(startDate, endDate)
	if err != nil {
		panic(err)
	}

	sd, _ := time.Parse(Layout, startDate)
	ed, _ := time.Parse(Layout, endDate)

	newSd := sd.AddDate(0, 0, -cycleLength-1)
	newEd := ed.AddDate(0, 0, -cycleLength-1)

	return newSd.Format(Layout), newEd.Format(Layout)
}

// from a given (startDate, endDate) cycle, return the previous N cycles prior to that cycle
func GetPreviousNCycles(startDate string, endDate string, numberOfCycles int) ([][]*DateObject, error) {
	cycleLength, err := daysDifference(startDate, endDate)

	if err != nil {
		return [][]*DateObject{}, err
	}

	sd, _ := time.Parse(Layout, startDate)
	ed, _ := time.Parse(Layout, endDate)

	var cycleList [][]*DateObject
	var tmpSd time.Time = sd
	var tmpEd time.Time = ed
	for i := 0; i < numberOfCycles; i++ {
		tmpSd = tmpSd.AddDate(0, 0, -cycleLength-1)
		tmpEd = tmpEd.AddDate(0, 0, -cycleLength-1)
		newCycleRange, _ := GenerateDateRange(tmpSd.Format(Layout), tmpEd.Format(Layout))
		cycleList = append(cycleList, newCycleRange)
	}
	return cycleList, nil
}

// from a given (startDate, endDate) cycle, return the next cycle start/end pair and the previous start/end pair
func GetNextPrevious(startDate string, endDate string) (string, string, string, string, error) {
	cycleLength, err := daysDifference(startDate, endDate)
	if err != nil {
		return "","","","", err
	}

	if cycleLength == 0 {
		cycleLength = 1
	}

	sd, _ := time.Parse(Layout, startDate)
	ed, _ := time.Parse(Layout, endDate)

	sdNext := sd.AddDate(0, 0, cycleLength+1)
	edNext := ed.AddDate(0, 0, cycleLength+1)
	sdPrevious := sd.AddDate(0, 0, -cycleLength-1)
	edPrevious := ed.AddDate(0, 0, -cycleLength-1)

	return sdNext.Format(Layout), edNext.Format(Layout), sdPrevious.Format(Layout), edPrevious.Format(Layout), nil
}