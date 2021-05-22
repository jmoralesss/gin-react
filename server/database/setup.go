package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Conn *gorm.DB

const DB_USERNAME = "array_user"
const DB_PASSWORD = "password123"
const DB_NAME = "array"
const DB_HOST = "127.0.0.1"
const DB_PORT = "3306"

func ConnectDatabase() {
	var err error
	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	Conn = db
}
