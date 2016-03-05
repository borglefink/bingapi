// Copyright 2016 Erlend Johannessen.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package bingapi

// Constants for bingapi
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
	SearchTypeComposite = "Composite" // Composite type Bing Web API search
	SearchTypeWeb       = "Web"       // Web type Bing Web API search
	SearchTypeNews      = "News"      // News type Bing Web API search
	SearchTypeImage     = "Image"     // Image type Bing Web API search
	SearchTypeVideo     = "Video"     // Video type Bing Web API search
	SearchTypeRelated   = "Related"   // Related type Bing Web API search
)

// MetaData holds meta data for each result
type MetaData struct {
	Uri        string `json:"uri"`
	ResultType string `json:"type"`
}

// Thumbnail holds thumbnail data from a Video or Image type Bing Web API search
type Thumbnail struct {
	MediaUrl    string
	ContentType string
	Width       int32
	Height      int32
	FileSize    int32
}
