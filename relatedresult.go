// Copyright 2012 Microsoft (data structure).
// Copyright 2016 Erlend Johannessen (code).
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package bingapi

// RelatedResult receives a single result from the Bing search.
type RelatedResult struct {
	MetaData MetaData `json:"__metadata"`
	ID       string
	Title    string
	BingURL  string `json:"BingUrl"`
}

// RelatedResultContainer receives a slice of RelatedResult
// and a pre-made URL for getting more results
type RelatedResultContainer struct {
	Results     []RelatedResult `json:"results"`
	NextPageURL string          `json:"__next"`
}

// RelatedResultWrapper is the outer wrapper for the result.
// Not used for more than receiving the resultcontainer.
type RelatedResultWrapper struct {
	Data RelatedResultContainer `json:"d"`
}
