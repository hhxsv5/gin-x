package helper

import (
	"testing"
)

func TestFileIsDir(t *testing.T) {
	f := "./xxx"
	r := FileIsDir(f)
	if r {
		t.Error("bad dir")
	}
}

func TestFileExists(t *testing.T) {
	f := "./xxx"
	r := FileExists(f)
	if r {
		t.Error("bad exist")
	}
}
