package schedule

import "trip_scheduler/event"

type ScheduleMaker struct {
	SchedulesList []Schedule
}

func (scheduleMaker *ScheduleMaker) ComputeSchedules(eventRangeList []event.EventRange) {
	events := make([]event.Event, 50, 50)
	for _, eventRange := range eventRangeList {
		... = eventRange.CreateEventsList()
	}
}
