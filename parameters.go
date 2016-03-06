// Copyright 2016 Erlend Johannessen.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package bingapi

import (
	"net/url"
	"strconv"
)

// Parameters holds the query parameters for all types of Bing Web API search
// See Bing Web API documentation for the correct combination for each type of search.
type Parameters struct {
	Query                string  // Required
	Skip                 int     // How may records to skip in the result
	Top                  int     // How may records to return (after Skip)
	Sources              string  // Composite, web+image+video+news+spell
	Adult                string  // Moderate
	ImageFilters         string  // Size:Small+Aspect:Square
	Latitude             float64 // 47.603450
	Longitude            float64 // -122.329696
	Market               string  // en-US
	NewsCategory         string  // rt_Business
	NewsLocationOverride string  // US.WA
	NewsSortBy           string  // Date
	Options              string  // EnableHighlighting
	VideoFilters         string  // Duration:Short+Resolution:High
	VideoSortBy          string  // Date
	WebFileType          string  // XLS
	WebSearchOptions     string  // DisableQueryAlterations
	RootURI              string  // Root URI for Bing search
}

// GetParameters gets a Parameters object with default values set.
func GetParameters(query string) Parameters {
	return Parameters{
		Query:   query,
		Market:  defaultMarket,
		Top:     defaultTop,
		RootURI: bingAPIRootURI,
	}
}

// SetSources sets values for Sources.
// This is only used with composite search.
func (p *Parameters) SetSources(sources string) Parameters {
	p.Sources = defaultIfEmpty(p.Sources, sources)
	return *p
}

// SetPage sets Skip to the correct value,
// based on the current Top (page size) value
// and the given 0-based page number.
func (p *Parameters) SetPage(page int) Parameters {
	p.Skip = page * p.Top
	return *p
}

// ExactSearch sets the option to disable query alterations
// for web type search.
func (p *Parameters) ExactSearch() Parameters {
	p.WebSearchOptions = "DisableQueryAlterations"
	return *p
}

// GetURI returns the URI string based on the parameters set,
// and the given search type.
func (p Parameters) GetURI(searchType string) string {
	var uri = p.RootURI + searchType

	uri += addFormat("json")
	uri += addEscaped("Query", p.Query)
	uri += addPaging("top", p.Top)
	uri += addPaging("skip", p.Skip)
	uri += addParameter("Adult", p.Adult)
	uri += addFloat("Latitude", p.Latitude)
	uri += addFloat("Longitude", p.Longitude)
	uri += addParameter("Market", p.Market)
	uri += addParameter("Options", p.Options)

	switch searchType {
	case SearchTypeComposite:
		p.SetSources(defaultSources)
		uri += addEscaped("Sources", p.Sources)
		uri += addParameter("ImageFilters", p.ImageFilters)
		uri += addParameter("NewsCategory", p.NewsCategory)
		uri += addParameter("NewsLocationOverride", p.NewsLocationOverride)
		uri += addParameter("NewsSortBy", p.NewsSortBy)
		uri += addParameter("VideoFilters", p.VideoFilters)
		uri += addParameter("VideoSortBy", p.VideoSortBy)
		uri += addParameter("WebFileType", p.WebFileType)
		uri += addParameter("WebSearchOptions", p.WebSearchOptions)
	case SearchTypeWeb:
		uri += addParameter("WebFileType", p.WebFileType)
		uri += addParameter("WebSearchOptions", p.WebSearchOptions)
	case SearchTypeNews:
		uri += addParameter("NewsCategory", p.NewsCategory)
		uri += addParameter("NewsLocationOverride", p.NewsLocationOverride)
		uri += addParameter("NewsSortBy", p.NewsSortBy)
	case SearchTypeImage:
		uri += addParameter("ImageFilters", p.ImageFilters)
	case SearchTypeVideo:
		uri += addParameter("VideoFilters", p.VideoFilters)
		uri += addParameter("VideoSortBy", p.VideoSortBy)
	}

	return uri
}

func addFormat(format string) string {
	return "?$format=" + format
}

func addEscaped(key, value string) string {
	return "&" + key + "=" + url.QueryEscape("'"+value+"'")
}

func addOptionalEscaped(key, value string) string {
	if value != "" {
		return "&" + key + "=" + url.QueryEscape("'"+value+"'")
	}
	return ""
}

func addPaging(key string, value int) string {
	return "&$" + key + "=" + strconv.Itoa(value)
}

func addParameter(key, value string) string {
	if value != "" {
		return "&" + addKeyValue(key, value)
	}
	return ""
}

func addKeyValue(key, value string) string {
	return key + "=%27" + value + "%27"
}

func addFloat(key string, value float64) string {
	if value != 0 {
		return "&" + addKeyValue(key, strconv.FormatFloat(value, 'f', 3, 64))
	}
	return ""
}

func defaultIfEmpty(value, defaultValue string) string {
	if value == "" {
		return defaultValue
	}
	return value
}
