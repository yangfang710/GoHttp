package util

import (
	"math/rand"
	"time"
	"unsafe"
)

const (
	NumberLetterBytes = "123456789"
	LetterBytes       = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func RandString(n int, source string) string {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, n)
	sourceLen := len(source)

	for i := range b {
		b[i] = source[r.Intn(sourceLen)]
	}

	return *(*string)(unsafe.Pointer(&b))
}
