package time_extended

import (
	"fmt"
	"time"
)

type WeekdayParser struct {
	isInitialized bool
	daysOfWeek    map[string]time.Weekday
}

func (weekdayParser *WeekdayParser) init() {
	weekdayParser.daysOfWeek = map[string]time.Weekday{}
	for d := time.Sunday; d <= time.Saturday; d++ {
		name := d.String()
		weekdayParser.daysOfWeek[name] = d
		weekdayParser.daysOfWeek[name[:3]] = d
	}
}

func (weekdayParser *WeekdayParser) Parse(dayStr string) time.Weekday {
	if !weekdayParser.isInitialized {
		weekdayParser.init()
		weekdayParser.isInitialized = true
	}

	if day, ok := weekdayParser.daysOfWeek[dayStr]; ok {
		return day
	}

	errMsg := fmt.Sprintf("Invalid weekday %s", dayStr)
	panic(errMsg)
}
