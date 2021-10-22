package config

import "goblong/pkg/config"

func init() {
	config.Add("session", config.StrMap{
		// default mode current is cookie
		"default": config.Env("SESSION_DRIVER", "cookie"),
		// Session cookie name
		"session_name": config.Env("SESSION_NAME", "goblog-session"),
	})
}
