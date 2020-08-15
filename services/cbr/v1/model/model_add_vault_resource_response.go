/*
 * Cbr
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
type AddVaultResourceResponse struct {
	// 已添加的资源ID列表
	AddResourceIds []string `json:"add_resource_ids,omitempty"`
}

func (o AddVaultResourceResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"AddVaultResourceResponse", string(data)}, " ")
}
