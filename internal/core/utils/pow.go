package utils

import (
	"crypto/rand"

	"golang.org/x/crypto/argon2"
)

const (
	OutputLength uint32 = 32
	SaltLength   uint32 = 16
)

var (
	Memory      uint32 = 512 * 1024
	Iteration   uint32 = 3
	Parallelism uint8  = 2
)

func GenerateProofOfWork(input string, salt []byte, iteration uint32, memory uint32) []byte {
	hash := argon2.IDKey([]byte(input), salt, iteration, memory, Parallelism, OutputLength)
	return hash
}

func GenerateSalt(length uint32) ([]byte, error) {
	salt := make([]byte, length)

	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}

	return salt, nil
}
