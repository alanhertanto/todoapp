package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// SetupDB : initializing mysql database
func SetupDB() *gorm.DB {
	USER := "root"
	PASS := "root"
	HOST := "localhost"
	PORT := "3000"
	DBNAME := "blogpost"
	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	db, err := gorm.Open(mysql.Open(URL))
	if err != nil {
		panic(err.Error())
	}
	return db
}
