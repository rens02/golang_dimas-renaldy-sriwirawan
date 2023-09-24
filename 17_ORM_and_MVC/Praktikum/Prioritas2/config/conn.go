package config

import (
	"Prioritas2/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var (
	DB *gorm.DB
)

func init() {
	InitDB()
	InitialMigration()
}

func InitDB() {
	username := "rens02"
	password := "mysql1234"
	host := "35.187.224.18"
	port := "3306"
	name := "beritabaru"

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		username,
		password,
		host,
		port,
		name,
	)
	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err.Error())
	}
}

func InitialMigration() {
	DB.AutoMigrate(&models.User{}, &models.Books{})
}
