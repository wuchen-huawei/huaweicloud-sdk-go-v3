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
type ListInstancesResponse struct {
	// 实例信息。
	Instances *[]QueryInstanceResponse `json:"instances,omitempty"`
	// 总记录数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListInstancesResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListInstancesResponse", string(data)}, " ")
}
