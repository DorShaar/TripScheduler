package main

import (
	"fmt"
	logging "trip_scheduler/logger"
	"trip_scheduler/queue_adapter"
	"trip_scheduler/schedule"
)

func main() {
	logger := logging.Logger{}
	logger.Init()

	queueAdapter := queue_adapter.QueueAdapter{}
	queueAdapter.Init(logger)

	queueAdapter.Connect()
	defer queueAdapter.Disconnect()

	databasePath := "C:\\Users\\Public\\DorShaar\\GolangWork\\src\\trip_scheduler\\db\\files\\stav_third_year"

	scheduleBuilder := schedule.ScheduleBuilder{}
	scheduleBuilder.Init(logger)
	list := scheduleBuilder.BuildSchedulesFromFiles(databasePath)

	for e := list.Back(); e != nil; e = e.Prev() {
		schedule, ok := e.Value.(schedule.Schedule)
		if !ok {
			errMsg := fmt.Sprintf("Some element was of type %T, expected type Schedule\n",
				schedule)
			logger.LogError(errMsg)
			panic(errMsg)
		}

		sendSchedule(queueAdapter, "schedules", schedule)
	}


// // After getting shcedules back.
// 	schedulePrinter := schedule.SchedulePrinter{}
// 	schedulePrinter.Init(logger)
}

func sendSchedule(queueAdapter queue_adapter.QueueAdapter, queueName string, schedule schedule.Schedule) {
	schedule.

	queueAdapter.SendString(, queueName)
}
