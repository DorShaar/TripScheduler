package schedule

import (
	"trip_scheduler/event"
)

type Schedule struct {
	id         int
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

func (schedule *Schedule) Id() int {
	return schedule.id
}

func (schedule *Schedule) GetEventsListCopy() []event.Event {
	eventsCopy := make([]event.Event, len(schedule.eventsList))
	copy(eventsCopy, schedule.eventsList)
	return eventsCopy
}
