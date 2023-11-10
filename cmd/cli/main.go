package main

import (
	"flag"
	"fmt"

	searchQueries "github.com/cocacolasante/googlecli/search"
)



func main() {
	searchString := flag.String("search", "google.com", "the search string")
	startIndex := flag.Int("start", 1, "page to start at")
	

	flag.Parse()

	search := searchString
	startIn := startIndex
	fmt.Printf("Search string is: %s\n", string(*search))
	fmt.Printf("Search start index is: %x\n", *startIn)

	url := searchQueries.GOOGLE_URL

	query := searchQueries.NewQuery(*search, *startIn)
	fmt.Printf("Search Query %s%s", url, query.SearchTerm)
}