/*
 * CBR
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// 备份时跳过的资源
type CheckpointCreateSkippedResource struct {
	// 资源ID
	Id *string `json:"id,omitempty"`
	// 资源类型
	Type *string `json:"type,omitempty"`
	// 资源名称
	Name *string `json:"name,omitempty"`
	// 请参见[错误码](ErrorCode.xml)。
	Code *string `json:"code,omitempty"`
	// 跳过原因，例如：该资源正在备份中。
	Reason *string `json:"reason,omitempty"`
}

func (o CheckpointCreateSkippedResource) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CheckpointCreateSkippedResource", string(data)}, " ")
}
