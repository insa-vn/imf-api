package googlesearch

import (
	"google.golang.org/api/customsearch/v1"
	"google.golang.org/api/googleapi/transport"
	"net/http"
)

const (
	ENDPOINT = "https://www.googleapis.com/customsearch/v1?"
)

type Searcher struct {
	engine			string
	isImageSearch	bool
	cseService  	*customsearch.CseService
}

func New(apiKey string, engine string, isImageSearch bool) (*Searcher, error) {
	searcher := &Searcher{}
	searcher.engine = engine
	searcher.isImageSearch = isImageSearch

	client := &http.Client{}
	client.Transport = &transport.APIKey { Key: apiKey }
	
	service, err := customsearch.New(client)
	if err != nil {
		return nil, err
	}

	service.BasePath = ENDPOINT
	searcher.cseService = customsearch.NewCseService(service)

	return searcher, nil
}

func (s Searcher) Search(query string, nbResults, startIdx int64) (*customsearch.Search, error) {
	cseListCall := s.cseService.List(query)
	cseListCall.Cx(s.engine)
	if s.isImageSearch {
		cseListCall.SearchType("image")
	}
	cseListCall.Num(nbResults)
	cseListCall.Start(startIdx)
	return cseListCall.Do()
}