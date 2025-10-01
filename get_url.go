package main

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return []string{}, errors.New("Error while creating an new doc")
	}
	urls := []string{}
	doc.Find("a[href]").Each(func(_ int, s *goquery.Selection) {
		linkstr, ok := s.Attr("href")
		if !ok {
			return

		}
		linkstr = strings.TrimSpace(linkstr)
		parsedlinkstr, err := url.Parse(linkstr)
		if err != nil {
			fmt.Println("Error while parsing: ", err)
			return
		}
		urls = append(urls, baseURL.ResolveReference(parsedlinkstr).String())
	})
	return urls, nil
}

func getImagesFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return []string{}, errors.New("Error while creating an new doc")
	}
	srcimages := []string{}
	doc.Find("img[src]").Each(func(_ int, s *goquery.Selection) {
		imgval, ok := s.Attr("src")
		if !ok {
			return
		}
		imgval = strings.TrimSpace(imgval)
		u, err := url.Parse(imgval)
		if err != nil {
			fmt.Println("Error while parsing: ", err)
			return
		}
		srcimages = append(srcimages, baseURL.ResolveReference(u).String())
	})
	return srcimages, nil
}
