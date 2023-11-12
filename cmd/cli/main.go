package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"

	"github.com/cocacolasante/googlecli/instructions"
	"github.com/cocacolasante/googlecli/openaiapi"
	searchQueries "github.com/cocacolasante/googlecli/search"
)



func main() {
	//  GOOGLE API FLAGS
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

	// OPEN AI FLAGS
	var usingAi bool
	var content string
	var stream bool
	flag.BoolVar(&usingAi, "ai", false, "Use this call to get ai input")
	flag.StringVar(&content, "content", "", "content to ask ai")
	flag.BoolVar(&stream,"stream", false, "change ai response type to stream answer")

	flag.Parse()

	if helpFlag {
		instructions.PrintHelpMenu()
		return
	}

	if descriptionFlag {
		instructions.PrintDescription()
		return
	}

	if usingAi {
		
		aiReq := openaiapi.NewAiRequest(content)
		if aiReq.Content == ""{
			log.Fatal("empty ai request")
			return 
		}
		if stream {
			aiReq.NewStreamCall()
		}
		
		aiReq.ApiCall()
		return
	}
	
	search := url.QueryEscape(phrase)
	exclude := url.QueryEscape(excludeterm)
	
	fmt.Printf("Search string is: %s\n", string(search))



	queryStruct := searchQueries.NewQuery(search, *targetSite, *sort, exclude)
	searchQuery := queryStruct.NewURL()
	fmt.Println(searchQuery)

	searchQueries.GetGoogleResponse(searchQuery)

	

	
}	

