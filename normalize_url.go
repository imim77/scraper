package main

import (
	"errors"
	"net/url"
	"strings"
)

func normalizeURL(rawurl string) (string, error) {
	parsedURL, err := url.Parse(rawurl)
	if err != nil {
		return "", errors.New("error while parsing the url")
	}
	fullPath := parsedURL.Host + parsedURL.Path
	fullPath = strings.TrimSuffix(strings.ToLower(fullPath), "/")
	return fullPath, nil
}
