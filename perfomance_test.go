package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func startTestServer() *httptest.Server {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	return httptest.NewServer(handler)
}

func TestServerStartDuration(t *testing.T) {
	startTime := time.Now()

	server := startTestServer()
	defer server.Close()

	duration := time.Since(startTime)

	// to adjust the threshold
	maxStartDuration := 2 * time.Second

	if duration > maxStartDuration {
		t.Errorf("Server took too long to start. Expected start time less than %s, but got %s", maxStartDuration, duration)
	}
}

//go test -run TestServerStartDuration > test_results.txt (for result logging)
