package helper

import (
	"time"
	"math/rand"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
)

var (
	codes   = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	codeLen = len(codes)
)

func RandStr(l int) string {
	data := make([]byte, l, l)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < l; i++ {
		data[i] = byte(codes[rand.Intn(codeLen)])
	}
	return string(data)
}

func Md5(s string) string {
	ctx := md5.New()
	ctx.Write([]byte(s))
	return hex.EncodeToString(ctx.Sum(nil))
}

func Json2String(d interface{}) string {
	j, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}
	return string(j)
}

func String2Json(j string, v interface{}) {
	err := json.Unmarshal([]byte(j), v)
	if err != nil {
		panic(err)
	}
}
