package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawurl string) (string, error) {

	req, err := http.NewRequest("GET", rawurl, nil)
	if err != nil {
		return "", fmt.Errorf("couldn't create a request: %v", err)
	}
	req.Header.Set("User-Agent", "BootCrawler/1.0")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("couldn't send a request: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode > 400 {
		return "", fmt.Errorf("error level code")
	}
	v, ok := resp.Header["Content-Type"]
	if !ok || len(v) == 0 {
		return "", fmt.Errorf("content-type header missing")
	}
	if !strings.HasPrefix(v[0], "text/html") {
		return "", fmt.Errorf("content type is not text/html, got: %s", v[0])
	}
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("could not read response body %v", err)
	}

	return string(resBody), nil
}
