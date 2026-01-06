package repository

import "database/sql"

type DataBase struct {
	db *sql.DB
}

func New(db *sql.DB) *DataBase {
	return &DataBase{db: db}
}
