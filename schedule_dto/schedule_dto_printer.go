package schedule_dto

import (
	fmt "fmt"
	"strconv"

	"github.com/golang/protobuf/ptypes"
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
	el := schedule.GetEventsList()

	logger.Log("Printing schedule id " + strconv.Itoa(int(schedule.GetID())))
	for _, event := range el {
		logger.Log(GetEventData(*event))
	}
}

func GetEventData(e Event) string {
	actualStartingTime, _ := ptypes.Timestamp(e.EventTime.ActualStartingTime)
	return fmt.Sprintf("'%s' at %s",
		e.EventName,
		actualStartingTime.Format("2006-01-02 15:04"))
}
