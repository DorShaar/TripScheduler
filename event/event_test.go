package event

import (
	"testing"
	"time"
)

const timeFormat = "Jan _2, 2006 15:04"

func Test_StartingTime(t *testing.T) {
	actualStartingTime, _ := time.Parse(timeFormat, "Sep 13, 2019 04:00")
	precautionDuration, _ := time.ParseDuration("45m")

	eventTime := EventTime{
		actualStartingTime: actualStartingTime,
		PrecautionDuration: precautionDuration}

	time := eventTime.StartingTime()
	if time.Day() != 13 || time.Hour() != 3 || time.Minute() != 15 {
		t.Errorf("StartingTime was not calculated well")
	}
}

func Test_EndingTime(t *testing.T) {
	actualStartingTime, _ := time.Parse(timeFormat, "Sep 13, 2019 16:00")
	duration, _ := time.ParseDuration("3h15m")

	eventTime := EventTime{
		actualStartingTime: actualStartingTime,
		Duration:           duration}

	time := eventTime.EndingTime()
	if time.Day() != 13 || time.Hour() != 19 || time.Minute() != 15 {
		t.Errorf("StartingTime was not calculated well")
	}
}

func Test_AreCoincide_False(t *testing.T) {
	actualStartingTime1, _ := time.Parse(timeFormat, "Sep 13, 2019 04:00")
	duration1, _ := time.ParseDuration("3h")
	precautionDuration1, _ := time.ParseDuration("45m")

	eventTime1 := EventTime{
		actualStartingTime: actualStartingTime1,
		Duration:           duration1,
		PrecautionDuration: precautionDuration1}

	actualStartingTime2, _ := time.Parse(timeFormat, "Sep 13, 2019 08:00")
	duration2, _ := time.ParseDuration("1h")
	precautionDuration2, _ := time.ParseDuration("10m")

	eventTime2 := EventTime{
		actualStartingTime: actualStartingTime2,
		Duration:           duration2,
		PrecautionDuration: precautionDuration2}

	if eventTime1.AreCoincide(eventTime2) {
		t.Errorf("Events are coinciding although they should not.")
	}
}

func Test_AreCoincide_True(t *testing.T) {
	actualStartingTime1, _ := time.Parse(timeFormat, "Sep 13, 2019 at 04:00pm")
	duration1, _ := time.ParseDuration("3h")
	precautionDuration1, _ := time.ParseDuration("45m")

	eventTime1 := EventTime{
		actualStartingTime: actualStartingTime1,
		Duration:           duration1,
		PrecautionDuration: precautionDuration1}

	actualStartingTime2, _ := time.Parse(timeFormat, "Sep 13, 2019 at 08:00pm")
	duration2, _ := time.ParseDuration("1h")
	precautionDuration2, _ := time.ParseDuration("2h")

	eventTime2 := EventTime{
		actualStartingTime: actualStartingTime2,
		Duration:           duration2,
		PrecautionDuration: precautionDuration2}

	if !eventTime1.AreCoincide(eventTime2) {
		t.Error("Events are not coinciding although they should coincide")
	}
}

func Test_CreateEventsList_ListLengthIs1(t *testing.T) {
	timeLayout := "15:04"
	startingTime, _ := time.Parse(timeLayout, "14:15")
	endingTime, _ := time.Parse(timeLayout, "15:45")
	timeRangeList := make([]TimeRange, 0)
	timeRangeList = append(timeRangeList, TimeRange{
		startTime:  startingTime,
		endingTime: endingTime})
	duration, _ := time.ParseDuration("1.5h")
	precautionDuration, _ := time.ParseDuration("0m")
	dummyActualStartingTime, _ := time.Parse(timeLayout, "00:00")
	event := BuildEventByParameters(
		"DummyEvent",
		"DummyLocation",
		timeRangeList,
		duration,
		precautionDuration,
		dummyActualStartingTime)

	if len(event.CreateEventsList()) != 1 {
		t.Error("CreateEventsList was not performed well")
	}
}
