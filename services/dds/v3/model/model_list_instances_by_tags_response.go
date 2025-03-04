/*
 * DDS
 *
 * API v3
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListInstancesByTagsResponse struct {
	// 实例列表。
	Instances *[]InstanceItem `json:"instances,omitempty"`
	// 总记录数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListInstancesByTagsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListInstancesByTagsResponse", string(data)}, " ")
}
