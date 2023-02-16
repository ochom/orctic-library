package utils

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
)

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

// ParseData ...
func ParseData[T any](data []byte) (*T, error) {
	var t T
	if err := json.Unmarshal(data, &t); err != nil {
		NewLogger().Error(fmt.Sprintf("Error parsing data: %s", err))
		return nil, err
	}

	return &t, nil
}
