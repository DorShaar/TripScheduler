package event

import (
	"time"
)

type EventRange struct {
	EventName         string
	Location          string
	EventDurationInfo EventDurationInfo
	TimeRangeList     []TimeRange
}

func (eventRange *EventRange) AddTimeRange(newTimeRange TimeRange) {
	var shouldAddEvent bool = true
	for _, timeRange := range eventRange.TimeRangeList {
		if timeRange.AreCoincide(newTimeRange) {
			shouldAddEvent = false
			break
		}
	}

	if shouldAddEvent {
		eventRange.TimeRangeList = append(eventRange.TimeRangeList, newTimeRange)
	}
}

func (eventRange *EventRange) CreateEventsList() []Event {
	const intervalInMinutes = 15
	eventsList := make([]Event, 0)
	for _, timeRange := range eventRange.TimeRangeList {
		possibleStartingTimes := timeRange.getAllPossibleStartingTimes(intervalInMinutes)
		for _, startingTime := range possibleStartingTimes {
			newEvent := Event{
				EventName: eventRange.EventName,
				Location:  eventRange.Location,
				EventTime: EventTime{
					ActualStartingTime: startingTime,
					EventDurationInfo:  eventRange.EventDurationInfo}}

			eventsList = append(eventsList, newEvent)
		}
	}

	return eventsList
}

type TimeRange struct {
	startTime  time.Time
	endingTime time.Time
}

func (timeRange *TimeRange) CreateRange(start time.Time, end time.Time) {
	if start.After(end) {
		panic("Given starting time is after given ending time")
	}

	timeRange.startTime = start
	timeRange.endingTime = end
}

func (timeRange *TimeRange) CreateRangeByDuration(start time.Time, duration time.Duration) {
	if duration.Nanoseconds() < 0 || duration.Nanoseconds() == 0 {
		panic("Duration cannot be negative or zerp")
	}

	timeRange.startTime = start
	timeRange.endingTime = start.Add(duration)
}

func (timeRange TimeRange) StartingTime() time.Time {
	return timeRange.startTime
}

func (timeRange TimeRange) EndingTime() time.Time {
	return timeRange.endingTime
}

// The ActualStartingTime minus the PrecautionTime
func (currentTimeRange TimeRange) AreCoincide(timeRange TimeRange) bool {
	if currentTimeRange.StartingTime().After(timeRange.EndingTime()) {
		return false
	}

	if currentTimeRange.EndingTime().Before(timeRange.StartingTime()) {
		return false
	}

	return true
}

func (timeRange TimeRange) getAllPossibleStartingTimes(intervalInMinutes int) []time.Time {
	timeList := make([]time.Time, 0)
	var shouldStop bool = false
	for timeCounter := 0; !shouldStop; timeCounter += intervalInMinutes {
		currentTime := timeRange.startTime.Add(time.Duration(timeCounter) * time.Minute)
		shouldStop =
			currentTime.After(timeRange.endingTime) || currentTime == timeRange.endingTime
		if !shouldStop {
			timeList = append(timeList, currentTime)
		}
	}

	return timeList
}
