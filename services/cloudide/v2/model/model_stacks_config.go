/*
 * cloudide
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

type StacksConfig struct {
	Attributes *StacksAttribute `json:"attributes,omitempty"`
	Recipe     *Recipe          `json:"recipe,omitempty"`
}

func (o StacksConfig) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"StacksConfig", string(data)}, " ")
}
