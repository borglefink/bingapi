// Copyright 2016 Erlend Johannessen.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package bingapi

import (
	"encoding/json"
	"net/http"
	"time"
)

// Client contains http and auth parameters for the API call.
type Client struct {
	AccKey     string // Bing access key
	Method     string // Http method
	UserAgent  string // Bing call UserAgent
	ReqTimeout int32  // Request timeout
}

// GetClient creates a new set of API call parameters
func GetClient(accessKey string) Client {
	return Client{
		AccKey:     accessKey,
		Method:     bingAPIMethodPost,
		UserAgent:  bingAPIUserAgent,
		ReqTimeout: bingAPIReqTimeout,
	}
}

// Composite - executing Composite type Bing Search API search. Returns a CompositeResultContainer
// containing a slice of CompositeResult, where only the first element is used.
// See documentation for CompositeResultContainer and CompositeResult.
func (client Client) Composite(parameters Parameters) (CompositeResultContainer, error) {
	var result compositeResultWrapper

	if err := client.search(parameters.GetURI(SearchTypeComposite), &result); err != nil {
		return CompositeResultContainer{}, err
	}

	return result.Data, nil
}

// Web - executing Web type Bing Search API search. Returns a WebResultContainer
// containing a slice of WebResult and an URI for next page.
// See documentation for WebResultContainer and WebResult.
func (client Client) Web(parameters Parameters) (WebResultContainer, error) {
	var result webResultWrapper

	if err := client.search(parameters.GetURI(SearchTypeWeb), &result); err != nil {
		return WebResultContainer{}, err
	}

	return result.Data, nil
}

// News - executing News type Bing Search API search. Returns a NewsResultContainer
// containing a slice of NewsResult and an URI for next page.
// See documentation for NewsResultContainer and NewsResult.
func (client Client) News(parameters Parameters) (NewsResultContainer, error) {
	var result newsResultWrapper

	if err := client.search(parameters.GetURI(SearchTypeNews), &result); err != nil {
		return NewsResultContainer{}, err
	}

	return result.Data, nil
}

// Image - executing Image type Bing Image API search. Returns an ImageResultContainer
// containing a slice of ImageResult and an URI for next page.
// See documentation for ImageResultContainer and ImageResult.
func (client Client) Image(parameters Parameters) (ImageResultContainer, error) {
	var result imageResultWrapper

	if err := client.search(parameters.GetURI(SearchTypeImage), &result); err != nil {
		return ImageResultContainer{}, err
	}

	return result.Data, nil
}

// Video - executing Video type Bing Search API search. Returns a VideoResultContainer
// containing a slice of VideoResult and an URI for next page.
// See documentation for VideoResultContainer and VideoResult.
func (client Client) Video(parameters Parameters) (VideoResultContainer, error) {
	var result videoResultWrapper

	if err := client.search(parameters.GetURI(SearchTypeVideo), &result); err != nil {
		return VideoResultContainer{}, err
	}

	return result.Data, nil
}

// Related - executing Related type Bing Related API search. Returns a RelatedResultContainer
// containing a slice of RelatedResult and an URI for next page.
// See documentation for RelatedResultContainer and RelatedResult.
func (client Client) Related(parameters Parameters) (RelatedResultContainer, error) {
	var result relatedResultWrapper

	if err := client.search(parameters.GetURI(SearchTypeRelated), &result); err != nil {
		return RelatedResultContainer{}, err
	}

	return result.Data, nil
}

// search - the main search method, fills target struct with data from bing.
func (client Client) search(uri string, target interface{}) error {
	var httpClient = &http.Client{}
	httpClient.Timeout = time.Duration(client.ReqTimeout) * time.Millisecond

	var req, errReq = http.NewRequest(client.Method, uri, nil)
	if errReq != nil {
		return errReq
	}

	req.Header.Set("User-Agent", client.UserAgent)
	req.SetBasicAuth(client.AccKey, client.AccKey)

	var res, errDo = httpClient.Do(req)
	if errDo != nil {
		return errDo
	}

	defer res.Body.Close()
	return json.NewDecoder(res.Body).Decode(target)
}
