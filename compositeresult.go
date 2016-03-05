// Copyright 2012 Microsoft (data structure).
// Copyright 2016 Erlend Johannessen (code).
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package bingapi

// CompositeResult receives all results from a composite Bing search.
// The totals and offsets below are in the Bing documentation described as int64,
// but in the received json they were strings.
type CompositeResult struct {
	MetaData                 MetaData `json:"__metadata"`
	ID                       string
	WebTotal                 string // originally int64
	WebOffset                string // originally int64
	ImageTotal               string // originally int64
	ImageOffset              string // originally int64
	VideoTotal               string // originally int64
	VideoOffset              string // originally int64
	NewsTotal                string // originally int64
	NewsOffset               string // originally int64
	SpellingSuggestionsTotal string // originally int64
	AlteredQuery             string
	AlterationOverrideQuery  string
	Web                      []WebResult
	Image                    []ImageResult
	Video                    []VideoResult
	News                     []NewsResult
	RelatedSearch            []RelatedResult
	SpellingSuggestions      []SpellResult
}

// CompositeResultContainer receives a slice of CompositeResult.
// Only the first element in the slice is used.
type CompositeResultContainer struct {
	Results []CompositeResult `json:"results"`
}

// CompositeResultWrapper is the outer wrapper for the result.
// Not used for more than receiving the resultcontainer.
type CompositeResultWrapper struct {
	Data CompositeResultContainer `json:"d"`
}
