// Copyright 2012 Microsoft (data structure).
// Copyright 2016 Erlend Johannessen (code).
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package bingapi

// SpellResult receives a spelling suggestion from the Bing search.
// This struct is used with composite search.
type SpellResult struct {
	MetaData MetaData `json:"__metadata"`
	ID       string
	Value    string
}
