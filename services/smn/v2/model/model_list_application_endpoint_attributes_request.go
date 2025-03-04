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
type ListApplicationEndpointAttributesRequest struct {
	EndpointUrn string `json:"endpoint_urn"`
}

func (o ListApplicationEndpointAttributesRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListApplicationEndpointAttributesRequest", string(data)}, " ")
}
