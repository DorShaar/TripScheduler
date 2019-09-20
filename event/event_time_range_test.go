package event

import (
	"testing"
	"time"
)

func Test_getAllPossibleStartingTimes(t *testing.T) {
	startingTime, _ := time.Parse(timeFormat, "Sep 13, 2019 04:00")
	endingTime, _ := time.Parse(timeFormat, "Sep 13, 2019 05:00")

	timeRange := CreateRange(startingTime, endingTime)

	possibleStartingTimeList := timeRange.getAllPossibleStartingTimes(15)
	if len(possibleStartingTimeList) != 4 {
		t.Errorf("getAllPossibleStartingTimes was not calculated well")
	}
}
