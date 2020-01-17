package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var DB *gorm.DB

func init() {
	db, err := gorm.Open("sqlite3", "my.db")
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&User{})
	DB = db
}
