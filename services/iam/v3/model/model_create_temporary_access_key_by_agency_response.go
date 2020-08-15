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
type CreateTemporaryAccessKeyByAgencyResponse struct {
	Credential *Credential `json:"credential,omitempty"`
}

func (o CreateTemporaryAccessKeyByAgencyResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateTemporaryAccessKeyByAgencyResponse", string(data)}, " ")
}
