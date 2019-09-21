package schedule

import (
	"trip_scheduler/event"
)

type Schedule struct {
	eventsList []event.Event
}

func (schedule *Schedule) TryAddEvent(newEvent event.Event) bool {
	var shouldAddEvent bool = true
	for _, event := range schedule.eventsList {
		if event.EventTime.AreCoincide(newEvent.EventTime) {
			shouldAddEvent = false
			break
		}
	}

	if shouldAddEvent {
		schedule.eventsList = append(schedule.eventsList, newEvent)
	}

	return shouldAddEvent
}
