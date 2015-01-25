// Copyright 2015 Gavin Bong. All rights reserved.
// Use of this source code is governed by the Hugsware 
// license that can be found in the LICENSE file.

// Package charlieswithcrayons provides access to random.org's JSON RPC api
package charlieswithcrayons

const Hello = "Hello"

type RandomOrgClient struct {
  apiKey string
}

func New() *RandomOrgClient {
  client := new(RandomOrgClient)
  return client
}
