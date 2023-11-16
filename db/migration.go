package db

import (
	"app/cmd/config"
	"app/pkg/log"
	"fmt"

	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
)

func Migration(config *config.Configuration) {
	// Create a new migration instance
	dbURL := fmt.Sprintf("postgress://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s",
		config.Db.Username, config.Db.Password, config.Db.Host, config.Db.Port, config.Db.Database, config.Db.Schema)

	m, err := migrate.New(
		"file://"+"/home/yohan/workspace/db/migration/",
		dbURL,
	)
	if err != nil {
		log.Fatal("Failed to create migration instance")
	}

	// Run the specified migration action
	err = m.Up()

	if err != nil && err != migrate.ErrNoChange {
		log.Fatal("Migration failed")
	}

	// Get the current migration version
	version, dirty, err := m.Version()
	if err != nil {
		log.Fatal("Failed to get migration version")
	}

	log.System(fmt.Sprintf("Current migration version: %s (dirty: %s)\n", version, dirty))
}
