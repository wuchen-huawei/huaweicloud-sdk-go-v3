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

// Response Object
type CreateAgencyCustomPolicyResponse struct {
	Role *AgencyPolicyRoleResult `json:"role,omitempty"`
}

func (o CreateAgencyCustomPolicyResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateAgencyCustomPolicyResponse", string(data)}, " ")
}
