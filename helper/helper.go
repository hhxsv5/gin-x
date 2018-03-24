package helper

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
)

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
