package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	req, err := http.NewRequest("GET", "http://example.com/path", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "ApiKey asdf")
	got, err := GetAPIKey(req.Header)
	if err != nil || got != "asdf" {
		t.Fatal(err)
	}
}
