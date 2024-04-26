package dateutil
/*
import (
	"fmt"
	"time"
)

type DateObject struct {
	DayOfWeek string `json:"day_of_week"` // Monday, Tuesday, ...
	Date      string `json:"date"`		  // YYYY-MM-DD
}
// An empty date object has no day of week and no date.
func EmptyDateObject() *DateObject {
	return &DateObject{
		DayOfWeek: "",
		Date: "",
	}
}
// Create a date object from a date string.
func CreateFromDate(date string) (*DateObject, error) {
	t, err := time.Parse(Layout, date)
	if err != nil {
		return nil, fmt.Errorf("failed to create from date: %w", err)
	}
	return &DateObject{
		DayOfWeek: t.Weekday().String(),
		Date: date,
	}, nil
}
*/