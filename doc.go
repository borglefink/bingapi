// Copyright 2016 Erlend Johannessen.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

/*
Package bingapi is an implementation of the Bing Search API.
Read more about the Bing Search API at the Azure Data Marketplace
https://datamarket.azure.com/dataset/bing/search.

To be able to use the Bing Search API an authorization key is needed.
Get it from the Azure Data Marketplace Account Keys https://datamarket.azure.com/account/keys.

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

The datastructures for returning the result are described separately,
but the gist of it is that Bing returns a json structure similar to
the json below, and this needs to be put into Go structs.

  {
    "d": {
      "results": [{
        "__metadata": {
            "uri": "https://api.datamarket.azure.com/Data.ashx/Bing/Search/v1/Web?Query=\u0027xbox\u0027&$skip=0&$top=0",
            "type": "WebResult"
        },
        "ID": "03cfa295-fe87-4fd4-b919-0f85841a67d1",
        "Title": "Xbox | Official Site",
        "Description": "Experience the new generation of games and entertainment with Xbox. Play Xbox games and stream video on all your devices.",
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
