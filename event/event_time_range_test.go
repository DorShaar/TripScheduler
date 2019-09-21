package event

import (
	"testing"
	"time"
)

func Test_getAllPossibleStartingTimes_Expected3StartingTimes(t *testing.T) {
	startingTime, _ := time.Parse(timeFormat, "Sep 13, 2019 04:00")
	endingTime, _ := time.Parse(timeFormat, "Sep 13, 2019 05:00")
	intervalInMinuts := 15
	eventDuration, _ := time.ParseDuration("20m")
	timeRange := CreateRange(startingTime, endingTime)

	possibleStartingTimeList := timeRange.getAllPossibleStartingTimes(
		intervalInMinuts, eventDuration)
	if len(possibleStartingTimeList) != 3 {
		t.Errorf("getAllPossibleStartingTimes was not calculated well")
	}
}
