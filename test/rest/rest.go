package test

import (
	"net/http"
	"testing"
)

func TestLibraryWithDummyServer(t *testing.T) {
	resp, err := http.Get("http://dummy-server:8080/metrics")
	if err != nil {
		t.Fatalf("Failed to reach the dummy server: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 OK, got %d", resp.StatusCode)
	}
}
