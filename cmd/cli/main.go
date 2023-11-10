package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"

	"net/http"

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

	

	queryStruct := searchQueries.NewQuery(*search, *startIn)
	searchQuery := queryStruct.NewURL()
	fmt.Println(searchQuery)

	response, err := http.Get(searchQuery)
	if err != nil {
		fmt.Printf("Error making GET request: %s\n", err)
		return
	}

	defer response.Body.Close()
	

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %s\n", err)
		return
	}
	// fmt.Println(string(body))

	
	var items searchQueries.SearchResult
	err = json.Unmarshal(body, &items)
	if err != nil {
		fmt.Printf("Error umarshalling response body: %s\n", err)
		return
	}

	for _, item := range items.Items {
		fmt.Printf("Name: %s\nLink: %s\n\n", item.Title, item.Link)
	}

	
}	