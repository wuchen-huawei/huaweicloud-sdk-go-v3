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
type RestoreTablesRequest struct {
	XLanguage  *string                   `json:"X-Language,omitempty"`
	InstanceId string                    `json:"instance_id"`
	Body       *RestoreTablesRequestBody `json:"body,omitempty"`
}

func (o RestoreTablesRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"RestoreTablesRequest", string(data)}, " ")
}
