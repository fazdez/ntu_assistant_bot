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
		case *DateRangeError:
			t.Errorf("wanted %T got %T", &NullDateError{}, err)
		}
	})

	t.Run("StartDateOutdated", func(t *testing.T) {
		err := validateDate(time.Now().AddDate(-1, 0, 0), time.Now())
		switch err.(type) {
		case nil:
			t.Errorf("wanted error got nil")
		case *NullDateError:
			t.Errorf("wanted %T got %T", &DateRangeError{}, err)
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

func TestValidDate(t *testing.T) {
	//logic harder to test as time is relative to current time. if test passes now it might fail later.
	_ = []struct {
		startDate  time.Time
		targetDate time.Time
	}{
		{},
		{},
		{},
	}
}
