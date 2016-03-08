// Copyright 2012 Microsoft (data structure).
// Copyright 2016 Erlend Johannessen (code).
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package bingapi

// NewsResult receives a single result from the Bing search.
type NewsResult struct {
	MetaData    MetaData `json:"__metadata"`
	ID          string
	Title       string
	URL         string `json:"Url"`
	Source      string
	Description string
	Date        string
}

// NewsResultContainer receives a slice of NewsResult
// and a pre-made URL for getting more results
type NewsResultContainer struct {
	Results     []NewsResult `json:"results"`
	NextPageURL string       `json:"__next"`
}

// newsResultWrapper is the outer wrapper for the result.
// Not used for more than receiving the resultcontainer.
type newsResultWrapper struct {
	Data NewsResultContainer `json:"d"`
}
