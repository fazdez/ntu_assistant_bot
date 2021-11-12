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
