package random

import (
	"math/rand"
	"time"
)

var charsByte = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func RandStringByte(size int) string {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	myString := make([]byte, size)
	for i := range myString {
		myString[i] = charsByte[r.Intn(len(myString))]
	}

	return string(myString)
}
