/*
 * EPS
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

// Request Object
type ListApiVersionsRequest struct {
}

func (o ListApiVersionsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListApiVersionsRequest", string(data)}, " ")
}
