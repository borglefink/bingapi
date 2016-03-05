// Copyright 2012 Microsoft (data structure).
// Copyright 2016 Erlend Johannessen (code).
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package bingapi

// ImageResult holds the result from a Image type Bing Web API search
type ImageResult struct {
	MetaData    MetaData `json:"__metadata"`
	ID          string
	Title       string
	MediaUrl    string
	SourceUrl   string
	DisplayUrl  string
	Width       int32
	Height      int32
	FileSize    int64
	ContentType string
	Thumbnail   Thumbnail
}

// ImageResultContainer holds the ImageResults
type ImageResultContainer struct {
	Results     []ImageResult `json:"results"`
	NextPageURL string        `json:"__next"`
}

// ImageResultWrapper
type ImageResultWrapper struct {
	Data ImageResultContainer `json:"d"`
}
