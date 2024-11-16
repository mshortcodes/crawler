package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	resp, err := http.Get(rawURL)
	if err != nil {
		return "", fmt.Errorf("couldn't fetch page: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return "", fmt.Errorf("status code error: %s", resp.Status)
	}

	contentType := resp.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		return "", fmt.Errorf("non-HTML response: %s", contentType)
	}

	htmlBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("couldn't read response body: %v", err)
	}

	htmlBody := string(htmlBytes)
	return htmlBody, nil
}
