package migrations

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlserver"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var ErrorConnect = errors.New("Failed to connect to SQL Server")
var ErrorCreateDB = errors.New("Failed to create database")
var ErrorCreateMigrationDriver = errors.New("Failed to create migration driver")
var ErrorCreateMigrationInstance = errors.New("Failed to create migration instance")
var ErrorRunMigrations = errors.New("Failed to run migrations")

func RunMigrations() error {
	server := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")

	masterDB, err := connectSqlServer(server, port, user, password, "master")
	if err != nil {
		return err
	}
	defer masterDB.Close()

	database := os.Getenv("DB_NAME")

	err = createNewDataBase(masterDB, database)

	if err != nil {
		return err
	}

	db, err := connectSqlServer(server, port, user, password, database)
	if err != nil {
		return err
	}
	defer db.Close()

	err = migrateUP(db)
	if err != nil {
		return err
	}

	return nil
}
func connectSqlServer(server, port, user, password, database string) (*sql.DB, error) {
	masterConnString := fmt.Sprintf(
		"server=%s;port=%s;user id=%s;password=%s;database=%s",
		server, port, user, password, database)

	var db *sql.DB
	var err error

	for attempt := 1; attempt <= 30; attempt++ {
		db, err = sql.Open("sqlserver", masterConnString)

		if err == nil {
			err = db.Ping()
		}

		if err == nil {
			break
		}

		if db != nil {
			db.Close()
		}

		time.Sleep(2 * time.Second)
	}

	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrorConnect, err)
	}

	return db, nil
}
func createNewDataBase(database *sql.DB, databaseName string) error {
	_, err := database.Exec(`IF DB_ID(N'@databaseName') IS NULL
    BEGIN
        CREATE DATABASE [@databaseName]
    END`,
		sql.Named("databaseName", databaseName))

	if err != nil {
		return fmt.Errorf("%w: %v", ErrorCreateDB, err)
	}

	return nil
}
func migrateUP(database *sql.DB) error {
	driver, err := sqlserver.WithInstance(database, &sqlserver.Config{})
	if err != nil {
		return fmt.Errorf("%w: %v", ErrorCreateMigrationDriver, err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://migrations", "sqlserver", driver)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrorCreateMigrationInstance, err)
	}
	defer m.Close()

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("%w: %v", ErrorRunMigrations, err)
	}

	return nil
}
