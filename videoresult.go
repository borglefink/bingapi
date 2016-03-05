// Copyright 2012 Microsoft (data structure).
// Copyright 2016 Erlend Johannessen (code).
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package bingapi

// VideoResult receives a single result from the Bing search.
type VideoResult struct {
	MetaData   MetaData `json:"__metadata"`
	ID         string
	Title      string
	MediaUrl   string
	DisplayUrl string
	RunTime    int32
	Thumbnail  Thumbnail
}

// VideoResultContainer receives a slice of VideoResult
// and a pre-made URL for getting more results
type VideoResultContainer struct {
	Results     []VideoResult `json:"results"`
	NextPageURL string        `json:"__next"`
}

// VideoResultWrapper is the outer wrapper for the result.
// Not used for more than receiving the resultcontainer.
type VideoResultWrapper struct {
	Data VideoResultContainer `json:"d"`
}
