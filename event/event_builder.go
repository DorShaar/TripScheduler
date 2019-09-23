package event

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
	"unicode"

	"gopkg.in/yaml.v2"
)

type eventInterface struct {
	Name                string
	Location            string
	Time_Range          []string
	Duration            string
	Precaution_Duration string
}

func BuildEventFronYAML(yamlPath string) (Event, error) {
	var eventResult Event
	// Reads yaml file
	yamlData, err := ioutil.ReadFile(yamlPath)
	if err != nil {
		fmt.Printf("Error reading yaml file. Error:   #%v ", err)
		return eventResult, err
	}

	eventData := eventInterface{}
	err = yaml.Unmarshal(yamlData, &eventData)
	if err != nil {
		fmt.Printf("Error unmarsheling yaml data. Error:   #%v ", err)
		return eventResult, err
	}

	duration, _ := time.ParseDuration(eventData.Duration)
	precautionDuration, _ := time.ParseDuration(eventData.Precaution_Duration)

	eventResult.EventName = eventData.Name
	eventResult.Location = eventData.Location
	eventResult.TimeRangesList = parseTimeRanges(eventData.Time_Range)
	eventResult.EventTime.Duration = duration
	eventResult.EventTime.PrecautionDuration = precautionDuration

	return eventResult, err
}

func parseTimeRanges(timeRangesStrings []string) []TimeRange {
	timeRanges := make([]TimeRange, 0)

	for _, currentTimeRange := range timeRangesStrings {
		var rangeTimeOnly string
		var day string

		index := strings.IndexByte(currentTimeRange, ' ')

		if index == -1 {
			// Must be the case of 07:00
			rangeTimeOnly = currentTimeRange
		} else if unicode.IsDigit(rune(currentTimeRange[index-1])) {
			// In case of format 07:00 - 19:00
			rangeTimeOnly = currentTimeRange
		} else {
			// In case of format Sunday: 07:00 - 19:00
			day = currentTimeRange[:index]
			rangeTimeOnly = currentTimeRange[index:]
		}

		timeRange := parseTimeRange(rangeTimeOnly)
		timeRange.weekday = day
		timeRanges = append(timeRanges, timeRange)
	}

	return timeRanges
}

// Get string in the format of "08:00 - 19:00" and returns starting time as 08:00 and ending
// time as 19:00.
// In case there is no ending time such as 08:00, only starting time will be updated.
func parseTimeRange(timeRangesStr string) TimeRange {
	timeLayout := "15:04"
	var endingTime time.Time

	timeRangeSlice := strings.Split(timeRangesStr, "-")

	startingTime, _ := time.Parse(timeLayout, strings.TrimSpace(timeRangeSlice[0]))
	if len(timeRangeSlice) > 1 {
		endingTime, _ = time.Parse(timeLayout, strings.TrimSpace(timeRangeSlice[1]))
	}

	return CreateRange(startingTime, endingTime)
}

func BuildEventByParameters(
	name string,
	location string,
	timeRangesStrings []TimeRange,
	duration time.Duration,
	precautionDuration time.Duration,
	actualStartingTime time.Time) (eventResult Event) {

	eventResult.EventName = name
	eventResult.Location = location
	eventResult.TimeRangesList = timeRangesStrings
	eventResult.EventTime.Duration = duration
	eventResult.EventTime.PrecautionDuration = precautionDuration
	eventResult.EventTime.actualStartingTime = actualStartingTime

	return eventResult
}
