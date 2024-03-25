package dateutil

import (
	"fmt"
	"time"
)
// Ensure that a day of week is in Monday through Sunday
func ValidateDayOfWeek(dow string) error {
	if dow == "" {
		return fmt.Errorf("day of week cannot be empty")
	}
	if dow != Monday && dow != Tuesday && dow != Wednesday && dow != Thursday && dow != Friday && dow != Saturday && dow != Sunday {
		return fmt.Errorf("day of week is not in known days of week")
	}
	return nil
}
// Ensure that a date is parseable for the Layout 2006-01-02
func ValidateDate(date string) error {
	_, err := time.Parse(Layout, date)
	if err != nil {
		return fmt.Errorf("date failed to validate: %w", err)
	}
	return nil
}