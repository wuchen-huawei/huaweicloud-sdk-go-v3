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
type CreateTemporaryAccessKeyByTokenRequest struct {
	Body *CreateTemporaryAccessKeyByTokenRequestBody `json:"body,omitempty"`
}

func (o CreateTemporaryAccessKeyByTokenRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateTemporaryAccessKeyByTokenRequest", string(data)}, " ")
}
