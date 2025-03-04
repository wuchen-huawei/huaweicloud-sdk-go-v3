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

// Response Object
type AssociateVaultPolicyResponse struct {
	AssociatePolicy *VaultPolicyResp `json:"associate_policy,omitempty"`
	HttpStatusCode  int              `json:"-"`
}

func (o AssociateVaultPolicyResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"AssociateVaultPolicyResponse", string(data)}, " ")
}
