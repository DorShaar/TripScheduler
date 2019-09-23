package time_extended

import (
	"testing"
	"time"
)

func Test_Parse_Success(t *testing.T) {
	weekdayParser := WeekdayParser{}
	if weekdayParser.Parse("Monday") != time.Monday {
		t.Errorf("WeekdayParser Prase did not perform well")
	}
}

func Test_Parse_Panic(t *testing.T) {
	weekdayParser := WeekdayParser{}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Parse did not paniced while it should be")
		}
	}()

	weekdayParser.Parse("Funday")
}

func Test_IsInitialized_True(t *testing.T) {
	weekdayParser := WeekdayParser{}
	if weekdayParser.Parse("Monday") != time.Monday {
		t.Errorf("WeekdayParser Prase did not perform well")
	}

	if !weekdayParser.isInitialized {
		t.Errorf("WeekdayParser is marked as not initialized while it should be")
	}
}
