package config

import (
	"beritaalta/model/news"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var Db *gorm.DB

func Init() {
	initDatabase()
	migrate()
}

// pake .env
//func initDatabase() {
//	viper.SetConfigFile(".env")
//	if err := viper.ReadInConfig(); err != nil {
//		panic(err)
//	}
//
//	username := viper.Get("DB_USERNAME")
//	password := viper.Get("DB_PASSWORD")
//	host := viper.Get("DB_HOST")
//	port := viper.Get("DB_PORT")
//	name := viper.Get("DB_NAME")
//
//	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
//		username,
//		password,
//		host,
//		port,
//		name,
//	)
//	var err error
//	Db, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
//	if err != nil {
//		log.Fatalf("Failed to connect to database: %s", err.Error())
//	}
//}

func initDatabase() {
	username := "root"
	password := "mysql1234"
	host := "35.187.224.18"
	port := "33062"
	name := "beritabaru"

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		username,
		password,
		host,
		port,
		name,
	)
	var err error
	Db, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err.Error())
	}
}

func migrate() {
	Db.AutoMigrate(&news.News{})
}
