// Copyright 2016 Erlend Johannessen.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package bingapi

import (
	"encoding/json"
	"net/http"
	"time"
)

// Client will contain static parameters for the API call
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

// Composite - executing Composite type Bing Web API search
func (client Client) Composite(parameters Parameters) (CompositeResultContainer, error) {
	var result CompositeResultWrapper

	if err := client.search(parameters.GetURI(SearchTypeComposite), &result); err != nil {
		return CompositeResultContainer{}, err
	}

	return result.Data, nil
}

// Web - executing Web type Bing Web API search
func (client Client) Web(parameters Parameters) (WebResultContainer, error) {
	var result WebResultWrapper

	if err := client.search(parameters.GetURI(SearchTypeWeb), &result); err != nil {
		return WebResultContainer{}, err
	}

	return result.Data, nil
}

// News - executing News type Bing web API search
func (client Client) News(parameters Parameters) (NewsResultContainer, error) {
	var result NewsResultWrapper

	if err := client.search(parameters.GetURI(SearchTypeNews), &result); err != nil {
		return NewsResultContainer{}, err
	}

	return result.Data, nil
}

// Image - executing Image type Bing Image API search
func (client Client) Image(parameters Parameters) (ImageResultContainer, error) {
	var result ImageResultWrapper

	if err := client.search(parameters.GetURI(SearchTypeImage), &result); err != nil {
		return ImageResultContainer{}, err
	}

	return result.Data, nil
}

// Video - executing Video type Bing web API search
func (client Client) Video(parameters Parameters) (VideoResultContainer, error) {
	var result VideoResultWrapper

	if err := client.search(parameters.GetURI(SearchTypeVideo), &result); err != nil {
		return VideoResultContainer{}, err
	}

	return result.Data, nil
}

// Related - executing Related type Bing Related API search
func (client Client) Related(parameters Parameters) (RelatedResultContainer, error) {
	var result RelatedResultWrapper

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
