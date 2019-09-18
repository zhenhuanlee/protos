package gclient

import (
	"os"
	"strings"
	"testing"
)

func Test_NewSourceClient(t *testing.T) {
	os.Setenv("SOURCEURL", "localhost:1234")
	_, err := CheckIn("x", "x", "x")
	t.Log(err)
	if err == nil {
		t.Errorf("err should not be nil")
	}
	if !strings.Contains(err.Error(), "connection refused") {
		t.Errorf("err should contain 'connection refused' err: %v", err)
	}
}
