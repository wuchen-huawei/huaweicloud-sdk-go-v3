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
type SwitchSslResponse struct {
	// 任务ID。
	JobId *string `json:"job_id,omitempty"`
	// SSL开关状态。
	SslOption      *string `json:"ssl_option,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SwitchSslResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"SwitchSslResponse", string(data)}, " ")
}
