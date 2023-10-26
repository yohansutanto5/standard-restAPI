package db

import (
	"app/cmd/config"
	"fmt"
	"log"
	"testing"

	"gorm.io/gorm"
)

func TestMigration(t *testing.T) {
	Migration()
}

var dbg *gorm.DB

func TestMain(m *testing.M) {
	configs := config.Load("dev")
	var err error
	dbg, err = GormInit(configs)
	if err != nil {
		log.Fatal("asd")
	} else {
		m.Run()
	}

}
func TestGorm(t *testing.T) {
	configs := config.Load("dev")
	dbg, err := GormInit(configs)
	if err != nil {
		t.Fail()
	}
	err = dbg.AutoMigrate(&Product{})
	if err != nil {
		fmt.Println(err.Error())
		t.FailNow()
	}
	dbg.Create(&Product{Code: "D42", Price: 100})
	dbg.Create(&Product{Name: "D42123123123123123123", Price: 100})

}

type Product struct {
	ID     int     `gorm:"primaryKey"`
	Name   string  `gorm:"unique;size:12"`
	Price  float64 `gorm:"size:12;type:int"`
	Active bool
	Code   string `gorm:"index"`
	CartID int
}

type Cart struct {
	ID      int `gorm:"primaryKey"`
	Product []Product
	Total   int
}

func TestGormRelation(t *testing.T) {
	err := dbg.AutoMigrate(&Cart{}, &Product{})
	if err != nil {
		fmt.Println(err.Error())
		t.FailNow()
	}
}

func TestGormRelationCreate(t *testing.T) {
	res := dbg.Create(&Cart{
		ID: 1,
		Product: []Product{
			{ID: 1, Name: "hy"}, {ID: 2, Name: "lpl"},
		},
		Total: 123,
	})
	if res.Error != nil {
		t.FailNow()
		fmt.Println(res.Error.Error())
	}
}

// User has and belongs to many languages, use `user_languages` as join table
type User struct {
	gorm.Model
	Languages []*Language `gorm:"many2many:user_languages;"`
}

type Language struct {
	gorm.Model
	Name  string
	Users []*User `gorm:"many2many:user_languages;"`
}

func TestGormRelation2(t *testing.T) {
	err := dbg.AutoMigrate(&Language{}, &User{})
	if err != nil {
		fmt.Println(err.Error())
		t.FailNow()
	}
}

func TestGormRelationCreate2(t *testing.T) {
	res := dbg.Create(&Cart{
		ID: 1,
		Product: []Product{
			{ID: 1, Name: "hy"}, {ID: 2, Name: "lpl"},
		},
		Total: 123,
	})
	if res.Error != nil {
		t.FailNow()
		fmt.Println(res.Error.Error())
	}
}
