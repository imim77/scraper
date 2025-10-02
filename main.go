package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	progArgs := os.Args[1:]
	if len(progArgs) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(progArgs) > 3 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	BASE_URL := progArgs[0]
	max_conc, err := strconv.Atoi(progArgs[1])
	if err != nil {
		fmt.Printf("Error - conversion from string to int: %v", err)
		return
	}
	max_page, err := strconv.Atoi(progArgs[2])
	if err != nil {
		fmt.Printf("Error - conversion from string to int: %v", err)
		return
	}

	cfg, err := configure(BASE_URL, max_conc, max_page)
	if err != nil {
		fmt.Printf("Error - configure: %v", err)
		return
	}

	fmt.Printf("starting crawl of: %s", BASE_URL)

	cfg.wg.Add(1)
	go cfg.crawlPage(BASE_URL)
	cfg.wg.Wait()
	fmt.Println("\n=== Crawl Results ===")
	writeCSVReport(cfg.pages, "report.csv")
	/*for normalizedURL, count := range cfg.pages {
		fmt.Printf("%s: %+v\n", normalizedURL, count)
	}
	*/

}
