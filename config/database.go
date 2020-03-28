package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Database struct {
	Connection func()
}

var (
	username  = "root"
	password  = ""
	dbname    = "aa"
	charset   = "aa"
	parseTime = "aa"
	loc       = "aa"
)

func Connection() *gorm.DB {
	db, _ := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/localhost?charset=utf8mb4&parseTime=True&loc=Local")
	db.LogMode(true)
	//defer db.Close()
	return db
}
