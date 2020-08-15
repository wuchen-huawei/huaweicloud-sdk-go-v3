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
type KeystoneUpdateUserByAdminRequest struct {
	UserId string                                `json:"user_id"`
	Body   *KeystoneUpdateUserByAdminRequestBody `json:"body,omitempty"`
}

func (o KeystoneUpdateUserByAdminRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"KeystoneUpdateUserByAdminRequest", string(data)}, " ")
}
