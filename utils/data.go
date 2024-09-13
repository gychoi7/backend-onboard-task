package utils

import (
	"crypto/rand"
	"encoding/hex"
	"golang.org/x/crypto/scrypt"
)

func MakePasswordHash(password string) (hashedPassword string, salt string, err error) {
	saltBytes := make([]byte, 32)
	_, err = rand.Read(saltBytes)
	if err != nil {
		return "", "", err
	}
	salt = hex.EncodeToString(saltBytes)
	hashedBytes, err := scrypt.Key([]byte(password), saltBytes, 16384, 8, 1, 32)
	if err != nil {
		return "", "", err
	}
	hashedPassword = hex.EncodeToString(hashedBytes)
	return hashedPassword, salt, nil
}

func CheckPasswordHash(password string, hashedPassword, salt string) bool {
	saltBytes, err := hex.DecodeString(salt)
	if err != nil {
		return false
	}

	hashedBytes, err := scrypt.Key([]byte(password), saltBytes, 16384, 8, 1, 32)
	if err != nil {
		return false
	}

	return hex.EncodeToString(hashedBytes) == hashedPassword
}
