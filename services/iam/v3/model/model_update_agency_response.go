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
type UpdateAgencyResponse struct {
	Agency *AgencyResult `json:"agency,omitempty"`
}

func (o UpdateAgencyResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateAgencyResponse", string(data)}, " ")
}
