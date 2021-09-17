package types

import "strconv"

// Convert int64 to string
func Int64ToString(num int64) string {
	return strconv.FormatInt(num, 10)
}
