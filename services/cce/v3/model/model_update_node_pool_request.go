/*
 * CCE
 *
 * CCE开放API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateNodePoolRequest struct {
	ClusterId   string    `json:"cluster_id"`
	NodepoolId  string    `json:"nodepool_id"`
	ContentType string    `json:"Content-Type"`
	ErrorStatus *string   `json:"errorStatus,omitempty"`
	Body        *NodePool `json:"body,omitempty"`
}

func (o UpdateNodePoolRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateNodePoolRequest", string(data)}, " ")
}
