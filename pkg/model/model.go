package model

import (
	"fmt"
	"goblong/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB gorm.DB object
var DB *gorm.DB

// Initialize DB

func ConnectDB() *gorm.DB {
	var err error

	config := mysql.New(mysql.Config{
		DSN: "homestead:secret@tcp(192.168.56.38:3306)/goblog?charset=utf8&parseTime=True&loc=Local",
	})

	// Prepare connection pool
	DB, err = gorm.Open(config, &gorm.Config{})

	if err != nil {
		fmt.Println("Connection to db failed", err)
	}
	logger.LogError(err)
	return DB
}
