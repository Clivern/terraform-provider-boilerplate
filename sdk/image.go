// Copyright 2019 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package sdk

import (
	"encoding/json"
)

// Image struct
type Image struct {
	ID           int    `json:"id"`
	Slug         string `json:"slug"`
	Name         string `json:"name"`
	Distribution string `json:"distribution"`
	Private      bool   `json:"private"`
	MinDiskSize  int    `json:"min_disk_size"`
}

// LoadFromJSON update object from json
func (s *Image) LoadFromJSON(data []byte) (bool, error) {
	err := json.Unmarshal(data, &s)
	if err != nil {
		return false, err
	}
	return true, nil
}

// ConvertToJSON convert object to json
func (s *Image) ConvertToJSON() (string, error) {
	data, err := json.Marshal(&s)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
