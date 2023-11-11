package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	// "strings"

	"net/http"
	"net/url"

	searchQueries "github.com/cocacolasante/googlecli/search"
)



func main() {
	var phrase string
	flag.StringVar(&phrase, "search", "google.com", "the search string")
	targetSite := flag.String("site", "", "Search results for a specific site")
	sort := flag.String("sort", "", "Sort result by date")

	flag.Parse()

	
	search := url.QueryEscape(phrase)
	
	fmt.Printf("Search string is: %s\n", string(search))



	queryStruct := searchQueries.NewQuery(search, *targetSite, *sort)
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