package repository

import "database/sql"

type DataBase struct {
	DB  *sql.DB
	Dsn string
}
