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
type CreateAgencyRequest struct {
	Body *CreateAgencyRequestBody `json:"body,omitempty"`
}

func (o CreateAgencyRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateAgencyRequest", string(data)}, " ")
}
