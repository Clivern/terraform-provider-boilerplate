// Copyright 2019 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package sdk

import (
	"context"
	"fmt"
	"net/http"
)

// Client struct
type Client struct {
	APIKey     string
	APIURL     string
	HTTPClient *HTTP
}

// Init init http client
func (c *Client) Init() {
	c.HTTPClient = &HTTP{}
}

// CreateServer creates a server
func (c *Client) CreateServer(ctx context.Context, server *Server) (*Server, error) {
	data, err := server.ConvertToJSON()

	if err != nil {
		return nil, err
	}

	response, err := c.HTTPClient.Post(
		ctx,
		fmt.Sprintf("%s/server", c.APIURL),
		data,
		map[string]string{},
		map[string]string{"X-AUTH-TOKEN": c.APIKey, "Content-Type": "application/json"},
	)

	if err != nil {
		return nil, err
	}

	statusCode := c.HTTPClient.GetStatusCode(response)

	if statusCode != http.StatusOK && statusCode != http.StatusCreated && statusCode != http.StatusAccepted {
		return nil, fmt.Errorf("Invalid status code %d", statusCode)
	}

	responseBody, err := c.HTTPClient.ToString(response)

	if err != nil {
		return nil, err
	}

	server.LoadFromJSON([]byte(responseBody))

	return server, nil
}

// GetServer retrieves a server
func (c *Client) GetServer(ctx context.Context, id int) (*Server, error) {

	server := &Server{}

	response, err := c.HTTPClient.Get(
		ctx,
		fmt.Sprintf("%s/server/%d", c.APIURL, id),
		map[string]string{},
		map[string]string{"X-AUTH-TOKEN": c.APIKey, "Content-Type": "application/json"},
	)

	statusCode := c.HTTPClient.GetStatusCode(response)

	if statusCode != http.StatusOK {
		return nil, fmt.Errorf("Invalid status code %d", statusCode)
	}

	responseBody, err := c.HTTPClient.ToString(response)

	if err != nil {
		return nil, err
	}

	server.LoadFromJSON([]byte(responseBody))

	return server, nil
}

// GetServerByName retrieves a server by name
func (c *Client) GetServerByName(ctx context.Context, name string) (*Server, error) {

	server := &Server{}

	response, err := c.HTTPClient.Get(
		ctx,
		fmt.Sprintf("%s/server/%s", c.APIURL, name),
		map[string]string{},
		map[string]string{"X-AUTH-TOKEN": c.APIKey, "Content-Type": "application/json"},
	)

	statusCode := c.HTTPClient.GetStatusCode(response)

	if statusCode != http.StatusOK {
		return nil, fmt.Errorf("Invalid status code %d", statusCode)
	}

	responseBody, err := c.HTTPClient.ToString(response)

	if err != nil {
		return nil, err
	}

	server.LoadFromJSON([]byte(responseBody))

	return server, nil
}

// DeleteServer deletes a server
func (c *Client) DeleteServer(ctx context.Context, id int) (bool, error) {

	response, err := c.HTTPClient.Delete(
		ctx,
		fmt.Sprintf("%s/server/%d", c.APIURL, id),
		map[string]string{},
		map[string]string{"X-AUTH-TOKEN": c.APIKey, "Content-Type": "application/json"},
	)

	if err != nil {
		return false, err
	}

	statusCode := c.HTTPClient.GetStatusCode(response)

	if statusCode != http.StatusNoContent {
		return false, fmt.Errorf("Invalid status code %d", statusCode)
	}

	return true, nil
}

// UpdateServer updates a server
func (c *Client) UpdateServer(ctx context.Context, server *Server) (*Server, error) {
	data, err := server.ConvertToJSON()

	if err != nil {
		return nil, err
	}

	response, err := c.HTTPClient.Put(
		ctx,
		fmt.Sprintf("%s/server/%d", c.APIURL, server.ID),
		data,
		map[string]string{},
		map[string]string{"X-AUTH-TOKEN": c.APIKey, "Content-Type": "application/json"},
	)

	statusCode := c.HTTPClient.GetStatusCode(response)

	if statusCode != http.StatusOK {
		return nil, fmt.Errorf("Invalid status code %d", statusCode)
	}

	responseBody, err := c.HTTPClient.ToString(response)

	if err != nil {
		return nil, err
	}

	server.LoadFromJSON([]byte(responseBody))

	return server, nil
}

// GetImageBySlug retrieves an image by a slug
func (c *Client) GetImageBySlug(ctx context.Context, slug string) (*Image, error) {

	image := &Image{}

	response, err := c.HTTPClient.Get(
		ctx,
		fmt.Sprintf("%s/image/%s", c.APIURL, slug),
		map[string]string{},
		map[string]string{"X-AUTH-TOKEN": c.APIKey, "Content-Type": "application/json"},
	)

	statusCode := c.HTTPClient.GetStatusCode(response)

	if statusCode != http.StatusOK {
		return nil, fmt.Errorf("Invalid status code %d", statusCode)
	}

	responseBody, err := c.HTTPClient.ToString(response)

	if err != nil {
		return nil, err
	}

	image.LoadFromJSON([]byte(responseBody))

	return image, nil
}
