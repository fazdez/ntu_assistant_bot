package ntuapi

import (
	"testing"
	"time"
)

func TestInvalidDate(t *testing.T) {
	t.Run("StartDateUninitialized", func(t *testing.T) {
		err := validateDate(time.Time{}, time.Now())

		switch err.(type) {
		case nil:
			t.Errorf("wanted error got nil")
		case *ErrDateRange:
			t.Errorf("wanted %T got %T", &ErrNullDate{}, err)
		}
	})

	t.Run("StartDateOutdated", func(t *testing.T) {
		err := validateDate(time.Now().AddDate(-1, 0, 0), time.Now())
		switch err.(type) {
		case nil:
			t.Errorf("wanted error got nil")
		case *ErrNullDate:
			t.Errorf("wanted %T got %T", &ErrDateRange{}, err)
		}
	})

	t.Run("DatesOutOfRange", func(t *testing.T) {
		tests := []struct {
			startDate  time.Time
			targetDate time.Time
		}{
			{time.Now(), time.Now().AddDate(1, 0, 0)},                    //target date in the future
			{time.Now().AddDate(0, 0, 5), time.Now()},                    //start date in the future
			{time.Now().AddDate(0, -1, 0), time.Now().AddDate(0, -2, 0)}, //target date before start date
		}

		for idx, test := range tests {
			err := validateDate(test.startDate, test.targetDate)
			if err == nil {
				t.Errorf("test case %d: wanted err got nil", idx+1)
			}
		}
	})
}

func TestCalculateWeek(t *testing.T) {
	sample := time.Date(2021, time.August, 9, 0, 0, 0, 0, time.Local) //this sample follows 21/22 Sem 1 start-date.

	tests := []struct {
		startDate  time.Time
		targetDate time.Time
		want       int
	}{
		{sample, sample.AddDate(0, 0, 7), 2},
		{sample, sample.AddDate(0, 0, 3), 1},
		{sample, sample.AddDate(0, 1, 0), 5},
		{sample, sample.AddDate(0, 3, 0), 14}, //should return 14, even though it is teaching week 13, due to week 7 being recess week (but not a teaching week)
	}

	for idx, test := range tests {
		got := calculateWeek(test.startDate, test.targetDate)
		if got != test.want {
			t.Errorf("test case %d: wanted week %d got week %d <start_date: %s> <target_date: %s>", idx, test.want, got,
				test.startDate.Format(TIMEFORMAT), test.targetDate.Format(TIMEFORMAT))
		}
	}
}
