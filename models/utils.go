package models

import (
	"crypto/sha1"
	"encoding/hex"
	"math/rand"
)

// Function generator array of rune
func generator(length int, alphabet []rune) []rune {
	str := make([]rune, length)
	for i := range str {
		str[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return str
}

// Insert rune into array of rune
func insert(str []rune, index int, value rune) []rune {
	if len(str) == index {
		return append(str, value)
	}
	str = append(str[:index+1], str[index:]...)
	str[index] = value
	return str
}

// Getting SHA-1 hash string
func shaHash(pin, salt string) string {
	sha := sha1.New()
	sha.Write([]byte(pin + salt))
	h := sha.Sum(nil)
	return hex.EncodeToString(h)
}
