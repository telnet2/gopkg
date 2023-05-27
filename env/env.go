package env

import (
	"os"
	"strconv"
)

// Get returns an env with a given key and returns `def` if its empty.
func Get(key string, def string) string {
	val, ok := os.LookupEnv(key)
	if !ok || val == "" {
		return def
	}
	return val
}

// GetInt returns an int value of an env variable.
func GetInt(key string, def int) int {
	val, ok := os.LookupEnv(key)
	if !ok || val == "" {
		return def
	}
	intVal, err := strconv.Atoi(val)
	if err != nil {
		return def
	}
	return intVal
}

// IsSet returns a bool value of an env variable.
func IsSet(key string) bool {
	val, ok := os.LookupEnv(key)
	result := ok && val != ""
	return result
}
