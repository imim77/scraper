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
	fmt.Printf("starting crawl of: %s", BASE_URL)
	websiteHTML, err := getHTML(BASE_URL)
	if err != nil {
		fmt.Println("Error occured while getting the website's HTML")
		os.Exit(1)
	}
	fmt.Println(websiteHTML)

}
