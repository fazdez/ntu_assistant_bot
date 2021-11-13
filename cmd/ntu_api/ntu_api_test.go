package ntuapi

import "testing"

func TestGetTeachingWeek(t *testing.T) {
	ntuapi := New(Config{})

	tests := []struct {
		week, teachingWeek int
	}{
		{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6}, {7, 7},
		{9, 8}, {10, 9}, {11, 10}, {12, 11}, {13, 12}, {14, 13},
		{8, 0}, {0, -1}, {15, -1}, {-100, -1}, {16, -1},
	}

	for _, test := range tests {
		got, _ := ntuapi.getTeachingWeek(test.week)
		if got != test.teachingWeek {
			t.Errorf("wanted teaching week %d got %d <input %d>", test.teachingWeek, got, test.week)
		}
	}
}
