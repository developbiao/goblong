package model

import (
	"goblong/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
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
	DB, err = gorm.Open(config, &gorm.Config{
		Logger: gormlogger.Default.LogMode(gormlogger.Warn),
	})

	logger.LogError(err)
	return DB
}
