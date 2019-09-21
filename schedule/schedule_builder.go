package schedule

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"trip_scheduler/event"
)

func BuildSchedulesFromFiles(filesDirectory string) (schedules []Schedule) {
	schedules = make([]Schedule, 1)
	schedules[0] = Schedule{}
	// NOT FINISHED
	allRegisteredEvents := getAllEventsList(filesDirectory)
	for _, registeredEvent := range allRegisteredEvents {
		for _, possibleEvent := range registeredEvent.CreateEventsList() {
			// for _, schedule := range schedules {
			if !schedules[0].TryAddEvent(possibleEvent) {
				fmt.Printf("Event %s could not be added", possibleEvent.EventName)
			} else {
				fmt.Printf("schedule contains %d events\n", len(schedules[0].eventsList))
			}
			// }
		}
	}
	return schedules
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
