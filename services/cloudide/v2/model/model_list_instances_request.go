/*
 * cloudide
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

// Request Object
type ListInstancesRequest struct {
	Limit   int64  `json:"limit,omitempty"`
	Offset  int64  `json:"offset,omitempty"`
	Search  string `json:"search,omitempty"`
	SortDir string `json:"sort_dir,omitempty"`
	SortKey string `json:"sort_key,omitempty"`
}

func (o ListInstancesRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListInstancesRequest", string(data)}, " ")
}
