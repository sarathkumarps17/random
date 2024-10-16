package random

import (
	"math/rand"
	"time"
)

func GenerateRandomNumber(min, max int) int {
	salt := rand.NewSource(time.Now().UnixNano())
	r := rand.New(salt)
	return r.Intn(max-min+1) + min
}

func GenerateRandomString(length int) string {
	salt := rand.NewSource(time.Now().UnixNano())
	r := rand.New(salt)
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}
	return string(b)
}
