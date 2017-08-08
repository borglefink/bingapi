## NB: Bing Search API v2 was closed down 31st of March 2017

[![GoDoc](https://godoc.org/github.com/borglefink/bingapi?status.svg)](https://godoc.org/github.com/borglefink/bingapi)
# bingapi #

An implementation in Go of the Bing Search API v2, which was closed down 31st of March 2017. 

Read more about the Bing Search API at the [Azure Data Marketplace](https://datamarket.azure.com/dataset/bing/search).


## Usage

To be able to use the Bing Search API an authorization key is needed. Get it from the [Azure Data Marketplace | Account Keys](https://datamarket.azure.com/account/keys).

```Go
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
```


## Search types and parameters

To understand the Bing Search API in depth, please refer to the Azure web site describing it in [more detail](https://datamarket.azure.com/dataset/bing/search).

There are 7 different search formats; Web, News, Image, Video, Related, Spelling and Composite, the latter being a combination of the previous formats, and Spelling being returned with the Composite format.

The parameters to the various formats differ, so be sure to use the correct parameters with each format. The relevant parameters are described below.

```
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
```


## License

A MIT license is used here - do what you want with this. Nice though if improvements and corrections could trickle back to me somehow. :-)
