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
type UpdateInstancePortResponse struct {
	// 任务ID。
	JobId *string `json:"job_id,omitempty"`
	// 实例当前端口号。
	Port           *int32 `json:"port,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o UpdateInstancePortResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateInstancePortResponse", string(data)}, " ")
}
