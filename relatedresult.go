// Copyright 2012 Microsoft (data structure).
// Copyright 2016 Erlend Johannessen (code).
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package bingapi

// RelatedResult holds releated results from a Bing Web API search
type RelatedResult struct {
	MetaData MetaData `json:"__metadata"`
	ID       string
	Title    string
	BingUrl  string
}

// RelatedResultContainer holds the RelatedResults
type RelatedResultContainer struct {
	Results     []RelatedResult `json:"results"`
	NextPageURL string          `json:"__next"`
}

// RelatedResultWrapper
type RelatedResultWrapper struct {
	Data RelatedResultContainer `json:"d"`
}
