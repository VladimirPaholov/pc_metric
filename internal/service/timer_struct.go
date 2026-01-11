package service

import (
	"fmt"
	"os"
	"time"
)

type TimeCfg struct {
	DefaultTimeWork      time.Duration
	DefaultTimeGetMetric time.Duration
}

func NewTimeCfg() *TimeCfg {
	return &TimeCfg{
		DefaultTimeWork:      ParseTimeCfg("DEFAULT_TIME_WORK", 1*time.Minute),
		DefaultTimeGetMetric: ParseTimeCfg("DEFAULT_TIME_GET_METRIC", 1*time.Second),
	}
}

func ParseTimeCfg(env string, defaultValue time.Duration) time.Duration {
	value := os.Getenv(env)

	if value == "" {
		fmt.Printf("ENV %s is not set, using default: %v\n", env, defaultValue)
		return defaultValue
	}

	d, err := time.ParseDuration(value)
	if err != nil {
		fmt.Printf("Failed to parse %s=%s,. Set default time %v: %v\n", env, value, defaultValue, err)

		return defaultValue

	}
	return d
}
