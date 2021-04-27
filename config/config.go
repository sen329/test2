package config

import (
	"../structs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DBInit create connection to database
func DBInit() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/testgolangdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(structs.User{})
	return db
}
