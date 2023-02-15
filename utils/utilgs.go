package utils

import (
	"crypto/rand"
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
