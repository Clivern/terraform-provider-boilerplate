// Copyright 2019 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package boilerplate

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

// dataSourceBoilerplateImage defines image schema
func dataSourceBoilerplateImage() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceBoilerplateImageRead,
		Schema: map[string]*schema.Schema{
			"slug": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "slug of the image",
				ValidateFunc: validation.NoZeroValues,
			},
			"name": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "name of the image",
				ValidateFunc: validation.NoZeroValues,
			},
			"distribution": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "distribution of the OS of the image",
			},
			"private": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Is the image private or non-private",
			},
			"min_disk_size": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "minimum disk size required by the image",
			},
		},
	}
}

// dataSourceBoilerplateImageRead retrieves an image
func dataSourceBoilerplateImageRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client).Client

	slug, hasSlug := d.GetOk("slug")

	if !hasSlug {
		return fmt.Errorf("`slug` must be assigned")
	}

	image, err := client.GetImageBySlug(context.Background(), slug.(string))

	log.Printf("[INFO] Getting Image By Slug")

	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(image.ID))
	d.Set("name", image.Name)
	d.Set("slug", image.Slug)
	d.Set("distribution", image.Distribution)
	d.Set("min_disk_size", image.MinDiskSize)
	d.Set("private", image.Private)

	return nil
}
