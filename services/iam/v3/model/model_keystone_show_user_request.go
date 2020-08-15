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
type KeystoneShowUserRequest struct {
	UserId string `json:"user_id"`
}

func (o KeystoneShowUserRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"KeystoneShowUserRequest", string(data)}, " ")
}
