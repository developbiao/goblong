package database

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"goblong/pkg/logger"
	"time"
)

// DB database object
var DB *sql.DB

func Initialize() {
	initDB()
	createTables()
}

func initDB() {
	var err error
	config := mysql.Config{
		User:                 "homestead",
		Passwd:               "secret",
		Addr:                 "192.168.56.38",
		Net:                  "tcp",
		DBName:               "goblog",
		AllowNativePasswords: true,
	}

	// Prepare database pool
	DB, err = sql.Open("mysql", config.FormatDSN())
	// fmt.Printf("DSN:%v\n", config.FormatDSN())
	logger.LogError(err)

	// Set maximum connections
	DB.SetMaxIdleConns(25)

	// Set maximum connection idle time
	DB.SetMaxIdleConns(25)

	// Set each connection expire time
	DB.SetConnMaxLifetime(5 * time.Minute)

	// Connection to database
	err = DB.Ping()
	logger.LogError(err)

}

// Create tables
func createTables() {
	createArticlesSQL := `CREATE TABLE IF NOT EXISTS articles(
	id BIGINT(20) PRIMARY KEY AUTO_INCREMENT NOT NULL,
	title VARCHAR(255) COLLATE utf8mb4_unicode_ci NOT NULL,
	body longtext COLLATE utf8mb4_unicode_ci
);`

	_, err := DB.Exec(createArticlesSQL)
	logger.LogError(err)
}
