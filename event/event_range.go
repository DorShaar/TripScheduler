package event

import (
	"fmt"
	"time"
)

const timeFormat = "Jan _2, 2006 15:04"

type EventRange struct {
	EventName     string
	Location      string
	TimeRangeList []TimeRange
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

// func (eventRange *EventRange) CreateEventsList() []Event {
// 	eventsList := make([]Event, 50, 50)
// 	for _, timeRange := range eventRange.TimeRangeList {
// 		newEvent := Event{
// 			EventName: eventRange.EventName,
// 			Location:  eventRange.Location,
// 			EventTime: eventRange.TimeRangeList,
// 		}

// 		append(eventsList, newEvent)
// 	}
// }

type TimeRange struct {
	startTime  time.Time
	endingTime time.Time
}

func (timeRange *TimeRange) CreateRange(start time.Time, end time.Time) {
	if start.After(end) {
		panic("Given starting time is after given ending time")
	}

	timeRange.startTime = start
	fmt.Println(start)
	timeRange.endingTime = end
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

func (timeRange TimeRange) getAllPossibleStartingTimes(intervalInMinuts int) []time.Time {
	timeList := make([]time.Time, 20, 20)
	var shouldStop bool = false
	for timeCounter := 0; !shouldStop; timeCounter += intervalInMinuts {
		currentTime := timeRange.startTime.Add(time.Duration(timeCounter))
		shouldStop = currentTime.Before(timeRange.endingTime)
		timeList = append(timeList, currentTime)
	}

	return timeList
}
