package utils

import (
	"math/rand"
	"time"
)

var randSrc = rand.NewSource(time.Now().UnixNano())

const idBytes = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
const idBits = 6
const idMask = 1<<idBits - 1

// GenerateID will generate a 32 bytes ID for machine
func GenerateID() string {
	return generateID(32)
}

// GenerateShortID will generate a 6 bytes short ID for human being
func GenerateShortID() string {
	return generateID(6)
}

func generateID(length int) string {
	b := make([]byte, length)
	for i, r, remain := length-1, randSrc.Int63(), 10; i >= 0; {
		if remain == 0 {
			r, remain = randSrc.Int63(), 10
		}
		if idx := int(r & idMask); idx < 62 {
			b[i] = idBytes[idx]
			i--
		}
		r >>= idBits
		remain--
	}

	return string(b)
}
