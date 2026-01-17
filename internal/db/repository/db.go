package repository

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func DBconnection() (*Repository, error) {

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

	return &Repository{DB: db, Dsn: dsn}, nil
}

func (d *Repository) RunMigration() error {

	cwd, err := os.Getwd()

	if err != nil {
		return fmt.Errorf("Error get current dir: %v", err)
	}

	pathWorkDir := os.Getenv("WORK_DIR_PATH")

	if pathWorkDir == "" {
		pathWorkDir = cwd
	}

	pathToMigration := filepath.Join(pathWorkDir, "migrations/pg")

	absMigrationPath, err := filepath.Abs(pathToMigration)

	if err != nil {
		return fmt.Errorf("Failed to get absolute path for %s: %w", absMigrationPath, err)
	}

	migration, err := migrate.New("file://"+filepath.ToSlash(absMigrationPath), d.Dsn)
	if err != nil {
		return err
	}
	defer migration.Close()
	errM := migration.Up()
	if errM != nil && errM != migrate.ErrNoChange {
		return errM
	}
	return nil
}

func (d *Repository) InsertData(createdAt time.Time, message string) error {
	_, err := d.DB.Exec(
		`INSERT INTO logs_metric (created_at, message)
 		VALUES ($1, $2)`,
		createdAt,
		message,
	)
	if err != nil {
		return fmt.Errorf("Insert log error: %w", err)
	}
	return nil
}
