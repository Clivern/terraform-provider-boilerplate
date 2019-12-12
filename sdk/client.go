// Copyright 2019 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package sdk

import (
	"net/http"
	"fmt"
)

type Client struct{
	ApiKey string
    ApiUrl string
    HttpClient *HTTP
}

func(c *Client) Init() {
	c.HttpClient = &HTTP{}
}

func(c *Client) CreateServer(server *Server)(*Server, error){
	data, err := server.ConvertToJSON()

	if err != nil {
		return nil, err
	}

	response, err := c.HttpClient.Post(
	    fmt.Sprintf("%s/server", c.ApiUrl),
	   	data,
	    map[string]string{},
	    map[string]string{"X-AUTH-TOKEN": c.ApiKey, "Content-Type": "application/json"},
	)

	if err != nil {
		return nil, err
	}

	statusCode := c.HttpClient.GetStatusCode(response)

	if statusCode != http.StatusOK && statusCode != http.StatusCreated && statusCode != http.StatusAccepted {
		return nil, fmt.Errorf("Invalid status code %d", statusCode)
	}

	responseBody, err := c.HttpClient.ToString(response)

	if err != nil {
		return nil, err
	}

	server.LoadFromJSON([]byte(responseBody))

	return server, nil
}

func(c *Client) GetServer(id int)(*Server, error){

	server := &Server{}

	response, err := c.HttpClient.Get(
	    fmt.Sprintf("%s/server/%d", c.ApiUrl, id),
	    map[string]string{},
	    map[string]string{"X-AUTH-TOKEN": c.ApiKey, "Content-Type": "application/json"},
	)

	statusCode := c.HttpClient.GetStatusCode(response)

	if statusCode != http.StatusOK {
		return nil, fmt.Errorf("Invalid status code %d", statusCode)
	}

	responseBody, err := c.HttpClient.ToString(response)

	if err != nil {
		return nil, err
	}

	server.LoadFromJSON([]byte(responseBody))

	return server, nil
}

func(c *Client) DeleteServer(id int)(bool, error){

	response, err := c.HttpClient.Delete(
	    fmt.Sprintf("%s/server/%d", c.ApiUrl, id),
	    map[string]string{},
	    map[string]string{"X-AUTH-TOKEN": c.ApiKey, "Content-Type": "application/json"},
	)

	if err != nil {
		return false, err
	}

	statusCode := c.HttpClient.GetStatusCode(response)

	if statusCode != http.StatusNoContent {
		return false, fmt.Errorf("Invalid status code %d", statusCode)
	}

	return true, nil
}

func(c *Client) UpdateServer(server *Server)(*Server, error){
	data, err := server.ConvertToJSON()

	if err != nil {
		return nil, err
	}

	response, err := c.HttpClient.Put(
	    fmt.Sprintf("%s/server/%d", c.ApiUrl, server.Id),
	   	data,
	    map[string]string{},
	    map[string]string{"X-AUTH-TOKEN": c.ApiKey, "Content-Type": "application/json"},
	)

	statusCode := c.HttpClient.GetStatusCode(response)

	if statusCode != http.StatusOK {
		return nil, fmt.Errorf("Invalid status code %d", statusCode)
	}

	responseBody, err := c.HttpClient.ToString(response)

	if err != nil {
		return nil, err
	}

	server.LoadFromJSON([]byte(responseBody))

	return server, nil
}
