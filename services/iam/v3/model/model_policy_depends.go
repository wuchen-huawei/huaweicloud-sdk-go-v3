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
type PolicyDepends struct {
	// 权限所在目录。
	Catalog string `json:"catalog"`
	// 权限展示名。
	DisplayName string `json:"display_name"`
}

func (o PolicyDepends) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"PolicyDepends", string(data)}, " ")
}
