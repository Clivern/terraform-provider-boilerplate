// Copyright 2019 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package example

import (
	"log"
	"strconv"
	"strings"

	"github.com/clivern/terraform-provider-example/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceExampleServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceExampleServerCreate,
		Read:   resourceExampleServerRead,
		Update: resourceExampleServerUpdate,
		Delete: resourceExampleServerDelete,
		Schema: map[string]*schema.Schema{
			"image": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.NoZeroValues,
			},
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},
			"region": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				StateFunc: func(val interface{}) string {
					return strings.ToLower(val.(string))
				},
				ValidateFunc: validation.NoZeroValues,
			},
			"size": {
				Type:     schema.TypeString,
				Required: true,
				StateFunc: func(val interface{}) string {
					return strings.ToLower(val.(string))
				},
				ValidateFunc: validation.NoZeroValues,
			},
		},
	}
}

func resourceExampleServerCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*ExampleClient).Client

	server, err := client.CreateServer(&sdk.Server{
		Image:  d.Get("image").(string),
		Name:   d.Get("name").(string),
		Size:   d.Get("size").(string),
		Region: d.Get("region").(string),
	})

	log.Printf("[INFO] Creating Server")

	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(server.Id))
	d.Set("name", server.Name)
	d.Set("size", server.Size)
	d.Set("image", server.Image)
	d.Set("region", server.Region)

	return resourceExampleServerRead(d, m)
}

func resourceExampleServerRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*ExampleClient).Client

	id, err := strconv.Atoi(d.Id())

	if err != nil {
		return err
	}

	server, err := client.GetServer(id)

	log.Printf("[INFO] Getting Server")

	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(server.Id))
	d.Set("name", server.Name)
	d.Set("size", server.Size)
	d.Set("image", server.Image)
	d.Set("region", server.Region)

	return nil
}

func resourceExampleServerUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*ExampleClient).Client

	id, err := strconv.Atoi(d.Id())

	if err != nil {
		return err
	}

	server, err := client.CreateServer(&sdk.Server{
		Id:     id,
		Image:  d.Get("image").(string),
		Name:   d.Get("name").(string),
		Size:   d.Get("size").(string),
		Region: d.Get("region").(string),
	})

	log.Printf("[INFO] Creating Server")

	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(server.Id))
	d.Set("name", server.Name)
	d.Set("size", server.Size)
	d.Set("image", server.Image)
	d.Set("region", server.Region)

	return resourceExampleServerRead(d, m)
}

func resourceExampleServerDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*ExampleClient).Client

	id, err := strconv.Atoi(d.Id())

	if err != nil {
		return err
	}

	client.DeleteServer(id)

	log.Printf("[INFO] Deleting Server")

	return nil
}
