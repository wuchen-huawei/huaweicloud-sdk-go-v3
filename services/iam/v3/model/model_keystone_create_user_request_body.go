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

//
type KeystoneCreateUserRequestBody struct {
	User *KeystoneCreateUserOption `json:"user"`
}

func (o KeystoneCreateUserRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"KeystoneCreateUserRequestBody", string(data)}, " ")
}
