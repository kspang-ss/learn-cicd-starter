package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKeyValidHeader(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey my-secret-key")
	got, err := GetAPIKey(headers)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := "my-secret-key"
	if got != want {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

func TestGetAPIKeyNoHeader(t *testing.T) {
	headers := http.Header{}
	_, err := GetAPIKey(headers)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if err != ErrNoAuthHeaderIncluded {
		t.Fatalf("expected: %v, got: %v", ErrNoAuthHeaderIncluded, err)
	}
}

func TestGetAPIKeyMalformedHeader(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer my-secret-key")
	_, err := GetAPIKey(headers)
	if err == nil {
		t.Fatal("expected error for malformed header, got nil")
	}
}

