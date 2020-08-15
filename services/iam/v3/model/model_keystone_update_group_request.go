/*
 * IAM
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
type KeystoneUpdateGroupRequest struct {
	GroupId string                          `json:"group_id"`
	Body    *KeystoneUpdateGroupRequestBody `json:"body,omitempty"`
}

func (o KeystoneUpdateGroupRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"KeystoneUpdateGroupRequest", string(data)}, " ")
}
