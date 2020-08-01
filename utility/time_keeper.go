package utility

import "time"

type timeKeeper struct {
	Now time.Time
}

func NewTimeKeeper () *timeKeeper {
	return &timeKeeper{
		Now: time.Now(),
	}
}

func NewTimeKeeperWithNow(now time.Time) *timeKeeper {
	return &timeKeeper{
		Now: now,
	}
}

func (tk *timeKeeper) IsBeforeToday(from time.Time) bool {
	return time.Now().Sub(from).Hours() >= 24
}

func (tk *timeKeeper) GetWeekdaysSince(from time.Time) []time.Time {
	current := from
	var daysSince []time.Time
	for current.Before(tk.Now) || current.Equal(tk.Now) {
		if isTradingDay(current) {
			daysSince = append(daysSince, current)
		}
		current = current.AddDate(0, 0, 1)
	}

	return daysSince
}

func isTradingDay(day time.Time) bool {
	return day.Weekday() != time.Saturday && day.Weekday() != time.Sunday
}
