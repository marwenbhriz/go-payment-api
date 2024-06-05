package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:root@tcp(localhost:3307)/payments?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Company{})
	database.AutoMigrate(&User{})
	database.AutoMigrate(&Invoice{})

	DB = database
}
