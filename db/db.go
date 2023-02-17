package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetMysqlConnection() string {
	dbPort := os.Getenv("DB_PORT")
	dbUsername := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")

	MysqlConnection := fmt.Sprintf("%s@tcp(127.0.0.1:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUsername, dbPort, dbName)
	return MysqlConnection
}

func ConnectToDB() {
	var err error
	dsn := GetMysqlConnection()
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	
	if err != nil {
		log.Fatal("error connecting to database : " + err.Error())
	}
}