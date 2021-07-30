package db

import (
	"gomodtest/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db  *gorm.DB
	err error
)

func Init() {
	db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=gomodtest password=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}
	autoMigration()
	user := models.User{
		Id:    1,
		Name:  "aa",
		Posts: []models.Post{{ID: 1, Content: "t1"}},
	}
	db.Create(&user)
}

func GetDB() *gorm.DB {
	return db
}

func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

func autoMigration() {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Post{})
}
