// Copyright 2019 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package boilerplate

import (
	"log"

	"github.com/clivern/terraform-provider-boilerplate/sdk"
)

// Client provider client
type Client struct {
	Client *sdk.Client
}

// Config provider config
type Config struct {
	APIKey string
	APIURL string
}

// Client gets an instance of the provider client
func (c *Config) Client() (*Client, error) {
	cli := &sdk.Client{
		APIKey: c.APIKey,
		APIURL: c.APIURL,
	}

	cli.Init()

	log.Printf("[INFO] Upstream Client Configured")

	return &Client{Client: cli}, nil
}
