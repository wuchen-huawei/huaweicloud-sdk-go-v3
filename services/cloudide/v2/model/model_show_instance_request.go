/*
 * cloudide
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
type ShowInstanceRequest struct {
	InstanceId string `json:"instance_id"`
}

func (o ShowInstanceRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowInstanceRequest", string(data)}, " ")
}
