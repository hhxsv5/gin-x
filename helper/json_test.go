package helper

import (
	"testing"
)

func TestJson2String(t *testing.T) {
	a := struct {
		A string `json:"aaa"`
		B uint   `json:"bbb"`
	}{A: "a1", B: 123}
	r, err := Json2String(a)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(r)
	}
}

func TestString2Json(t *testing.T) {
	s := "{\"aaa\":\"aaa\",\"bbb\":123}"
	a := struct {
		A string `json:"aaa"`
		B uint   `json:"bbb"`
	}{}
	err := String2Json(s, &a)
	if err != nil {
		t.Error(err)
	}
}
