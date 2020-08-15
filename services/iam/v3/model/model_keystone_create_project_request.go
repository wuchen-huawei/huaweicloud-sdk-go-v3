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
type KeystoneCreateProjectRequest struct {
	Body *KeystoneCreateProjectRequestBody `json:"body,omitempty"`
}

func (o KeystoneCreateProjectRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"KeystoneCreateProjectRequest", string(data)}, " ")
}
