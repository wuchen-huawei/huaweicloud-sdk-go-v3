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

// Request Object
type ResizeInstanceRequest struct {
	InstanceId string                     `json:"instance_id"`
	Body       *ResizeInstanceRequestBody `json:"body,omitempty"`
}

func (o ResizeInstanceRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ResizeInstanceRequest", string(data)}, " ")
}
