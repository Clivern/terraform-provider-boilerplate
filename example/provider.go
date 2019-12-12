// Copyright 2019 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package example

import (
	"log"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() *schema.Provider {
    return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("EXAMPLE_API_KEY", nil),
			},
			"api_url": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "api.example.com",
			},
		},
        ResourcesMap: map[string]*schema.Resource{
        	"example_server": resourceExampleServer(),
        },
        ConfigureFunc: providerConfigure,
    }
}

func providerConfigure(data *schema.ResourceData) (interface{}, error) {
	log.Println("[INFO] Initializing client")

	config := Config{
		ApiKey: data.Get("api_key").(string),
		ApiUrl: data.Get("api_url").(string),
	}

	return config.Client()
}
