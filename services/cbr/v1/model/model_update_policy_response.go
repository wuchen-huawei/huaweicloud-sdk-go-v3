/*
 * Cbr
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

// Response Object
type UpdatePolicyResponse struct {
	Policy *Policy `json:"policy,omitempty"`
}

func (o UpdatePolicyResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdatePolicyResponse", string(data)}, " ")
}
