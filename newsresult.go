// Copyright 2012 Microsoft (data structure).
// Copyright 2016 Erlend Johannessen (code).
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package bingapi

// NewsResult holds the result from a News type Bing Web API search
type NewsResult struct {
	MetaData    MetaData `json:"__metadata"`
	ID          string
	Title       string
	Url         string
	Source      string
	Description string
	Date        string
}

// NewsResultContainer holds the NewsResults
type NewsResultContainer struct {
	Results     []NewsResult `json:"results"`
	NextPageURL string       `json:"__next"`
}

// NewsResultWrapper
type NewsResultWrapper struct {
	Data NewsResultContainer `json:"d"`
}
