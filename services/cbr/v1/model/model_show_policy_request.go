/*
 * CBR
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
type ShowPolicyRequest struct {
	PolicyId string `json:"policy_id"`
}

func (o ShowPolicyRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowPolicyRequest", string(data)}, " ")
}
