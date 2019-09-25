package schedule

import logging "trip_scheduler/logger"

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

	logger.Log("Printing schedule")
	for _, event := range schedule.eventsList {
		logger.Log(event.GetEventData())
	}
}
