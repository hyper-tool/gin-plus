package main

import (
	"gin-plus/routes"
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

func main() {
	router := routes.Api()
	router.Run()
}
