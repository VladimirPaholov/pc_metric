package main

import (
	"fmt"
	"os"
	"pc_metric/internal/lifecycle"
	"pc_metric/internal/service"
	"pc_metric/repository"
	"time"
)

func main() {

	var (
		workTime,
		metricInterval time.Duration
	)

	service.GetENV()

	db, err := repository.DBconnection()
	if err != nil {
		os.Exit(1)
	}
	defer db.DB.Close()

	if err := db.RunMigration(); err != nil {
		fmt.Println("Migration error:", err)
		os.Exit(1)
	}
	t := service.NewTimeStruct()

	cfg := service.ParseFlags(t.DefaultTimeWork, t.DefaultTimeMetric)
	if cfg.CustomWorkTime > 0 {
		workTime = cfg.CustomWorkTime
		metricInterval = cfg.CustomMetricInterval
	} else {
		workTime = t.DefaultTimeWork
		metricInterval = t.DefaultTimeMetric
	}

	lifecycle.Start(workTime, metricInterval, db)
}
