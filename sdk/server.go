// Copyright 2019 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package sdk

import (
	"encoding/json"
)

// Server struct
type Server struct {
	ID     int    `json:"id"`
	Image  string `json:"image"`
	Name   string `json:"name"`
	Size   string `json:"size"`
	Region string `json:"region"`
}

// LoadFromJSON update object from json
func (s *Server) LoadFromJSON(data []byte) (bool, error) {
	err := json.Unmarshal(data, &s)
	if err != nil {
		return false, err
	}
	return true, nil
}

// ConvertToJSON convert object to json
func (s *Server) ConvertToJSON() (string, error) {
	data, err := json.Marshal(&s)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
