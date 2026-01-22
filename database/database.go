package database

import (
	"go-mysql-gorm-gin/model"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	uri := "root:@tcp(localhost:3306)/users_db?parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(uri), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to Connect to MySql", err)
	}
	if err := db.AutoMigrate(&model.User{}); err != nil {
		log.Fatal("Migration Failed", err)
	}
	DB = db
}
