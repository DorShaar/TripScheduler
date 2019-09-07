package schedule

import (
	"trip_scheduler/event"
)

type Schedule struct {
	EventsList []event.Event
}

func (schedule *Schedule) AddEvent(newEvent event.Event) {
	var shouldAddEvent bool = true
	for _, event := range schedule.EventsList {
		if event.EventTime.AreCoincide(newEvent.EventTime) {
			shouldAddEvent = false
			break
		}
	}

	if shouldAddEvent {
		schedule.EventsList = append(schedule.EventsList, newEvent)
	}
}
