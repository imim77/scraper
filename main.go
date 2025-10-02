package main

import (
	"fmt"
	"os"
)

func main() {

	progArgs := os.Args[1:]
	if len(progArgs) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(progArgs) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	BASE_URL := progArgs[0]

	pages := make(map[string]int)

	fmt.Printf("starting crawl of: %s", BASE_URL)

	crawlPage(BASE_URL, BASE_URL, pages)
	fmt.Println("\n=== Crawl Results ===")

	for normalizedURL, count := range pages {
		fmt.Printf("%s: %d\n", normalizedURL, count)
	}

}
