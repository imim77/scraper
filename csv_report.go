package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strings"
)

func writeCSVReport(pages map[string]PageData, filename string) error {
	if len(pages) == 0 {
		fmt.Println("No data to write to CSV")
		return nil
	}
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error while creating a file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{"page_url", "h1", "first_paragraph", "outgoing_link_urls", "image_urls"}
	err = writer.Write(headers)
	if err != nil {
		return fmt.Errorf("error while writing headers: %v", err)
	}
	keys := make([]string, 0, len(pages))
	for k := range pages {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Rows
	for _, normalizedURL := range keys {
		p := pages[normalizedURL]
		outgoing := strings.Join(p.OutgoingLinks, ";")
		images := strings.Join(p.ImageURLs, ";")
		row := []string{
			p.URL,
			p.H1,
			p.FirstParagraph,
			outgoing,
			images,
		}
		if err := writer.Write(row); err != nil {
			return fmt.Errorf("write row for %s: %w", p.URL, err)
		}
	}

	fmt.Printf("Report written to %s\n", filename)
	return nil
}
