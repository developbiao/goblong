package config

import (
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"goblong/pkg/logger"
)

// Viper instance
var Viper *viper.Viper

// Config map
type StrMap map[string]interface{}

// When import init() automatic load
func init() {
	// 1. Initialization Viper
	Viper = viper.New()

	// 2. Set config name
	Viper.SetConfigName(".env")

	// 3. Set config type
	// Support "json", "yaml", "yml", "env", "dotenv", "props", "prop", "properties"
	Viper.SetConfigType("env")

	// 4. Set Variable path related with main.go some path
	Viper.AddConfigPath(".")

	// 5. Read config from project root
	err := Viper.ReadInConfig()
	logger.LogError(err)

	// 6. Set environment prefix
	Viper.SetEnvPrefix("appenv")

	// 7. When use Viper.Get() priority read it
	Viper.AutomaticEnv()

}

// Env read environment support default value
func Env(envName string, defaultValue ...interface{}) interface{} {
	if len(defaultValue) > 0 {
		return Get(envName, defaultValue[0])
	}
	return Get(envName)
}

// Get config option support use dot example: app.name
func Get(path string, defaultValue ...interface{}) interface{} {
	if !Viper.IsSet(path) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return nil
	}
	return Viper.Get(path)
}

// Add config option
func Add(name string, configuration map[string]interface{}) {
	Viper.Set(name, configuration)
}

// Get type of string config
func GetString(path string, defaultValue ...interface{}) string {
	return cast.ToString(Get(path, defaultValue...))
}

// Get type of integer config
func GetInt(path string, defaultValue ...interface{}) int {
	return cast.ToInt(Get(path, defaultValue))
}

// Get type of integer 64 config
func GetInt64(path string, defaultValue ...interface{}) int64 {
	return cast.ToInt64(Get(path, defaultValue))
}

// Get type of Uint config
func GetUint(path string, defaultValue ...interface{}) uint {
	return cast.ToUint(Get(path, defaultValue))
}

// Get type of boolean config
func GetBool(path string, defaultValue ...interface{}) bool {
	return cast.ToBool(Get(path, defaultValue))
}
