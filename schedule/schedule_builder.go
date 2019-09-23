package schedule

import (
	"container/list"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"trip_scheduler/event"
)

func BuildSchedulesFromFiles(filesDirectory string) list.List {
	shouldProcessQueue := list.New()
	doneQueue := list.New()

	emptySchedule := Schedule{}
	shouldProcessQueue.PushBack(emptySchedule)

	allRegisteredEvents := getAllEventsList(filesDirectory)
	for _, registeredEvent := range allRegisteredEvents {
		for shouldProcessQueue.Len() > 0 {
			element := shouldProcessQueue.Front()
			shouldProcessQueue.Remove(element)

			originalSchedule, ok := element.Value.(Schedule)
			if !ok {
				errMsg := fmt.Sprintf("Some element was of type %T, expected type Schedule\n",
					originalSchedule)
				panic(errMsg)
			}

			for _, possibleEvent := range registeredEvent.CreateEventsList() {
				newSchedule := originalSchedule
				if newSchedule.TryAddEvent(possibleEvent) {
					doneQueue.PushBack(newSchedule)
				}
			}
		}

		tempQueue := shouldProcessQueue
		shouldProcessQueue = doneQueue
		doneQueue = tempQueue
	}

	doneQueue = shouldProcessQueue
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
