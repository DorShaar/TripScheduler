package event

import (
	"time"
)

type TimeRange struct {
	weekday    string
	startTime  time.Time
	endingTime time.Time
}

func CreateRange(start time.Time, end time.Time) (timeRange TimeRange) {
	if start.After(end) {
		panic("Given starting time is after given ending time")
	}

	timeRange.startTime = start
	timeRange.endingTime = end
	return timeRange
}

func CreateRangeByDuration(start time.Time, duration time.Duration) (timeRange TimeRange) {
	if duration.Nanoseconds() <= 0 {
		panic("Duration cannot be negative or zero")
	}

	timeRange.startTime = start
	timeRange.endingTime = start.Add(duration)
	return timeRange
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

func (timeRange TimeRange) getAllPossibleStartingTimes(
	intervalInMinutes int,
	eventDuration time.Duration) []time.Time {
	var timeList []time.Time
	var shouldStop bool = false
	for timeCounter := 0; !shouldStop; timeCounter += intervalInMinutes {
		currentStartTime := timeRange.startTime.Add(time.Duration(timeCounter) * time.Minute)
		currentEndingTime := currentStartTime.Add(eventDuration)
		shouldStop =
			currentStartTime.After(timeRange.endingTime) ||
				currentStartTime == timeRange.endingTime ||
				currentEndingTime.After(timeRange.endingTime)

		if !shouldStop {
			timeList = append(timeList, currentStartTime)
		}
	}

	return timeList
}
