package clients

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// SendHTTPRequest sends an HTTP GET request to the given URL
func SendHTTPRequest(url string, timeout int, headers map[string]string) (string, error) {
	// Initialize an HTTP client with a timeout
	client := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}

	// Create a new HTTP GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("error creating request for %s: %v", url, err)
	}

	// Add headers to the request, if any
	for key, value := range headers {
		req.Header.Add(key, value)
	}

	// Send the HTTP request
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request to %s: %v", url, err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body from %s: %v", url, err)
	}

	return string(body), nil
}
