package schedule

import "trip_scheduler/event"

func BuildSchedulesFromFiles(filesDirectory string) {
	
	schedulesList := make([]Schedule, 0)
	for _, eventRange := range eventRangeList {
		... = eventRange.CreateEventsList()
	}
}
