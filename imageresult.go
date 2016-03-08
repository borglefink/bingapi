// Copyright 2012 Microsoft (data structure).
// Copyright 2016 Erlend Johannessen (code).
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package bingapi

// ImageResult receives a single result from the Bing search.
type ImageResult struct {
	MetaData    MetaData `json:"__metadata"`
	ID          string
	Title       string
	MediaURL    string `json:"MediaUrl"`
	SourceURL   string `json:"SourceUrl"`
	DisplayURL  string `json:"DisplayUrl"`
	Width       int32
	Height      int32
	FileSize    int64
	ContentType string
	Thumbnail   Thumbnail
}

// ImageResultContainer receives a slice of ImageResult
// and a pre-made URL for getting more results
type ImageResultContainer struct {
	Results     []ImageResult `json:"results"`
	NextPageURL string        `json:"__next"`
}

// imageResultWrapper is the outer wrapper for the result.
// Not used for more than receiving the resultcontainer.
type imageResultWrapper struct {
	Data ImageResultContainer `json:"d"`
}
