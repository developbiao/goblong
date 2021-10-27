package config

import (
	"fmt"
	"goblong/pkg/config"
)

func init() {
	fmt.Println("App Config loading")
	config.Add("app", config.StrMap{
		// App name
		"name": config.Env("APP_NAME", "GoBlog"),

		// Current environment
		"env": config.Env("APP_NAME", "production"),

		// Debug mode
		"debug": config.Env("APP_DEBUG", false),

		// App port
		"port": config.Env("APP_PORT", "3000"),

		// App key for cookie gorilla/session
		"key": config.Env("APP_KEY", "66333999dcf9ea060a0a6532b166da32f304af0de"),

		// app url
		"url": config.Env("APP_URL", "http://localhost:3000"),
	})
	fmt.Println("App Config Done!")
}
