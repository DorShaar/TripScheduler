package main

import (
	"trip_scheduler/event"
)

func main() {
	e := event.EventRange{}
	e.Test()
	// duration, _ := time.ParseDuration("90m")
	// precautionTime, _ := time.ParseDuration("25m")

	// e := event.Event{
	// 	EventName: "Shopping !",
	// 	Location:  "Oxford Street",
	// 	EventTime: event.EventTime{
	// 		ActualStartingTime: time.Now(),
	// 		Duration:           duration,
	// 		PrecautionDuration: precautionTime}}

	// d := schedule.Schedule{}
	// fmt.Println(d.EventLists == nil)
	// fmt.Printf("%d\n", len(d.EventLists))

	// d.AddEvent(e)
	// fmt.Printf("%d", len(d.EventLists))
}
