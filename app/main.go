package main

import (
	"time"
	"trip_scheduler/event"
)

func main() {
	duration, _ := time.ParseDuration("90m")
	precautionTime, _ := time.ParseDuration("25m")

	e := event.Event{
		EventName:      "Shopping !",
		StartingTime:   time.Now(),
		Duration:       duration,
		PrecautionTime: precautionTime,
		Location:       "Oxford Street"}

	e.PrintEvent()
}
