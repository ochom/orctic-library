package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

// GetEnv ...
func GetEnv[T comparable](key string, defaultValue T) (res T) {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	val, err := ParseData[T]([]byte(value))
	if err != nil {
		return defaultValue
	}
	return *val
}

// MustGetEnv ...
func MustGetEnv[T comparable](key string) (res T) {
	value, ok := os.LookupEnv(key)
	if !ok {
		NewLogger().Error(fmt.Sprintf("Environment variable %s is not set", key))
	}

	val, err := ParseData[T]([]byte(value))
	if err != nil {
		NewLogger().Error(fmt.Sprintf("Environment variable %s is wrongly set", key))
	}
	return *val
}

// ParseData ...
func ParseData[T any](data []byte) (*T, error) {
	var t T
	if err := json.Unmarshal(data, &t); err != nil {
		NewLogger().Error(fmt.Sprintf("Error parsing data: %s", err))
		return nil, err
	}

	return &t, nil
}
