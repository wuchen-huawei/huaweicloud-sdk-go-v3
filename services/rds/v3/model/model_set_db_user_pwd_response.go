/*
 * RDS
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
type SetDbUserPwdResponse struct {
	// 操作结果。
	Resp           *string `json:"resp,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SetDbUserPwdResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"SetDbUserPwdResponse", string(data)}, " ")
}
