package schedule

import "trip_scheduler/event"

type ScheduleMaker struct {
	SchedulesList []Schedule
}

func (scheduleMaker *ScheduleMaker) ComputeSchedules(eventRangeList []event.EventRange) {
	events := make([]event.Event, 50, 50)
	for _, eventRange := range eventRangeList {
		for _, timeRange := range eventRange.TimeRangeList {
			newEvent := event.Event {
				EventName:	eventRange.EventName,
				Location:  	eventRange.Location,
				EventTime:	EventTime
			}
			
			append(events, newEvent)
		}
	}
}
