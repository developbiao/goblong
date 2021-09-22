package types

import (
	"goblong/pkg/logger"
	"strconv"
)

// Convert int64 to string
func Int64ToString(num int64) string {
	return strconv.FormatInt(num, 10)
}

// Convert string to integer
func StringToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		logger.LogError(err)
	}
	return i
}
