package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DbMysql *gorm.DB

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
//	DbMysql, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
//	if err != nil {
//		log.Fatalf("Failed to connect to database: %s", err.Error())
//	}
//}

func initDatabase() {
	username := "root"
	host := "localhost"
	port := "3306"
	name := "beritabaru"

	connectionString := fmt.Sprintf("%s:@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		username,
		host,
		port,
		name,
	)
	var err error
	DbMysql, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err.Error())
	}
}
