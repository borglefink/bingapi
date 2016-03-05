// Copyright 2012 Microsoft (data structure).
// Copyright 2016 Erlend Johannessen (code).
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package bingapi

// WebResult receives a single result from the Bing search.
type WebResult struct {
	MetaData    MetaData `json:"__metadata"`
	ID          string
	Title       string
	Description string
	DisplayURL  string   `json:"DisplayUrl"`
	URL         string   `json:"Url"`
}

// WebResultContainer receives a slice of WebResult
// and a pre-made URL for getting more results
type WebResultContainer struct {
	Results     []WebResult `json:"results"`
	NextPageURL string      `json:"__next"`
}

// WebResultWrapper is the outer wrapper for the result.
// Not used for more than receiving the resultcontainer.
type WebResultWrapper struct {
	Data WebResultContainer `json:"d"`
}
