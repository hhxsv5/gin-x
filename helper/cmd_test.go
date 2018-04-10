package helper

import (
	"testing"
)

func TestExecShell(t *testing.T) {
	r, err := ExecShell("ls")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(r)
	}
}
