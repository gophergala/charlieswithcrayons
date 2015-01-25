// Copyright 2015 Gavin Bong. All rights reserved.
// Use of this source code is governed by the Hugsware
// license that can be found in the LICENSE file.

// Package charlieswithcrayons provides access to random.org's JSON RPC api
package charlieswithcrayons

const Version = "1.0.0"
const StdUrl = "https://api.random.org/json-rpc/1/invoke"

type Arguments struct {
	ApiKey string
	Url    string
}

type RandomOrgClient struct {
	apiKey string
	url    string
}

func New(args Arguments) *RandomOrgClient {
	client := new(RandomOrgClient)
	client.apiKey = args.ApiKey
	client.url = args.Url
	return client
}
