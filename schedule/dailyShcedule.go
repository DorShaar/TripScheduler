package schedule

import (
	"trip_scheduler/event"
)

type Schedule struct {
	EventLists []event.Event
}

func (schedule *Schedule) AddEvent(newEvent event.Event) {
	var shouldAddEvent bool = true
	for _, event := range schedule.EventLists {
		if event.EventTime.AreCoincide(newEvent.EventTime) {
			shouldAddEvent = false
			break
		}
	}

	if shouldAddEvent {
		schedule.EventLists = append(schedule.EventLists, newEvent)
	}
}
