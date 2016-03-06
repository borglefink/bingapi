// Copyright 2016 Erlend Johannessen.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package bingapi

// Private constants for bingapi
const (
	bingAPIRootURI    = "https://api.datamarket.azure.com/Bing/Search/v1/" // Bing Search API URI
	bingAPIUserAgent  = "Bing Search Client for Go"                        // Bing UserAgent
	bingAPIMethodGet  = "GET"                                              // Http method GET
	bingAPIMethodPost = "POST"                                             // Http method POST
	bingAPIReqTimeout = 8000                                               // Request Timeout (in milliseconds)
	defaultMarket     = "en-US"                                            // Default Market
	defaultSources    = "web"                                              // Default Sources for Composite search
	defaultTop        = 50                                                 // Default number of results (limited to 50 by API)
)

// Constants for searchtypes
const (
	SearchTypeComposite = "Composite" // Composite type Bing Search API search
	SearchTypeWeb       = "Web"       // Web type Bing Search API search
	SearchTypeNews      = "News"      // News type Bing Search API search
	SearchTypeImage     = "Image"     // Image type Bing Search API search
	SearchTypeVideo     = "Video"     // Video type Bing Search API search
	SearchTypeRelated   = "Related"   // Related type Bing Search API search
)

// MetaData holds meta data for each result.
// This is really just the URI for the specific result
// Metadata is describing, and the type of result.
// See e.g. the WebResult struct.
type MetaData struct {
	URI        string `json:"uri"`
	ResultType string `json:"type"`
}

// Thumbnail holds thumbnail data for an ImageResult or a VideoResult
// from a Bing Search API search.
type Thumbnail struct {
	MediaURL    string `json:"MediaUrl"`
	ContentType string
	Width       int32
	Height      int32
	FileSize    int32
}
