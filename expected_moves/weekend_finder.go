package expected_moves

import (
	"time"
)

type CalendarInterface interface {
	Today() time.Time
}

type WeekendFinder struct {
	Calendar CalendarInterface
}

type Calendar struct {}

func (c Calendar) Today() time.Time {
	return time.Now()
}

func (wf WeekendFinder) GetNextFriday() time.Time {
	weekday := wf.Calendar.Today().Weekday()

	return wf.Calendar.Today().Add(time.Duration(getDaysToAdd(weekday)) * 24 * time.Hour)
}

func getDaysToAdd(weekday time.Weekday) int {
	switch weekday {
	case time.Thursday:
		return 1
	case time.Wednesday:
		return 2
	case time.Tuesday:
		return 3
	case time.Monday:
		return 4
	case time.Sunday:
		return 5
	case time.Saturday:
		return 6
	case time.Friday:
		return 7
	}

	return 0
}
