package database

import (
	"io/ioutil"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Conn *gorm.DB

func init() {
	db, err := gorm.Open(sqlite.Open("array.db"), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	file, err := ioutil.ReadFile("database/PostDeploy.sql")

	if err != nil {
		panic(err.Error())
	}

	db.Exec(string(file))

	Conn = db
}
