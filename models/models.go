package models

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(sqlite.Open("demo-api.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln("无法连接数据库:", err.Error())
	}

	db.AutoMigrate(&Message{}, &Keyword{})
}
