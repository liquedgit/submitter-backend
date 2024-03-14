package utils

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetConnection() {
	dsn := "root:@tcp(127.0.0.1:3306)/challdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}
	db.AutoMigrate()
}
