package schedule

import (
	"testing"
	"time"
	"trip_scheduler/event"
)

const timeFormat = "Jan _2, 2006 15:04"

func Test_TryAddEvent_EmptySchedule_Success(t *testing.T) {
	actualStartingTime, _ := time.Parse(timeFormat, "Sep 13, 2019 04:00")
	duration, _ := time.ParseDuration("90m")
	precautionDuration, _ := time.ParseDuration("25m")

	event := CreateEvent(actualStartingTime, duration, precautionDuration)

	schedule := Schedule{}
	schedule.TryAddEvent(event)
	if len(schedule.eventsList) != 1 {
		t.Error("Event was not added to the schedule")
	}
}

func Test_TryAddEvent_Success(t *testing.T) {
	// Setting first event.
	actualStartingTime1, _ := time.Parse(timeFormat, "Sep 13, 2019 04:00")
	duration1, _ := time.ParseDuration("90m")
	precautionDuration1, _ := time.ParseDuration("25m")

	event1 := CreateEvent(actualStartingTime1, duration1, precautionDuration1)

	schedule := Schedule{}
	schedule.TryAddEvent(event1)

	// Setting second event.
	actualStartingTime2, _ := time.Parse(timeFormat, "Sep 13, 2019 08:00")
	duration2, _ := time.ParseDuration("90m")
	precautionDuration2, _ := time.ParseDuration("25m")

	event2 := CreateEvent(actualStartingTime2, duration2, precautionDuration2)
	schedule.TryAddEvent(event2)

	if len(schedule.eventsList) != 2 {
		t.Error("Event was not added to the schedule")
	}
}

func Test_TryAddEvent_Fail(t *testing.T) {
	// Setting first event.
	actualStartingTime1, _ := time.Parse(timeFormat, "Sep 13, 2019 04:00")
	duration1, _ := time.ParseDuration("90m")
	precautionDuration1, _ := time.ParseDuration("25m")

	event1 := CreateEvent(actualStartingTime1, duration1, precautionDuration1)

	schedule := Schedule{}
	schedule.TryAddEvent(event1)

	// Setting second event.
	actualStartingTime2, _ := time.Parse(timeFormat, "Sep 13, 2019 05:00")
	duration2, _ := time.ParseDuration("90m")
	precautionDuration2, _ := time.ParseDuration("25m")

	event2 := CreateEvent(actualStartingTime2, duration2, precautionDuration2)
	schedule.TryAddEvent(event2)

	if len(schedule.eventsList) != 1 {
		t.Error("Schedule should have only one event")
	}
}

func CreateEvent(
	actualStartingTime time.Time,
	duration time.Duration,
	precautionDuration time.Duration) event.Event {

	timeRangesList := make([]event.TimeRange, 0)
	return event.BuildEventByParameters(
		"Shopping !",
		"Oxford Street",
		timeRangesList,
		duration,
		precautionDuration,
		actualStartingTime)
}
