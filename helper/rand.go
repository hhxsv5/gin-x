package helper

import (
	"time"
	"math/rand"
)

var (
	Chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
)

func RandStr(l int) string {
	le := len(Chars)
	data := make([]byte, l, l)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < l; i++ {
		data[i] = byte(Chars[rand.Intn(le)])
	}
	return string(data)
}
