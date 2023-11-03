package db

import (
	"app/cmd/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	connection *gorm.DB
}

func NewDatabase(config config.Configuration) *Database {
	connectionString := "host=localhost user=app password=app dbname=app search_path=app port=5432 sslmode=disable search_path=app"
	connection, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &Database{
		connection: connection,
	}
}

func ConnectDB(username, password, host, dbName string) (*sqlx.DB, error) {
	connectionString := "user=" + username + " password=" + password + " host=" + host + " dbname=" + dbName + " sslmode=disable" + " search_path=app"
	db, err := sqlx.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	// Set the maximum number of open and idle connections
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func GormInit(config config.Configuration) (*gorm.DB, error) {
	connectionString := "host=localhost user=app password=app dbname=app search_path=app port=5432 sslmode=disable search_path=app"
	return gorm.Open(postgres.Open(connectionString), &gorm.Config{})
}
