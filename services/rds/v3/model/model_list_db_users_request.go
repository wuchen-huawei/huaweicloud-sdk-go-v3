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

// Request Object
type ListDbUsersRequest struct {
	XLanguage  *string `json:"X-Language,omitempty"`
	InstanceId string  `json:"instance_id"`
	Page       int32   `json:"page"`
	Limit      int32   `json:"limit"`
}

func (o ListDbUsersRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListDbUsersRequest", string(data)}, " ")
}
