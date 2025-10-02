package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

/*type config struct {
	pages              map[string]PageData
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}
*/

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

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	parsedBase, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("Error parsing the baseURL to the URL struct%v\n", err)
		return
	}

	parsedCurrent, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error parsing the baseURL to the URL struct%v\n", err)
		return
	}

	if parsedBase.Hostname() != parsedCurrent.Hostname() {
		return
	}
	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error normalizing URL %s: %v\n", rawCurrentURL, err)
		return
	}

	if _, ok := pages[normalizedURL]; ok {
		pages[normalizedURL]++
		return
	}
	pages[normalizedURL] = 1

	fmt.Printf("Crawling: %s\n", rawCurrentURL)
	htmlfromurl, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error getting the HTML from URL: %s: %v\n", rawCurrentURL, err)
	}

	urls, err := getURLsFromHTML(htmlfromurl, parsedCurrent)
	if err != nil {
		fmt.Printf("Error parsing URLs from %s: %v\n", rawCurrentURL, err)
		return
	}

	for _, url := range urls {
		crawlPage(rawBaseURL, url, pages)
	}

}
