package ntuapi

import (
	"fmt"
	"time"
)

const TIMEFORMAT = "January 2, 2006"

func validateDate(startDate, targetDate time.Time) error {
	var currentDate = time.Now()
	if currentDate.Before(targetDate) {
		return &DateRangeError{message: fmt.Sprintf("%s is in the future", targetDate.Format(TIMEFORMAT))}
	}

	if startDate.IsZero() {
		return &NullDateError{message: "start date is not initialized"}
	}

	if startDate.AddDate(0, 6, 0).Before(currentDate) {
		return &DateRangeError{message: "start date is too outdated"}
	}

	if targetDate.Before(startDate) {
		return &DateRangeError{message: "target date is before start date"}
	}

	return nil
}

//calculateWeek gets the week number of target date, if the start date's week number is 1
//this function is not tied to the context of NTU academic week. it may even return negative.
//no validation checks here.
func calculateWeek(startDate, targetDate time.Time) int {
	_, start_week := startDate.ISOWeek()
	_, target_week := targetDate.ISOWeek()

	return target_week - start_week + 1
}
