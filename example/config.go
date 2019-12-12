// Copyright 2019 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package example

import (
    "log"
    "github.com/clivern/terraform-provider-example/sdk"
)

type ExampleClient struct{
	Client *sdk.Client
}

type Config struct {
    ApiKey string
    ApiUrl string
}

func (c *Config) Client() (*ExampleClient, error) {
    cli := &sdk.Client{
        ApiKey: c.ApiKey,
        ApiUrl: c.ApiUrl,
    }

    cli.Init()

    log.Printf("[INFO] upstream client configured")

    return &ExampleClient{Client: cli}, nil
}