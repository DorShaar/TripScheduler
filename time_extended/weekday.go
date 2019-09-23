package time_extended

import (
	"fmt"
	"time"
)

var daysOfWeek = map[string]time.Weekday{}

func init() {
	for d := time.Sunday; d <= time.Saturday; d++ {
		name := d.String()
		daysOfWeek[name] = d
		daysOfWeek[name[:3]] = d
	}
}

func ParseWeekday(dayStr string) (time.Weekday, error) {
	if day, ok := daysOfWeek[dayStr]; ok {
		return day, nil
	}

	errMsg := fmt.Sprintf("Invalid weekday %s", dayStr)
	panic(errMsg)
}
