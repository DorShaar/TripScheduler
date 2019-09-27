package schedule

import (
	"sort"
	"strconv"
	"trip_scheduler/event"
	logging "trip_scheduler/logger"
)

type Logger interface {
	Log(string)
	LogError(string)
}

type SchedulePrinter struct {
	logger Logger
}

func (schedulePrinter *SchedulePrinter) Init(logger logging.Logger) {
	schedulePrinter.logger = logger
}

func (schedulePrinter *SchedulePrinter) PrintSchedule(schedule Schedule) {
	logger := schedulePrinter.logger
	el := schedule.eventsList

	logger.Log("Printing schedule id " + strconv.Itoa(schedule.id))
	sort.Slice(el, func(i, j int) bool {
		return el[i].EventTime.ActualStartingTime().Before(
			el[j].EventTime.ActualStartingTime())
	})

	for _, event := range el {
		logger.Log(event.GetEventData())
	}
}

type eventsList []event.Event

// // Forward request for length
// func (el eventsList) Len() int {
// 	return len(el)
// }

// // Define compare
// func (el eventsList) Less(i, j int) bool {
// 	return el[i].EventTime.ActualStartingTime().Before(el[j].EventTime.ActualStartingTime())
// }

// // Define swap over an array
// func (el eventsList) Swap(i, j int) {
// 	el[i], el[j] = el[j], el[i]
// }
