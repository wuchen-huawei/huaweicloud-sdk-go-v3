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
type ListDatabasesRequest struct {
	XLanguage  *string `json:"X-Language,omitempty"`
	InstanceId string  `json:"instance_id"`
	Page       int32   `json:"page"`
	Limit      int32   `json:"limit"`
}

func (o ListDatabasesRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListDatabasesRequest", string(data)}, " ")
}
