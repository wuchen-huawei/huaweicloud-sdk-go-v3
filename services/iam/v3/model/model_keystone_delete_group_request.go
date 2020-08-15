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
type KeystoneDeleteGroupRequest struct {
	GroupId string `json:"group_id"`
}

func (o KeystoneDeleteGroupRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"KeystoneDeleteGroupRequest", string(data)}, " ")
}
