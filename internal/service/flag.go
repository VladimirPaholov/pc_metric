package service

import (
	"flag"
	"fmt"
	"os"
	"time"
)

type FlagConfig struct {
	CustomWorkTime       time.Duration
	CustomMetricInterval time.Duration
}

func ParseFlags(d time.Duration, i time.Duration) FlagConfig {

	work := flag.Duration("d", d, "Default time for working app - 1 min")

	metric := flag.Duration("i", i, "Default time for metrics monitoring - 1 second")

	flag.Parse()

	if work == nil || metric == nil {
		fmt.Printf("Return nil pointer %v:%v", work, metric)
		os.Exit(1)
	}

	if *work <= 0 {
		fmt.Println("Duration work time must be > 0")
		os.Exit(1)
	}

	if *metric <= 0 {
		fmt.Println("Duration get metric must be > 0")
		os.Exit(1)
	}
	return FlagConfig{CustomWorkTime: *work, CustomMetricInterval: *metric}
}
