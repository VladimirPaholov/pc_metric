package main

import (
	"os"
	"pc_metric/internal/app"
	"pc_metric/internal/db"
	"pc_metric/internal/db/migrations"
	"pc_metric/internal/db/repository"
	"pc_metric/internal/logger"
	"pc_metric/internal/service"
	"time"

	"github.com/joho/godotenv"
)

func main() {

	var (
		workTime,
		metricInterval time.Duration
	)
	logger.InitSysLogger()
	logger.SysLogger.Info("service started")

	if err := godotenv.Load(); err != nil {
		logger.SysLogger.Error("file .env not found", "error", err)
		os.Exit(1)
	}
	logger.SysLogger.Info(".env loaded successfully")

	logger.SysLogger.Info("connecting to database")
	db, err := db.InitDB()
	if err != nil {
		logger.SysLogger.Error("database connection failed", "error", err)
		os.Exit(1)
	}
	logger.SysLogger.Info("connected to database is successfully")
	defer db.Close()
	repo := repository.NewRepository(db)
	userService := service.NewUserService(repo)

	logger.SysLogger.Info("running migrations")
	if err := migrations.RunMigration(); err != nil {
		logger.SysLogger.Error("database migration failed", "error", err)
		os.Exit(1)
	}
	logger.SysLogger.Info("migration completed successfully")

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
