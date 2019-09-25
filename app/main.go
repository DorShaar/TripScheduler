package main

import (
	logging "trip_scheduler/logger"
	"trip_scheduler/schedule"
)

func main() {
	databasePath := "C:\\Users\\Public\\DorShaar\\GolangWork\\src\\trip_scheduler\\db\\files\\stav_third_year"
	logger := logging.Logger{}
	logger.Init()

	scheduleBuilder := schedule.ScheduleBuilder{}
	scheduleBuilder.Init(logger)
	scheduleBuilder.BuildSchedulesFromFiles(databasePath)
}
