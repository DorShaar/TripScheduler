package schedule

import (
	"sort"
	"strconv"
)

type Logger interface {
	Log(string)
	LogError(string)
}

type SchedulePrinter struct {
	logger Logger
}

func (schedulePrinter *SchedulePrinter) Init(logger Logger) {
	schedulePrinter.logger = logger
}

func (schedulePrinter *SchedulePrinter) PrintSchedule(schedule Schedule) {
	logger := schedulePrinter.logger
	el := schedule.eventsList

	logger.Log("Printing schedule id " + strconv.Itoa(schedule.Id()))
	sort.Slice(el, func(i, j int) bool {
		return el[i].EventTime.ActualStartingTime().Before(
			el[j].EventTime.ActualStartingTime())
	})

	for _, event := range el {
		logger.Log(event.GetEventData())
	}
}
