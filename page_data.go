package main

import (
	"fmt"
	"net/url"
)

type PageData struct {
	H1             string
	URL            string
	FirstParagraph string
	OutgoingLinks  []string
	ImageURLs      []string
}

func extractPageData(html, pageURL string) PageData {
	url, err := url.Parse(pageURL)
	if err != nil {
		fmt.Println("there is an error wihle parsing the url")
		return PageData{}
	}
	h1 := getH1FromHTML(html)
	p := getFirstParagraphFromHTML(html)
	urls, err := getURLsFromHTML(html, url)
	if err != nil {
		fmt.Println("error wihle getting the url's from the html")
		return PageData{}
	}
	srcurls, err := getImagesFromHTML(html, url)
	if err != nil {
		fmt.Println("error wihle getting the image url's from the html")
		return PageData{}
	}
	return PageData{
		H1:             h1,
		URL:            pageURL,
		FirstParagraph: p,
		OutgoingLinks:  urls,
		ImageURLs:      srcurls,
	}
}
