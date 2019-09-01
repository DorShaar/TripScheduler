package main

import (  
    "fmt"
    "time"
)

type Event struct {  
    EventName		string
    StartingTime    time.Time
    Duration		time.Duration
    PrecautionTime 	time.Duration
}

func (e Event) PrintEvent() {  
    fmt.Printf("%s starts on %s %s for %.1f hours. Should Come %.0f minutes before.", 
    	e.EventName,
    	e.StartingTime.Weekday(), 
    	e.StartingTime.Format("2006-01-02 15:04"),
    	e.Duration.Hours(),
    	e.PrecautionTime.Minutes())
}

func main() {
	duration, _ := time.ParseDuration("90m")
	precautionTime, _ := time.ParseDuration("25m")

	e:= Event {
			EventName: 		"London!!",
		 	StartingTime:	 time.Now(),
		 	Duration:		 duration,
		 	PrecautionTime:	 precautionTime}
	e.PrintEvent()
}