package ntuapi

import (
	"fmt"
	"time"
)

type Config struct {
}

type NTUApi struct {
	config    Config
	startDate time.Time
	endDate   time.Time
}

func New(config Config) *NTUApi {
	return &NTUApi{config, time.Time{}, time.Time{}}
}

//GetWeek returns the current teaching week based on the start date and target date
//-1 with the appropriate error if validation check fails
// 0 represents recess week
//Use UpdateStartDate() to initialize start_date.
func (ntu *NTUApi) GetWeek(targetDate time.Time) (int, error) {
	err := validateDate(ntu.startDate, targetDate)
	if err != nil {
		return -1, err
	}

	week := calculateWeek(ntu.startDate, targetDate)
	return ntu.getTeachingWeek(week)
}

//GetCurrentWeek calls GetWeek() with current date.
func (ntu *NTUApi) GetCurrentWeek() (int, error) {
	return ntu.GetWeek(time.Now())
}

func (ntu *NTUApi) UpdateStartDate(startDate time.Time) {
	ntu.startDate = startDate
}

//getTeachingWeek returns the teaching week (between 1 to 13)
//0 represents recess week
//else -1, error is returned
func (ntu *NTUApi) getTeachingWeek(week int) (int, error) {
	if week < 1 || week > 14 {
		return -1, &ErrWeekOutOfRange{message: fmt.Sprintf("week (got: %d) should be between 1 and 14", week)}
	}

	if week < 8 {
		return week, nil
	}
	//recess week
	if week == 8 {
		return 0, nil
	}

	return week - 1, nil
}
