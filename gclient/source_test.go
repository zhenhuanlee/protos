package gclient

import (
	"os"
	"strings"
	"testing"
)

func Test_NewSourceClient(t *testing.T) {
	os.Setenv("SOURCEURL", "localhost:3345")
	r, err := CheckIn("x", "x", "x")
	t.Log(err)
	if err == nil {
		t.Errorf("err should not be nil")
		if r.Code != "updated" {
			t.Errorf("code should be updated but %v", r)
		}
	} else if !strings.Contains(err.Error(), "connection refused") {
		t.Errorf("err should contain 'connection refused' err: %v", err)
	}
}
