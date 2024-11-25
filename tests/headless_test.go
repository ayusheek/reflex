package tests

import (
	"testing"

	"github.com/ayusheek/reflex/clients"
)

func TestSendHeadlessRequest(t *testing.T) {
	url := "https://ifconfig.me/all"
	var showBrowser bool = true
	body, err := clients.SendHeadlessRequest(url, showBrowser)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	t.Logf("Response: %s", body)
}
