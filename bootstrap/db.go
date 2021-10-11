package bootstrap

import (
	"goblong/app/models/article"
	"goblong/app/models/user"
	"goblong/pkg/model"
	"gorm.io/gorm"
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

	// Create table and maintain
	migration(db)

}

func migration(db *gorm.DB) {
	// Automatic migration
	db.AutoMigrate(
		&user.User{},
		&article.Article{},
	)

}
