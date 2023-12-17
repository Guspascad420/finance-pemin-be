package database

import (
	"finance-be/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var Db *gorm.DB
var err error

func Connect() {
	dbUrl := os.Getenv("DB_URL")
	dsn := dbUrl
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Cannot connect to DB")
	} else {
		log.Println("Connected to Database!")
	}
}

func Migrate() {
	Db.AutoMigrate(&models.User{})
	Db.AutoMigrate(&models.Balance{})
	Db.AutoMigrate(&models.Transaction{})

	log.Println("Database Migration Completed!")
}
