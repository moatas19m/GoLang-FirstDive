package config

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Connect() {
	dsn := "your-connection-string"
	log.Printf("Connecting to database %s", dsn)
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
