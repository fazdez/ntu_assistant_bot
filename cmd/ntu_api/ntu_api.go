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

//GetCurrentWeek returns the current academic week based on the start date.
//-1 with the appropriate error is returned when start_date is not initialized.
//Use UpdateStartDate() to initialize start_date.
func (ntu *NTUApi) GetCurrentWeek(currentDay time.Time) (int, error) {
	err := validateDate(ntu.startDate, time.Now())
	if err != nil {
		return -1, err
	}

	//to add logic here.
	return 0, nil
}

func (ntu *NTUApi) UpdateStartDate(startDate time.Time) {
	ntu.startDate = startDate
}
