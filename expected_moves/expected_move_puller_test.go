package expected_moves

import (
	"testing"
	"time"
)

func PullExpectedMoveForUpcomingWeek (t *testing.InternalTest) {

}

func TestGetDateForNextFriday (t *testing.T) {
	today, _ := time.ParseInLocation("2006-01-02", "2020-08-22", time.Local)
	expected_friday, _ := time.ParseInLocation("2006-01-02", "2020-08-28", time.Local)
	finder := &WeekendFinder{FakeCalendar {
		T: today,
	}}

	friday := finder.GetNextFriday()

	if !friday.Equal(expected_friday) {
		t.Errorf("Friday not found: expected 2020-08-28, found %s", friday)
	}
}

type FakeCalendar struct {
	T time.Time
}

func (clock FakeCalendar) Today() time.Time {
	return clock.T
}
