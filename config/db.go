package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/perkantoran"))

	if err != nil {
		panic(err)
	}
}
