package main

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strings"
)

var numbersAndLetters = []rune("abcdefghijklmnopqrstuvwxyz1234567890")
var letters = []rune("abcdefghijklmnopqrstuvwxyz")
var numbers = []rune("1234567890")

type Answer struct {
	Pin  string
	Salt string
	Hash string
}

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

// JSON response with pin, salt and sha-1 hash
func pinSaltHash(w http.ResponseWriter, r *http.Request) {
	var n int = 8
	pin := string(generator(n-2, numbers))
	salt := string(generator(10, letters))
	h := shaHash(pin, salt)

	// Response
	resp := Answer{pin, salt, strings.ToUpper(h)}

	// Translate to JSON
	js, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return JSON response
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	http.HandleFunc("/api", pinSaltHash)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
