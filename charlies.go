// Copyright 2015 Gavin Bong. All rights reserved.
// Use of this source code is governed by the Hugsware
// license that can be found in the LICENSE file.

// Package charlieswithcrayons provides access
// to random.org's JSON RPC api.
package charlieswithcrayons

import (
	"encoding/json"
	"log"
	"net/http"
    "time"
    "reflect"
    "bytes"
)

const Version = "1.0.0"
const StdUrl = "https://api.random.org/json-rpc/1/invoke"
const methodBlob = "generateBlobs"
const methodBlobSigned = "generateSignedBlobs"

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

type BlobParams struct {
	ApiKey string `json:"apiKey"`
	N      int    `json:"n"`
	Size   int    `json:"size"`
}

type BlobArguments struct {
	JsonRpc    string `json:"jsonrpc"`
	Method     string `json:"method"`
	BlobParams `json:"params"`
	Id         int `json:"id"`
}

func (c *RandomOrgClient) newBlobArguments(method string, bitSize int, txId int) *BlobArguments {
	retval := new(BlobArguments)
	retval.JsonRpc = "2.0"
	retval.Method = method
	retval.Id = txId
	retval.BlobParams.ApiKey = c.apiKey
	retval.BlobParams.N = 1
	retval.BlobParams.Size = bitSize
	return retval
}

// You must have a valid Api key to invoke this method.
// Argument `txId` can be used to differentiate calls.
func (c *RandomOrgClient) GetRandomBits(bitSize int, count int, txId int) error {
	log.Println("invoke generateBlobs")
	blobArguments := c.newBlobArguments(methodBlob, bitSize, txId)
	client := http.Client{
		Timeout: time.Duration(2 * time.Second),
	}

	postBody, err := json.Marshal(blobArguments)
    if err != nil {
		panic(err)
	}
    log.Println(reflect.TypeOf(postBody))

	req, err := http.NewRequest("POST", c.url, bytes.NewBuffer(postBody))
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
    log.Println(resp)
	return err
}

func (c *RandomOrgClient) GetSignedRandom(bitSize int, count int) error {
	log.Println("invoke generateSignedBlobs")
	return nil
}
