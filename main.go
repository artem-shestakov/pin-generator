package main

import (
	"fmt"
	"math/rand"
	"time"
)

var numbersAndLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var numbers = []rune("1234567890")

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

func main() {
	var n int = 8
	var strong bool = true
	var pin []rune
	rand.Seed(time.Now().UnixNano())
	if strong {
		pin = generator(n-2, numbers)
		for i := 0; i < 2; i++ {
			pin = insert(pin, rand.Intn(len(pin)), letters[rand.Intn(len(letters))])
		}
		fmt.Println(string(pin))
	} else {
		pin = generator(n, numbers)
		fmt.Println(string(pin))
	}
}
