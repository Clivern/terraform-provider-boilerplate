// Copyright 2019 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package boilerplate

import (
	"context"
	"log"
	"strconv"
	"strings"

	"github.com/clivern/terraform-provider-boilerplate/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

// resourceBoilerplateServer defines a server schema
func resourceBoilerplateServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceBoilerplateServerCreate,
		Read:   resourceBoilerplateServerRead,
		Update: resourceBoilerplateServerUpdate,
		Delete: resourceBoilerplateServerDelete,
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

// resourceBoilerplateServerCreate creates a server
func resourceBoilerplateServerCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client).Client

	server, err := client.CreateServer(context.Background(), &sdk.Server{
		Image:  d.Get("image").(string),
		Name:   d.Get("name").(string),
		Size:   d.Get("size").(string),
		Region: d.Get("region").(string),
	})

	log.Printf("[INFO] Creating Server")

	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(server.ID))
	d.Set("name", server.Name)
	d.Set("size", server.Size)
	d.Set("image", server.Image)
	d.Set("region", server.Region)

	return resourceBoilerplateServerRead(d, m)
}

// resourceBoilerplateServerRead retrieves a server
func resourceBoilerplateServerRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client).Client

	serverName := d.Get("name").(string)

	server, err := client.GetServerByName(context.Background(), serverName)

	log.Printf("[INFO] Getting Server")

	if err != nil {
		// If server not found
		if strings.Contains(err.Error(), "404") {
			d.SetId("")
			return nil
		}

		return err
	}

	d.SetId(strconv.Itoa(server.ID))
	d.Set("name", server.Name)
	d.Set("size", server.Size)
	d.Set("image", server.Image)
	d.Set("region", server.Region)

	return nil
}

// resourceBoilerplateServerUpdate updates a server
func resourceBoilerplateServerUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client).Client

	id, err := strconv.Atoi(d.Id())

	if err != nil {
		return err
	}

	server, err := client.UpdateServer(context.Background(), &sdk.Server{
		ID:     id,
		Image:  d.Get("image").(string),
		Name:   d.Get("name").(string),
		Size:   d.Get("size").(string),
		Region: d.Get("region").(string),
	})

	log.Printf("[INFO] Updating Server")

	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(server.ID))
	d.Set("name", server.Name)
	d.Set("size", server.Size)
	d.Set("image", server.Image)
	d.Set("region", server.Region)

	return resourceBoilerplateServerRead(d, m)
}

// resourceBoilerplateServerDelete deletes a server
func resourceBoilerplateServerDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client).Client

	id, err := strconv.Atoi(d.Id())

	if err != nil {
		return err
	}

	client.DeleteServer(context.Background(), id)

	log.Printf("[INFO] Deleting Server")

	return nil
}
