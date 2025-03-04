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
type AwakeClusterRequest struct {
	ClusterId   string `json:"cluster_id"`
	ContentType string `json:"Content-Type"`
}

func (o AwakeClusterRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"AwakeClusterRequest", string(data)}, " ")
}
