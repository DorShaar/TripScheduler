package schedule_dto

import (
	"trip_scheduler/event"
	"trip_scheduler/schedule"

	"github.com/golang/protobuf/ptypes"
)

func CreateDTOSchedule(schedule schedule.Schedule) *Schedule {
	scheduleDTO := &Schedule{}
	scheduleDTO.ID = int32(schedule.Id())

	eventsList := schedule.GetEventsListCopy()
	eventsListDTO := make([]*Event, 0)
	for _, event := range eventsList {
		eventsListDTO = append(eventsListDTO, createDTOEvent(event))
	}

	scheduleDTO.EventsList = eventsListDTO
	return scheduleDTO
}

func createDTOEvent(event event.Event) *Event {
	eventDTO := &Event{}
	eventDTO.EventName = event.EventName
	eventDTO.Location = event.Location
	eventDTO.EventTime = createDTOEventTime(event.EventTime)

	return eventDTO
}

func createDTOEventTime(eventTime event.EventTime) *EventTime {
	eventTimeDTO := &EventTime{}
	eventTimeDTO.DurationInSec = int32(eventTime.Duration.Seconds())
	eventTimeDTO.PrecautionDurationInSec = int32(eventTime.PrecautionDuration.Seconds())

	actualStartingTime, _ := ptypes.TimestampProto(eventTime.ActualStartingTime())
	eventTimeDTO.ActualStartingTime = actualStartingTime
	return eventTimeDTO
}
