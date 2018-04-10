package helper

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

func Md5(s string) string {
	ctx := md5.New()
	ctx.Write([]byte(s))
	return hex.EncodeToString(ctx.Sum(nil))
}

func StringToInt64(s string) int64 {
	r, e := strconv.ParseInt(s, 10, 64)
	if e != nil {
		return 0
	}
	return r
}

func StringToFloat64(s string) float64 {
	r, e := strconv.ParseFloat(s, 64)
	if e != nil {
		return 0
	}
	return r
}

func StringSliceToStringMap(input []string) map[string]string {
	r := map[string]string{}
	for _, s := range input {
		r[s] = s
	}
	return r
}

func StringSliceUnique(input []string) []string {
	r := StringSliceToStringMap(input)
	nr := make([]string, 0, len(r))
	for key, _ := range r {
		nr = append(nr, key)
	}
	return nr
}
