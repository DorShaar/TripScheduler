package schedule

import (
	"container/list"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"trip_scheduler/event"
	logging "trip_scheduler/logger"
)

type ScheduleBuilder struct {
	logger        LoggerInterface
	itemIdCreator ItemIdCreator
}

func (scheduleBuilder *ScheduleBuilder) Init(logger logging.Logger) {
	scheduleBuilder.logger = logger
}

func (scheduleBuilder *ScheduleBuilder) BuildSchedulesFromFiles(filesDirectory string) list.List {
	logger := scheduleBuilder.logger

	logger.Log("Building schedules from files in " + filesDirectory)

	shouldProcessQueue := list.New()
	doneQueue := list.New()

	emptySchedule := Schedule{}
	shouldProcessQueue.PushBack(emptySchedule)
	logger.Log("Empty schedule pushed back to queue")

	allRegisteredEvents := getAllEventsList(filesDirectory)
	for _, registeredEvent := range allRegisteredEvents {
		logger.Log("Queue size: " + strconv.Itoa(shouldProcessQueue.Len()))
		for shouldProcessQueue.Len() > 0 {
			element := shouldProcessQueue.Front()
			shouldProcessQueue.Remove(element)

			originalSchedule, ok := element.Value.(Schedule)
			if !ok {
				errMsg := fmt.Sprintf("Some element was of type %T, expected type Schedule\n",
					originalSchedule)
				logger.LogError(errMsg)
				panic(errMsg)
			}

			for _, possibleEvent := range registeredEvent.CreateEventsList() {
				newSchedule := scheduleBuilder.copySchedule(originalSchedule)
				logger.Log("Trying to add to schedule " + strconv.Itoa(originalSchedule.id) +
					" event " + possibleEvent.GetEventData())
				if newSchedule.TryAddEvent(possibleEvent) {
					doneQueue.PushBack(newSchedule)
					logger.Log("Event added to schedule id " + strconv.Itoa(newSchedule.id))
				} else {
					logger.Log("Event was not added")
				}
			}
		}

		tempQueue := shouldProcessQueue
		shouldProcessQueue = doneQueue
		doneQueue = tempQueue
	}

	doneQueue = shouldProcessQueue
	logger.Log("Done building " + strconv.Itoa(doneQueue.Len()) + " schedules")
	return *doneQueue
}

func getAllEventsList(filesDirectory string) []event.Event {
	allEvents := make([]event.Event, 0)
	filePaths := getFilesFromDirectory(filesDirectory)
	for _, filePath := range filePaths {
		event, err := event.BuildEventFronYAML(filePath)
		if err != nil {
			fmt.Printf("Error parsing %s to event", filePath)
			continue
		}

		allEvents = append(allEvents, event)
	}

	return allEvents
}

func getFilesFromDirectory(filesDirectory string) (filePaths []string) {
	files, err := ioutil.ReadDir(filesDirectory)
	if err != nil {
		fmt.Println(err)
		return filePaths
	}

	for _, file := range files {
		filePaths = append(filePaths, filepath.Join(filesDirectory, file.Name()))
	}

	return filePaths
}

func (scheduleBuilder *ScheduleBuilder) copySchedule(originalSchedule Schedule) (copiedSchedule Schedule) {
	copiedSchedule = originalSchedule
	copiedSchedule.id = scheduleBuilder.itemIdCreator.NextId()
	return copiedSchedule
}
