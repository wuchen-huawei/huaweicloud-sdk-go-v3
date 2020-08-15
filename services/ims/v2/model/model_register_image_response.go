/*
 * IMS
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

// Response Object
type RegisterImageResponse struct {
	// 异步任务ID。
	JobId string `json:"job_id,omitempty"`
}

func (o RegisterImageResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"RegisterImageResponse", string(data)}, " ")
}
