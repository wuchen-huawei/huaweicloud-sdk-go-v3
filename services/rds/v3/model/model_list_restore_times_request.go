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
type ListRestoreTimesRequest struct {
	XLanguage  *string `json:"X-Language,omitempty"`
	InstanceId string  `json:"instance_id"`
	Date       *string `json:"date,omitempty"`
}

func (o ListRestoreTimesRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListRestoreTimesRequest", string(data)}, " ")
}
