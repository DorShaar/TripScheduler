package schedule

import "trip_scheduler/event"

type ScheduleMaker struct {
	SchedulesList []Schedule
}

func (ScheduleMaker *ScheduleMaker) ComputeSchedules(eventRangeList []event.EventRange) {
	for _, eventRange := range eventRangeList {
		for _, timeRange := range eventRange.TimeRangeList {
			timeRange.
		}
	}
}
