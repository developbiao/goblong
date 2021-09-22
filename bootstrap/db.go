package bootstrap

import (
	"goblong/pkg/model"
	"time"
)

func SetupDB() {
	// Connect to database pool
	db := model.ConnectDB()

	// Command print db request information
	sqlDB, _ := db.DB()

	// Set maximum connections
	sqlDB.SetMaxOpenConns(100)

	// Set Maximum idles
	sqlDB.SetMaxIdleConns(25)

	// Set each connection expired time
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

}
