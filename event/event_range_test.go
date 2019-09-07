package event

import (
	"testing"
	"time"
)

// const timeFormat = "Jan _2, 2006 15:04"

func Test_getAllPossibleStartingTimes(t *testing.T) {
	startingTime, _ := time.Parse(timeFormat, "Sep 13, 2019 04:00")
	endingTime, _ := time.Parse(timeFormat, "Sep 13, 2019 05:00")

	timeRange := TimeRange{}
	timeRange.CreateRange(startingTime, endingTime)

	possibleStartingTimeList := timeRange.getAllPossibleStartingTimes(15)
	if len(possibleStartingTimeList) != 4 {
		t.Errorf("getAllPossibleStartingTimes was not calculated well")
	}
}
