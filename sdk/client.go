// Copyright 2019 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package sdk

import (
	"fmt"
	"net/http"
)

type Client struct {
	APIKey     string
	APIURL     string
	HTTPClient *HTTP
}

func (c *Client) Init() {
	c.HTTPClient = &HTTP{}
}

// CreateServer creates a server
func (c *Client) CreateServer(server *Server) (*Server, error) {
	data, err := server.ConvertToJSON()

	if err != nil {
		return nil, err
	}

	response, err := c.HTTPClient.Post(
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
func (c *Client) GetServer(id int) (*Server, error) {

	server := &Server{}

	response, err := c.HTTPClient.Get(
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

// DeleteServer deletes a server
func (c *Client) DeleteServer(id int) (bool, error) {

	response, err := c.HTTPClient.Delete(
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
func (c *Client) UpdateServer(server *Server) (*Server, error) {
	data, err := server.ConvertToJSON()

	if err != nil {
		return nil, err
	}

	response, err := c.HTTPClient.Put(
		fmt.Sprintf("%s/server/%d", c.APIURL, server.Id),
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
