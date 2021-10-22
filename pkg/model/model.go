package model

import (
	"fmt"
	c "goblong/pkg/config"
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
	var (
		host     = c.GetString("database.mysql.host")
		port     = c.GetString("database.mysql.port")
		database = c.GetString("database.mysql.database")
		username = c.GetString("database.mysql.username")
		password = c.GetString("database.mysql.password")
		charset  = c.GetString("database.mysql.charset")
	)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=%s",
		username, password, host, port, database, charset, "Local")

	config := mysql.New(mysql.Config{
		DSN: dsn,
	})

	// Prepare connection pool
	DB, err = gorm.Open(config, &gorm.Config{
		Logger: gormlogger.Default.LogMode(gormlogger.Warn),
	})

	logger.LogError(err)
	return DB
}
