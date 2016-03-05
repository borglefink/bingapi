// Copyright 2012 Microsoft (data structure).
// Copyright 2016 Erlend Johannessen (code).
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package bingapi

// CompositeResult holds the result from a composite Bing Web API search
type CompositeResult struct {
	MetaData                 MetaData `json:"__metadata"`
	ID                       string
	WebTotal                 string
	WebOffset                string
	ImageTotal               string
	ImageOffset              string
	VideoTotal               string
	VideoOffset              string
	NewsTotal                string
	NewsOffset               string
	SpellingSuggestionsTotal string
	AlteredQuery             string
	AlterationOverrideQuery  string
	Web                      []WebResult
	Image                    []ImageResult
	Video                    []VideoResult
	News                     []NewsResult
	RelatedSearch            []RelatedResult
	SpellingSuggestions      []SpellResult
}

// CompositeResultContainer holds the CompositeResult
type CompositeResultContainer struct {
	Results []CompositeResult `json:"results"`
}

// CompositeResultWrapper
type CompositeResultWrapper struct {
	Data CompositeResultContainer `json:"d"`
}
