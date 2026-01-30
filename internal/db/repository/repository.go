package repository

import (
	"database/sql"
	"time"
)

type Repository struct {
	db *sql.DB
}

type MetricRepository interface {
	AddMetricDB(timestamp time.Time, message string) error
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}
