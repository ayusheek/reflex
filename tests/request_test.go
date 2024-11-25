package tests

import (
	"testing"

	"github.com/ayusheek/reflex/clients"
)

func TestSendHTTPRequest(t *testing.T) {
	timeout := 10
	url := "https://ifconfig.me/all"

	headers := map[string]string{
		"User-Agent": "reflex/1.0",
		"Accept":     "*/*",
	}

	body, err := clients.SendHTTPRequest(url, timeout, headers)

	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	t.Logf("Response: %s", body)
}
