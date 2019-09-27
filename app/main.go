package main

import (
	"fmt"
	logging "trip_scheduler/logger"
	"trip_scheduler/schedule"
)

func main() {
	databasePath := "C:\\Users\\Public\\DorShaar\\GolangWork\\src\\trip_scheduler\\db\\files\\stav_third_year"
	logger := logging.Logger{}
	logger.Init()

	scheduleBuilder := schedule.ScheduleBuilder{}
	scheduleBuilder.Init(logger)
	list := scheduleBuilder.BuildSchedulesFromFiles(databasePath)

	schedulePrinter := schedule.SchedulePrinter{}
	schedulePrinter.Init(logger)

	for e := list.Back(); e != nil; e = e.Prev() {
		schedule, ok := e.Value.(schedule.Schedule)
		if !ok {
			errMsg := fmt.Sprintf("Some element was of type %T, expected type Schedule\n",
				schedule)
			logger.LogError(errMsg)
			panic(errMsg)
		}

		schedulePrinter.PrintSchedule(schedule)
		logger.Log("")
	}
}
