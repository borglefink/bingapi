// Copyright 2012 Microsoft (data structure).
// Copyright 2016 Erlend Johannessen (code).
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package bingapi

// WebResult holds the result from a Web type Bing Web API search
type WebResult struct {
	MetaData    MetaData `json:"__metadata"`
	ID          string
	Title       string
	Description string
	DisplayUrl  string
	Url         string
}

// WebResultContainer holds the WebResults
type WebResultContainer struct {
	Results     []WebResult `json:"results"`
	NextPageURL string      `json:"__next"`
}

// WebResultWrapper
type WebResultWrapper struct {
	Data WebResultContainer `json:"d"`
}
