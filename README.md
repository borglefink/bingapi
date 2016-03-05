# bingapi #
[![GoDoc](https://godoc.org/github.com/borglefink/bingapi?status.svg)](https://godoc.org/github.com/borglefink/bingapi)

An implementation of the Bing Search API in Go. Read more about the Bing Search API at the [Azure Data Marketplace](https://datamarket.azure.com/dataset/bing/search).


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

There are 7 different search formats; Web, News, Image, Video, Related, Spelling and Composite, the latter being a combination of the previous formats. The parameters to the mentioned formats differ, so when using the parameter class some care is necessary. The relevant parameters are categorised below.

```Go
type Parameters struct {
	Query                string  // All formats
	Sources              string  // Composite, web+image+video+news+spell
	Adult                string  // All formats
	ImageFilters         string  // Image, Composite (when Sources=image)
	Latitude             float64 // All formats
	Longitude            float64 // All formats
	Market               string  // All formats
	NewsCategory         string  // News, Composite (when Sources=news)
	NewsLocationOverride string  // News, Composite (when Sources=news)
	NewsSortBy           string  // News, Composite (when Sources=news)
	Options              string  // All formats
	VideoFilters         string  // Video, Composite (when Sources=video)
	VideoSortBy          string  // Video, Composite (when Sources=video)
	WebFileType          string  // Web, Composite (when Sources=web)
	WebSearchOptions     string  // Web, Composite (when Sources=web)
}
```


## License

A MIT license is used here - do what you want with this. Nice though if improvements and corrections could trickle back to me somehow. :-)
