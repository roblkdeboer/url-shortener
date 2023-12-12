package utils

import (
	"crypto/rand"
	"log"
	"time"
)

const (
	base62Chars  = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	base62Length = int64(len(base62Chars))
)

func GenerateShortURL() string {
	// Use a combination of timestamp and random number
	timestamp := time.Now().UnixNano()
	randomNumber := generateRandomNumber()
	combinedValue := timestamp + int64(randomNumber)

	// Convert the combined value to Base62
	shortURL := toBase62(combinedValue)

	return shortURL
}

// generateRandomNumber generates a random number to add variability.
func generateRandomNumber() int {
	var bytes [8]byte
	_, err := rand.Read(bytes[:])
	if err != nil {
		log.Fatal("Error generating random number:", err)
	}
	return int(bytes[0])
}

// toBase62 converts a decimal number to Base62.
func toBase62(num int64) string {
	var result string
	for num > 0 {
		remainder := num % base62Length
		result = string(base62Chars[remainder]) + result
		num /= base62Length
	}
	return result
}