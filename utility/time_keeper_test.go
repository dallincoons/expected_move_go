package utility

import (
	"expected_move/cmd"
	"testing"
	"time"
)

func TestTimeIsBeforeToday (t *testing.T) {
	tk := cmd.NewTimeKeeper()

	yesterday := time.Now().AddDate(0, 0, -1)

	if tk.IsBeforeToday(yesterday) != true {
		t.Error("Yesterday is not considered before today's date when subtracting day")
	}

	today := time.Now()

	if tk.IsBeforeToday(today) != false {
		t.Error("Yesterday is not considered before today's date when using today's date")
	}
}

func TestGetAllDatesBetweenThenAndNow (t *testing.T) {
	now := time.Date(2020, time.August, 3, 0, 0,0, 0, time.Local)
	tk := cmd.NewTimeKeeperWithNow(now)

	from := now.AddDate(0, 0, -14)
	daysSince := tk.GetWeekdaysSince(from)

	dayMap := make(map[string]bool)
	for _, day := range daysSince {
		dayMap[day.Format("2006-01-02")]	= true
	}

	if len(dayMap) != 11 {
		t.Error("Expected 11 days, got ", len(dayMap))
	}

	if found := dayMap["2020-08-02"]; found {
		t.Error("Found weekend day, 2020-08-02")
	}

	if found := dayMap["2020-08-01"]; found {
		t.Error("Found weekend day, 2020-08-01")
	}

	if found := dayMap["2020-07-31"]; !found {
		t.Error("Expected weekday, 2020-07-31")
	}

	if found := dayMap["2020-07-27"]; !found {
		t.Error("Expected weekday, 2020-07-27")
	}

	if found := dayMap["2020-07-20"]; !found {
		t.Error("Expected weekday, 2020-07-22")
	}

	if found := dayMap["2020-07-19"]; found {
		t.Error("Found weekend day, 2020-07-19")
	}

	if found := dayMap["2020-07-17"]; !found {
		t.Error("Expected weekday, 2020-07-17")
	}
}
