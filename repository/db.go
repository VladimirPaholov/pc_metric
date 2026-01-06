package repository

import (
	"database/sql"
	"fmt"
	"os"

	//"time"
	"embed"

	_ "github.com/lib/pq"
)

//go:embed migrations/pg/*.sql
var sqlFiles embed.FS

func DBconnection() (*sql.DB, error) {

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	sslMode := os.Getenv("DB_SSLMODE")

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		user,
		password,
		host,
		port,
		dbName,
		sslMode,
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("sql.Open error: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("data base ping error: %w", err)
	}

	fmt.Println("Connected to data base - successfull!")

	return db, nil
}

// переделать как методы
//
//	func NewTable(db *sql.DB) error {
//		_, err := db.Exec(`
//		CREATE TABLE IF NOT EXISTS logs_metric (
//		id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
//		created_at TIMESTAMPTZ NOT NULL,
//		message TEXT NOT NULL
//
// );`)
//
//		return err
//	}
func (d *DataBase) CreateTable() error {
	query, err := sqlFiles.ReadFile("migrations/pg/1-logs_metric.sql")
	if err != nil {
		return fmt.Errorf("Read sql file error: %w", err)
	}
	_, err = d.db.Exec(string(query))
	return nil
}

// func InsertLogMetric(db *sql.DB, createdAt time.Time, message string) error {
// 	_, err := db.Exec(
// 		`INSERT INTO logs_metric (created_at, message)
// 		VALUES ($1, $2)`,
// 		createdAt,
// 		message,
// 	)
// 	if err != nil {
// 		return fmt.Errorf("Insert log error: %w", err)
// 	}
// 	return nil
// }
