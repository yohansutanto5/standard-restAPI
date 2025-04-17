package dbtest

import (
	"app/cmd/config"
	"app/db"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var dbg *gorm.DB
var ds *db.DataStore
var configs config.Configuration

func TestMain(m *testing.M) {
	mode := os.Getenv("GIN_MODE")
	configfolder := os.Getenv("config")
	configs = config.Load(mode, configfolder)
	var err error
	ds = db.NewDatabase(configs)
	dbg = ds.Db
	if err != nil {
		log.Fatal("asd")
	} else {
		m.Run()
	}

}

func TestMigration(t *testing.T) {
	db.Migration(&configs, false)
}

func TestGormMigration(t *testing.T) {
	err := ds.Db.AutoMigrate()
	assert.Nil(t, err)
}
