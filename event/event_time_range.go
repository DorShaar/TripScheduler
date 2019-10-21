package event

import (
	"time"
	"trip_scheduler/time_extended"
)

var epochTimeYear = 1970

type TimeRange struct {
	weekday    string
	startTime  time.Time
	endingTime time.Time
}

func CreateRange(start time.Time, end time.Time) (timeRange TimeRange) {
	if start.After(end) {
		panic("Given starting time is after given ending time")
	}

	if start.Year() == 0 {
		start = start.AddDate(epochTimeYear, 0, 0)
	}

	if end.Year() == 0 {
		end = end.AddDate(epochTimeYear, 0, 0)
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
	if start.Year() == 0 {
		start = start.AddDate(epochTimeYear, 0, 0)
	}

	timeRange.endingTime = start.Add(duration)
	return timeRange
}

func (timeRange TimeRange) StartingTime() time.Time {
	return timeRange.startTime.AddDate(0, 0, getDaysToAdd(timeRange.weekday))
}

func (timeRange TimeRange) EndingTime() time.Time {
	return timeRange.endingTime.AddDate(0, 0, getDaysToAdd(timeRange.weekday))
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
		currentStartTime := timeRange.StartingTime().
			Add(time.Duration(timeCounter) * time.Minute)
		currentEndingTime := currentStartTime.Add(eventDuration)

		shouldStop =
			currentStartTime.After(timeRange.EndingTime()) ||
				currentStartTime == timeRange.EndingTime() ||
				currentEndingTime.After(timeRange.EndingTime())

		if !shouldStop {
			timeList = append(timeList, currentStartTime)
		}
	}

	return timeList
}

func getDaysToAdd(weekday string) int {
	var daysToAdd int
	if weekday != "" {
		weekdayParser := time_extended.WeekdayParser{}
		daysToAdd = (int)(weekdayParser.Parse(weekday))
	}

	return daysToAdd
}
