// Copyright 2012 Microsoft (data structure).
// Copyright 2016 Erlend Johannessen (code).
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package bingapi

// VideoResult holds the result from a Video type Bing Web API search
type VideoResult struct {
	MetaData   MetaData `json:"__metadata"`
	ID         string
	Title      string
	MediaUrl   string
	DisplayUrl string
	RunTime    int32
	Thumbnail  Thumbnail
}

// VideoResultContainer holds the VideoResults
type VideoResultContainer struct {
	Results     []VideoResult `json:"results"`
	NextPageURL string        `json:"__next"`
}

// VideoResultWrapper
type VideoResultWrapper struct {
	Data VideoResultContainer `json:"d"`
}
