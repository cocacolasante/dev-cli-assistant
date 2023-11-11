package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"

	// "strings"

	"net/http"
	"net/url"

	"github.com/cocacolasante/googlecli/instructions"
	searchQueries "github.com/cocacolasante/googlecli/search"
)



func main() {
	var phrase string
	var excludeterm string
	var helpFlag bool 
	var descriptionFlag bool

	flag.StringVar(&phrase, "search", "google.com", "the search string")
	targetSite := flag.String("site", "", "Search results for a specific site")
	sort := flag.String("sort", "", "Sort result by date")
	flag.StringVar(&excludeterm, "exclude", "", "terms to exclude")
	flag.BoolVar(&helpFlag, "help", false, "print usage instructions")
	flag.BoolVar(&descriptionFlag, "description", false, "prints cli description")

	flag.Parse()

	if helpFlag {
		instructions.PrintHelpMenu()
		return
	}

	if descriptionFlag {
		instructions.PrintDescription()
		return
	}
	
	search := url.QueryEscape(phrase)
	exclude := url.QueryEscape(excludeterm)
	
	fmt.Printf("Search string is: %s\n", string(search))



	queryStruct := searchQueries.NewQuery(search, *targetSite, *sort, exclude)
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