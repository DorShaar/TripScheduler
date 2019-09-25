package event

import (
	"fmt"
	"time"
)

type Event struct {
	EventName      string
	Location       string
	TimeRangesList []TimeRange
	EventTime      EventTime
}

func (e Event) GetEventData() string {
	return fmt.Sprintf("'%s' at %s",
		e.EventName,
		e.EventTime.actualStartingTime.Format("2006-01-02 15:04"))
}

func (event *Event) RegisterTimeRange(newTimeRange TimeRange) {
	var shouldAddEvent bool = true
	for _, timeRange := range event.TimeRangesList {
		if timeRange.AreCoincide(newTimeRange) {
			shouldAddEvent = false
			break
		}
	}

	if shouldAddEvent {
		event.TimeRangesList = append(event.TimeRangesList, newTimeRange)
	}
}

// Get all possible events with updated ActualStartingTime.
func (event Event) CreateEventsList() []Event {
	const intervalInMinutes = 15
	eventsList := make([]Event, 0)
	for _, timeRange := range event.TimeRangesList {
		possibleStartingTimes := timeRange.getAllPossibleStartingTimes(
			intervalInMinutes, event.EventTime.Duration)
		for _, startingTime := range possibleStartingTimes {
			event.EventTime.actualStartingTime = startingTime
			eventsList = append(eventsList, event)
		}
	}

	return eventsList
}

type EventTime struct {
	Duration           time.Duration
	PrecautionDuration time.Duration
	actualStartingTime time.Time
}

// The ActualStartingTime minus the PrecautionTime
func (et EventTime) StartingTime() time.Time {
	return et.actualStartingTime.Add(-et.PrecautionDuration)
}

// The ActualStartingTime plus the Duration of event
func (et EventTime) EndingTime() time.Time {
	return et.actualStartingTime.Add(et.Duration)
}

func (et EventTime) ActualStartingTime() time.Time {
	return et.actualStartingTime
}

func (currentEventTime EventTime) AreCoincide(eventTime EventTime) bool {
	if currentEventTime.StartingTime().After(eventTime.EndingTime()) {
		return false
	}

	if currentEventTime.EndingTime().Before(eventTime.StartingTime()) {
		return false
	}

	return true
}
