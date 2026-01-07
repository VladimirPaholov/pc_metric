package service

import (
	"fmt"
	"os"
	"time"
)

type TimeStruct struct {
	DefaultTimeWork   time.Duration
	DefaultTimeMetric time.Duration
}

func NewTimeStruct() *TimeStruct {
	t := &TimeStruct{}

	t.ReadTimeValue()
	t.ReadTimeGetMetrics()
	return t
}

func (t *TimeStruct) ReadTimeValue() {
	readValue := os.Getenv("DEFAULT_TIME_WORK")

	d, err := time.ParseDuration(readValue)
	if err != nil {
		fmt.Printf("Parse error time from .env: %v. Set default time 1m\n", err)
		d = 1 * time.Minute
	}
	t.DefaultTimeWork = d
}

func (t *TimeStruct) ReadTimeGetMetrics() {
	readValue := os.Getenv("DEFAULT_TIME_GET_METRIC")

	d, err := time.ParseDuration(readValue)
	if err != nil {
		fmt.Printf("Error parse getting time metric: %v. Set default time 1s\n", err)
		d = 1 * time.Second
	}
	t.DefaultTimeMetric = d
}
