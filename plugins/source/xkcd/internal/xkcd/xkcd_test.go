package xkcd

import (
	"net/http"
	"testing"
	"time"
)

func TestNewClientDefaultHTTPClientHasTimeout(t *testing.T) {
	c, err := NewClient()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if c.client == http.DefaultClient {
		t.Fatal("expected a dedicated HTTP client, got http.DefaultClient")
	}
	if c.client.Timeout != defaultTimeout {
		t.Fatalf("got timeout %v, want %v", c.client.Timeout, defaultTimeout)
	}
}

func TestNewClientWithHTTPClientKeepsCustomClient(t *testing.T) {
	custom := &http.Client{Timeout: time.Second}
	c, err := NewClient(WithHTTPClient(custom))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if c.client != custom {
		t.Fatal("expected the custom HTTP client to be used")
	}
}
