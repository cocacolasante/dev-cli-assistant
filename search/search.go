package searchQueries

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const GOOGLE_URL = "https://www.googleapis.com/customsearch/v1?key=%s&q=%s}&start=%s"

type SearchQuery struct {
	SearchTerm string `json:"search_terms"`
	TargetSite string `json:"target_site"`
}

type ResponsePayload struct {
	Title string `json:"title"`
	Link  string `json:"link"`
}

func NewQuery(search string, args ...string) *SearchQuery {
	targetSite := args[0]
	return &SearchQuery{
		SearchTerm: search,
		TargetSite: targetSite,
	}
}

func (s *SearchQuery) NewURL() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}


	API := os.Getenv("SEARCH_API_KEY")
	QUERY_URL := fmt.Sprintf("https://www.googleapis.com/customsearch/v1?key=%s&cx=f29208640b52548c8&q=%s&start=1", API, s.SearchTerm)


	
	if s.TargetSite != ""{
		QUERY_URL += fmt.Sprintf("&siteSearch=%s&siteSearchFilter=i", s.TargetSite)

	}

	
	return QUERY_URL

}


type SearchResult struct {
	Kind              string `json:"kind"`
	URL               struct {
		Type     string `json:"type"`
		Template string `json:"template"`
	} `json:"url"`
	Queries struct {
		Request []struct {
			Title          string `json:"title"`
			TotalResults   string `json:"totalResults"`
			SearchTerms    string `json:"searchTerms"`
			Count          int    `json:"count"`
			StartIndex     int    `json:"startIndex"`
			InputEncoding  string `json:"inputEncoding"`
			OutputEncoding string `json:"outputEncoding"`
			Safe           string `json:"safe"`
			Cx             string `json:"cx"`
		} `json:"request"`
		NextPage []struct {
			Title          string `json:"title"`
			TotalResults   string `json:"totalResults"`
			SearchTerms    string `json:"searchTerms"`
			Count          int    `json:"count"`
			StartIndex     int    `json:"startIndex"`
			InputEncoding  string `json:"inputEncoding"`
			OutputEncoding string `json:"outputEncoding"`
			Safe           string `json:"safe"`
			Cx             string `json:"cx"`
		} `json:"nextPage"`
	} `json:"queries"`
	Context struct {
		Title  string `json:"title"`
		Facets [][]struct {
			Anchor         string `json:"anchor"`
			Label          string `json:"label"`
			LabelWithOp    string `json:"label_with_op"`
		} `json:"facets"`
	} `json:"context"`
	SearchInformation struct {
		SearchTime           float64 `json:"searchTime"`
		FormattedSearchTime  string  `json:"formattedSearchTime"`
		TotalResults         string  `json:"totalResults"`
		FormattedTotalResults string  `json:"formattedTotalResults"`
	} `json:"searchInformation"`
	Items []struct {
		Kind           string `json:"kind"`
		Title          string `json:"title"`
		HtmlTitle      string `json:"htmlTitle"`
		Link           string `json:"link"`
		DisplayLink    string `json:"displayLink"`
		Snippet        string `json:"snippet"`
		HtmlSnippet    string `json:"htmlSnippet"`
		CacheId        string `json:"cacheId"`
		FormattedUrl   string `json:"formattedUrl"`
		HtmlFormattedUrl string `json:"htmlFormattedUrl"`
	} `json:"items"`
}
