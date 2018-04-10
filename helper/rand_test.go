package helper

import "testing"

func TestRandStr(t *testing.T) {
	r := RandStr(10)
	if len(r) != 10 {
		t.Error("bad rand string")
	}
}
