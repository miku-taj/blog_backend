package utils

import (
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	encodeToString := hex.EncodeToString(hash)

	return encodeToString, nil
}

func CheckPasswordHash(password string, hash string) bool {
	decodeString, err := hex.DecodeString(hash)
	if err != nil {
		return false
	}

	err = bcrypt.CompareHashAndPassword(decodeString, []byte(password))
	if err != nil {
		return false
	}
	return true
}

// GenerateRandomString генерирует случайную строку заданной длиной
func GenerateRandomString(length int) string {
	// Генерация случайного пароля
	b := make([]byte, length)
	for i := range b {
		b[i] = byte(97 + rand.Intn(26)) // a-z
	}
	return string(b)
}
