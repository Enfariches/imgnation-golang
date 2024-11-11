package random

import (
	"math/rand"
	"time"
)

var charsByte = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandStringByte(size int) string { //Усложнить рандомайзер
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	myString := make([]rune, size)
	for i := range myString {
		myString[i] = charsByte[r.Intn(len(charsByte))]
	}

	return string(myString)
}
