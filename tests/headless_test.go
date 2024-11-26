package tests

import (
	"testing"

	"github.com/ayusheek/reflex/clients"
)

func TestSendHeadlessRequest(t *testing.T) {
	url := "http://127.0.0.1:8000/testsite"
	var showBrowser bool = true
	waitTime := 2
	body, err := clients.SendHeadlessRequest(url, waitTime, showBrowser)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	t.Logf("Response: %s", body)
}
