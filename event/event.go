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
		e.EventTime.Duration.Hours(),
		e.EventTime.PrecautionTime.Minutes())

	return eventInfo
}

type EventTime struct {
	ActualStartingTime time.Time
	Duration           time.Duration
	PrecautionTime     time.Duration
}

// The ActualStartingTime minus the PrecautionTime
func (et EventTime) StartingTime() time.Time {
	return et.ActualStartingTime.Add(-et.PrecautionTime)
}

// The ActualStartingTime plus the Duration of event
func (et EventTime) EndingTime() time.Time {
	return et.ActualStartingTime.Add(et.Duration)
}

// The ActualStartingTime minus the PrecautionTime
func (currentEventTime EventTime) AreCoincide(eventTime EventTime) bool {
	if eventTime.StartingTime().After(currentEventTime.EndingTime()) {
		return false
	}

	if eventTime.EndingTime().Before(currentEventTime.StartingTime()) {
		return false
	}

	return true
}
