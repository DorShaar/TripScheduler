package event

import (
	"fmt"
	"time"
)

const timeFormat = "Jan _2, 2006 15:04"

type EventRange struct {
	EventName      string
	Location       string
	TimeRangeLists []TimeRange
}

func (eventRange EventRange) Test() {
	var t TimeRange

	startingTime, _ := time.Parse(timeFormat, "Sep 13, 2019 05:00")
	endingTime, _ := time.Parse(timeFormat, "Sep 13, 2019 08:00")
	t.CreateRange(startingTime, endingTime)

	// t.CreateRange(endingTime, startingTime)
	// fmt.Println("good")
}

func (eventRange *EventRange) AddTimeRange(newTimeRange TimeRange) {
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
