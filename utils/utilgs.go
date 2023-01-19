package utils

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

var logger = NewLogger()

// MustGetEnv ...
func MustGetEnv(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		logger.Error(fmt.Sprintf("Environment variable %s is not set", key))
	}
	return val
}

// GetEnvOrDefault ...
func GetEnvOrDefault(key string, defaultValue string) string {
	val, ok := os.LookupEnv(key)
	if ok {
		return val
	}
	return defaultValue
}

// GenerateOTP generates a random  OTP of a given size
func GenerateOTP(size int) string {
	rand.Seed(time.Now().UnixNano())
	var letterRunes = []rune("0123456789")
	b := make([]rune, size)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
