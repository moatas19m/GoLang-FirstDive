package config

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func Connect() {
	dsn := os.Getenv("DB_CONNECTION_STRING")
	d, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = d
	log.Println("Connected to database successfully")
}
func GetDB() *gorm.DB {
	return DB
}