package main

import (
	"fmt"
	"os"
	"pc_metric/internal/app"
	"pc_metric/internal/db"
	"pc_metric/internal/db/migrations"
	"pc_metric/internal/db/repository"
	"pc_metric/internal/service"
	"time"

	"github.com/joho/godotenv"
)

func main() {

	var (
		workTime,
		metricInterval time.Duration
	)

	if err := godotenv.Load(); err != nil {
		fmt.Println("File .env not found")
		os.Exit(1)
	}

	db, err := db.InitDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	repo := repository.NewRepository(db)
	userService := service.NewUserService(repo)

	if err := migrations.RunMigration(); err != nil {
		fmt.Println("Migration error:", err)
		os.Exit(1)
	}
	t := service.NewTimeCfg()

	cfg := service.ParseFlags(t.DefaultTimeWork, t.DefaultTimeGetMetric)
	if cfg.CustomWorkTime > 0 {
		workTime = cfg.CustomWorkTime
		metricInterval = cfg.CustomMetricInterval
	} else {
		workTime = t.DefaultTimeWork
		metricInterval = t.DefaultTimeGetMetric
	}

	app.Start(workTime, metricInterval, userService)
}
