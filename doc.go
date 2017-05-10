// Copyright 2016 Erlend Johannessen.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

/*
Package bingapi is an implementation of the Bing Search API v2.

  Bing Search API v2 was closed down 31st of March 2017.

So, what is the Bing Search API? Microsoft explains:

  The Bing Search API enables developers to embed search results in applications or
  websites using XML or JSON. Add search functionality to a website, create unique
  consumer or enterprise apps, or develop new mash-ups.

Registered users can (as of this writing) have 5000 queries/month
for free, after that a fee is applied (until the next month).
Read more about the Bing Search API at the Azure Data Marketplace,
see https://datamarket.azure.com/dataset/bing/search.

To be able to use the Bing Search API an authorization key is needed.
Get it from the Azure Data Marketplace Account Keys,
see https://datamarket.azure.com/account/keys.

Example of usage:

  package main

  import (
  	"fmt"
  	"github.com/borglefink/bingapi"
  )

  func main() {
  	var parameters = bingapi.GetParameters("[my query]")
  	var client = bingapi.GetClient("[MyBingAuthKey]")

  	var container, err = client.Web(parameters)

  	if err != nil {
  		panic(err)
  	}

  	for _, item := range container.Results {
  		fmt.Printf("%s\n", item.Title)
  	}
  }


There are 7 different search formats; Web, News, Image, Video, Related, Spelling and Composite, the latter being a combination of the previous formats, and Spelling being returned with the Composite format.
The parameters to the various searches differ, so be sure to use the correct parameters with each format, as described below.

  Query                string  // All formats
  Adult                string  // All formats
  Latitude             float64 // All formats
  Longitude            float64 // All formats
  Market               string  // All formats
  Options              string  // All formats

  Sources              string  // Composite, web+image+video+news+spell

  ImageFilters         string  // Image, Composite (when Sources=image)

  NewsCategory         string  // News, Composite (when Sources=news)
  NewsLocationOverride string  // News, Composite (when Sources=news)
  NewsSortBy           string  // News, Composite (when Sources=news)

  VideoFilters         string  // Video, Composite (when Sources=video)
  VideoSortBy          string  // Video, Composite (when Sources=video)

  WebFileType          string  // Web, Composite (when Sources=web)
  WebSearchOptions     string  // Web, Composite (when Sources=web)


The data structures for returning the result are described separately,
but the gist of it is that Bing returns a json structure similar to
the json below, and this needs to be put into Go structs.

  {                          // (wrapper)
    "d": {                   // e.g. WebResultContainer
      "results": [{          // e.g. []WebResult
        "__metadata": {
            "uri": "https://api.datamarket.azure.com/Data.ashx/Bing/Search/v1/Web?Query=\u0027xbox\u0027&$skip=0&$top=0",
            "type": "WebResult"
        },
        "ID": "03cfa295-fe87-4fd4-b919-0f85841a67d1",
        "Title": "Xbox | Official Site",
        "Description": "Experience the new generation of games and entertainment with Xbox.",
        "DisplayUrl": "www.xbox.com",
        "Url": "http://www.xbox.com/"
      }, {

        (plus 49 more results)

      }],
      "__next": "https://api.datamarket.azure.com/Data.ashx/Bing/Search/v1/Web?Query=\u0027xbox\u0027&$skip=50&$top=50"
    }
  }

NB: Bing Search API can return results as json and as XML.
Only json is implemented in this version of the package.

*/
package bingapi
