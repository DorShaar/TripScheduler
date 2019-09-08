package event

import (
	"fmt"
	"time"
)

type Event struct {
	EventName string
	Location  string
	EventTime EventTime
}

func (e Event) PrintEvent() string {
	eventInfo := fmt.Sprintf("'%s', located on %s, starts on %s %s for %.1f hours. Should Come %.0f minutes before.",
		e.EventName,
		e.Location,
		e.EventTime.ActualStartingTime.Weekday(),
		e.EventTime.ActualStartingTime.Format("2006-01-02 15:04"),
		e.EventTime.EventDurationInfo.Duration.Hours(),
		e.EventTime.EventDurationInfo.PrecautionDuration.Minutes())

	return eventInfo
}

type EventDurationInfo struct {
	Duration           time.Duration
	PrecautionDuration time.Duration
}

type EventTime struct {
	ActualStartingTime time.Time
	EventDurationInfo  EventDurationInfo
}

// The ActualStartingTime minus the PrecautionTime
func (et EventTime) StartingTime() time.Time {
	return et.ActualStartingTime.Add(-et.EventDurationInfo.PrecautionDuration)
}

// The ActualStartingTime plus the Duration of event
func (et EventTime) EndingTime() time.Time {
	return et.ActualStartingTime.Add(et.EventDurationInfo.Duration)
}

// The ActualStartingTime minus the PrecautionTime
func (currentEventTime EventTime) AreCoincide(eventTime EventTime) bool {
	if currentEventTime.StartingTime().After(eventTime.EndingTime()) {
		return false
	}

	if currentEventTime.EndingTime().Before(eventTime.StartingTime()) {
		return false
	}

	return true
}
