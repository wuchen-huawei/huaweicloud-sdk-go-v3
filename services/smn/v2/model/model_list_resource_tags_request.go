/*
 * SMN
 *
 * SMN Open API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListResourceTagsRequest struct {
	ResourceType string `json:"resource_type"`
	ResourceId   string `json:"resource_id"`
}

func (o ListResourceTagsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListResourceTagsRequest", string(data)}, " ")
}
