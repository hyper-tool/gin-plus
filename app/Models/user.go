package Models

import (
	"gin-plus/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Name string
	Age  int
}

func index() {

}

func Create(username string, age int) bool {
	user := User{Name: username, Age: age}
	db := config.Connection()
	db.Create(&user)
	return db.NewRecord(user)
}

func show() {

}
func destroy() {

}
