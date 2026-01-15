package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey_ValidHeader(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey test-api-key-123")

	apiKey, err := GetAPIKey(headers)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if apiKey != "test-api-key-123" {
		t.Errorf("Expected API key 'test-api-key-123', got: %s", apiKey)
	}
}

func TestGetAPIKey_MissingHeader(t *testing.T) {
	headers := http.Header{}

	apiKey, err := GetAPIKey(headers)
	if err == nil {
		t.Fatal("Expected error for missing authorization header, got nil")
	}
	if err != ErrNoAuthHeaderIncluded {
		t.Errorf("Expected ErrNoAuthHeaderIncluded, got: %v", err)
	}
	if apiKey != "" {
		t.Errorf("Expected empty API key, got: %s", apiKey)
	}
}

func TestGetAPIKey_MalformedHeader_WrongPrefix(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer token-123")

	apiKey, err := GetAPIKey(headers)
	if err == nil {
		t.Fatal("Expected error for malformed header, got nil")
	}
	if err.Error() != "malformed authorization header" {
		t.Errorf("Expected 'malformed authorization header' error, got: %v", err)
	}
	if apiKey != "" {
		t.Errorf("Expected empty API key, got: %s", apiKey)
	}
}

func TestGetAPIKey_MalformedHeader_NoSpace(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey")

	apiKey, err := GetAPIKey(headers)
	if err == nil {
		t.Fatal("Expected error for malformed header, got nil")
	}
	if err.Error() != "malformed authorization header" {
		t.Errorf("Expected 'malformed authorization header' error, got: %v", err)
	}
	if apiKey != "" {
		t.Errorf("Expected empty API key, got: %s", apiKey)
	}
}

