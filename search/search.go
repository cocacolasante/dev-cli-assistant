package searchQueries

const GOOGLE_URL="https://www.googleapis.com/customsearch/v1?"

type SearchQuery struct {
	SearchTerm string `json:"search_terms"`
	StartIndex int `json:"start_index"`
	
}


func NewQuery(search string, start int) *SearchQuery{
	return &SearchQuery{
		SearchTerm: search,
		StartIndex: start,
	}
}

