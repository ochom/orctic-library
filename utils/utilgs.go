package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
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

	var letterRunes = []rune("0123456789")
	b := make([]rune, size)

	for i := range b {
		index := func() int {
			bigN, err := rand.Int(rand.Reader, big.NewInt(9))
			if err != nil {
				return 0
			}
			return int(bigN.Int64())
		}()
		b[i] = letterRunes[index]
	}

	return string(b)
}
