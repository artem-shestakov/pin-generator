package models

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strings"
)

var numbersAndLetters = []rune("abcdefghijklmnopqrstuvwxyz1234567890")
var letters = []rune("abcdefghijklmnopqrstuvwxyz")
var numbers = []rune("1234567890")

type Answer struct {
	Pin  string `json:"pin"`
	Salt string `json:"salt"`
	Hash string `json:"hash"`
}

func (a *Answer) ToJSON(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	return encoder.Encode(a)
}

func (a *Answer) Generate(pinLen, saltLen int, strong bool, letterCount ...int) {
	if strong && len(letterCount) != 0 {
		pin := generator(pinLen-letterCount[0], numbers)
		for i := 0; i < letterCount[0]; i++ {
			pin = insert(pin, rand.Intn(len(pin)), letters[rand.Intn(len(letters))])
		}
		a.Pin = string(pin)
		a.Salt = string(generator(saltLen, letters))

	} else {
		a.Pin = string(generator(pinLen, numbers))
		a.Salt = string(generator(saltLen, letters))
	}
	a.Hash = strings.ToUpper(shaHash(a.Pin, a.Salt))
}
