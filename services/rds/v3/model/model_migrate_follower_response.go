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
type MigrateFollowerResponse struct {
	// 任务ID
	WorkflowId     *string `json:"workflowId,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o MigrateFollowerResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"MigrateFollowerResponse", string(data)}, " ")
}
