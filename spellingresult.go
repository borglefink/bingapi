// Copyright 2012 Microsoft (data structure).
// Copyright 2016 Erlend Johannessen (code).
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package bingapi

// SpellResult holds spelling suggestions from a Bing Web API search
type SpellResult struct {
	MetaData MetaData `json:"__metadata"`
	ID       string
	Value    string
}

// SpellResultContainer holds the SpellResults
type SpellResultContainer struct {
	Results     []SpellResult `json:"results"`
	NextPageURL string        `json:"__next"`
}

// SpellResultWrapper
type SpellResultWrapper struct {
	Data SpellResultContainer `json:"d"`
}
