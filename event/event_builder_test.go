package event

import (
	"testing"
	"time"
)

func Test_BuildEventFronYAML_TimeRangeWithOneValue(t *testing.T) {
	timeLayout := "22:22"
	expectedDuration, _ := time.ParseDuration("2h")
	expectedPrecautionDuration, _ := time.ParseDuration("15m")
	expectedStartingTime, _ := time.Parse(timeLayout, "07:00")
	expectedEndingTime, _ := time.Parse(timeLayout, "19:00")

	event, _ := BuildEventFronYAML("test_files/hyde_park.txt")
	if event.EventName != "Walking in Hyde Park" ||
		event.Location != "Hyde Park" ||
		event.EventTime.Duration != expectedDuration ||
		event.EventTime.PrecautionDuration != expectedPrecautionDuration ||
		event.TimeRangesList[0].StartingTime() == expectedStartingTime ||
		event.TimeRangesList[0].EndingTime() == expectedEndingTime {
		t.Errorf("Event time could not be parsed from YAML file well")
	}
}

func Test_BuildEventFronYAML_TimeRangeWithTwoValues(t *testing.T) {
	timeLayout := "15:04"

	expectedDay1 := "Sunday"
	expectedStartingTime1, _ := time.Parse(timeLayout, "07:00")
	expectedStartingTime1 = expectedStartingTime1.AddDate(0, 0, getDaysToAdd(expectedDay1))
	expectedEndingTime1, _ := time.Parse(timeLayout, "15:00")
	expectedEndingTime1 = expectedEndingTime1.AddDate(0, 0, getDaysToAdd(expectedDay1))

	expectedDay2 := "Saturday"
	expectedStartingTime2, _ := time.Parse(timeLayout, "08:00")
	expectedStartingTime2 = expectedStartingTime2.AddDate(0, 0, getDaysToAdd(expectedDay2))
	expectedEndingTime2, _ := time.Parse(timeLayout, "22:00")
	expectedEndingTime2 = expectedEndingTime2.AddDate(0, 0, getDaysToAdd(expectedDay2))

	event, _ := BuildEventFronYAML("test_files/camden_town.txt")
	if event.TimeRangesList[0].weekday != expectedDay1 ||
		event.TimeRangesList[0].StartingTime() != expectedStartingTime1 ||
		event.TimeRangesList[0].EndingTime() != expectedEndingTime1 ||
		event.TimeRangesList[1].weekday != expectedDay2 ||
		event.TimeRangesList[1].StartingTime() != expectedStartingTime2 ||
		event.TimeRangesList[1].EndingTime() != expectedEndingTime2 {
		t.Errorf("Event time could not be parsed from YAML file well")
	}
}

func Test_BuildEventFronYAML_TimeRangeWithStartingTimeOnly(t *testing.T) {
	timeLayout := "15:04"
	expectedStartingTime, _ := time.Parse(timeLayout, "19:30")

	event, _ := BuildEventFronYAML("test_files/lion_king.txt")
	if event.TimeRangesList[0].StartingTime() != expectedStartingTime {
		t.Errorf("Event time could not be parsed from YAML file well")
	}
}
