package main

import (
	"fmt"
	"log"
	logging "trip_scheduler/logger"
	"trip_scheduler/queue_adapter"
	"trip_scheduler/schedule"
	"trip_scheduler/schedule_dto"

	"github.com/gogo/protobuf/proto"
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

	// TODO: After getting shcedules back.
	// 	schedulePrinter := schedule.SchedulePrinter{}
	// 	schedulePrinter.Init(logger)
}

func sendSchedule(queueAdapter queue_adapter.QueueAdapter, queueName string, schedule schedule.Schedule) {
	scheduleDTO := schedule_dto.CreateDTOSchedule(schedule)

	serializedSchedule, err := proto.Marshal(scheduleDTO)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}

	deserializedschedule := &schedule_dto.Schedule{}
	err = proto.Unmarshal(serializedSchedule, deserializedschedule)
	if err != nil {
		panic(err)
	}

	fmt.Println(deserializedschedule)

	// TODO queueAdapter.SendBytes(serializedSchedule, queueName)
}
