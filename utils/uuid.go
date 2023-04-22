package utils

import (
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"io"
	"sync/atomic"
	"time"
)

var objectIDCounter = readRandomUint32()

func processUniqueBytes() [5]byte {
	var b [5]byte
	_, err := io.ReadFull(rand.Reader, b[:])
	if err != nil {
		panic(err)
	}

	return b
}

func readRandomUint32() uint32 {
	var b [4]byte
	_, err := io.ReadFull(rand.Reader, b[:])
	if err != nil {
		panic(err)
	}

	return (uint32(b[0]) << 0) | (uint32(b[1]) << 8) | (uint32(b[2]) << 16) | (uint32(b[3]) << 24)
}

func putUint24(b []byte, v uint32) {
	b[0] = byte(v >> 16)
	b[1] = byte(v >> 8)
	b[2] = byte(v)
}

func generateUUID() [12]byte {

	var b [12]byte

	// Convert the time to a byte array
	binary.BigEndian.PutUint32(b[0:4], uint32(time.Now().Unix()))

	// Generate 5 random bytes
	unique := processUniqueBytes()
	copy(b[4:9], unique[:])

	// Generate 3 random bytes
	putUint24(b[9:12], atomic.AddUint32(&objectIDCounter, 1))

	return b
}

func NewID() string {
	// Generate a UUID
	uuid := generateUUID()

	var buf [24]byte
	hex.Encode(buf[:], uuid[:])

	return string(buf[:])
}
