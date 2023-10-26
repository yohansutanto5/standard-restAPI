package db

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
)

func Migration() {
	// Create a new migration instance
	dbURL := "postgres://postgres@localhost:5432/app?sslmode=disable&search_path=app"
	m, err := migrate.New(
		"file://"+"/home/yohan/workspace/db/migration/",
		dbURL,
	)
	if err != nil {
		log.Fatalf("Failed to create migration instance: %v", err)
	}

	// Run the specified migration action
	err = m.Up()

	if err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Migration failed: %v", err)
	}

	// Get the current migration version
	version, dirty, err := m.Version()
	if err != nil {
		log.Fatalf("Failed to get migration version: %v", err)
	}
	fmt.Printf("Current migration version: %v (dirty: %v)\n", version, dirty)
}
