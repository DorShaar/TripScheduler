package main

import (
	"fmt"
	"time"
)

type Event struct {
	EventName      string
	StartingTime   time.Time
	Duration       time.Duration
	PrecautionTime time.Duration
	Location       string
}

func (e Event) PrintEvent() {
	fmt.Printf("'%s', located on %s, starts on %s %s for %.1f hours. Should Come %.0f minutes before.",
		e.EventName,
		e.Location,
		e.StartingTime.Weekday(),
		e.StartingTime.Format("2006-01-02 15:04"),
		e.Duration.Hours(),
		e.PrecautionTime.Minutes())
}

func main() {
	duration, _ := time.ParseDuration("90m")
	precautionTime, _ := time.ParseDuration("25m")

	e := Event{
		EventName:      "Shopping !",
		StartingTime:   time.Now(),
		Duration:       duration,
		PrecautionTime: precautionTime,
		Location:       "Oxford Street"}
	e.PrintEvent()
}
