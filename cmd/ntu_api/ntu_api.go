package ntuapi

import (
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

//GetWeek returns the current academic week based on the start date and target date
//-1 with the appropriate error if validation check fails
//Use UpdateStartDate() to initialize start_date.
func (ntu *NTUApi) GetWeek(targetDate time.Time) (int, error) {
	err := validateDate(ntu.startDate, targetDate)
	if err != nil {
		return -1, err
	}

	return calculateWeek(ntu.startDate, targetDate), nil
}

//GetCurrentWeek calls GetWeek() with current date.
func (ntu *NTUApi) GetCurrentWeek() (int, error) {
	return ntu.GetWeek(time.Now())
}

func (ntu *NTUApi) UpdateStartDate(startDate time.Time) {
	ntu.startDate = startDate
}
