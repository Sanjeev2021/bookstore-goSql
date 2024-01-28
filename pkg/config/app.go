package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

)

var (
	db *gorm.DB // THIS WILL CONNECT TO THE DATABASE
)

func Connect() {
	database, err := gorm.Open("mysql", "root:root@/bookstore?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = database // here we are establishing the connection
}

func GetDB() *gorm.DB {
	return db
}
