package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d , err := gorm.Open("mysql","surya:suryapass@tcp(127.0.0.1:3306)/bookdb?charset=utf8&parseTime=True&loc=Local")   // user : password/db?
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB()  *gorm.DB {
	return db
}