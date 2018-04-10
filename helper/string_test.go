package helper

import "testing"

func TestStrToInt64(t *testing.T) {
	i := StringToInt64("3000")
	if i != 3000 {
		t.Error(i)
	}
	t.Log(i)

}

func TestStrToFloat64(t *testing.T) {
	i := StringToFloat64("3000.1234")
	if i != 3000.1234 {
		t.Error(i)
	}
	t.Log(i)
}

func TestMd52(t *testing.T) {
	r := Md5("123")
	if r != "202cb962ac59075b964b07152d234b70" {
		t.Error("bad md5")
	}
}

func TestStringSliceToStringMap(t *testing.T) {
	a := []string{"a", "b", "b", "a"}
	r := StringSliceToStringMap(a)
	if len(r) != 2 {
		t.Error("bad map")
	}
}

func TestStringSliceUnique(t *testing.T) {
	a := []string{"a", "b", "b", "a"}
	r := StringSliceUnique(a)
	if len(r) != 2 {
		t.Error("bad unique")
	}
}
