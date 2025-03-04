/*
 * BSS
 *
 * Business Support System API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type AutoRenewalResourcesRequest struct {
	ResourceId string `json:"resource_id"`
}

func (o AutoRenewalResourcesRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"AutoRenewalResourcesRequest", string(data)}, " ")
}
