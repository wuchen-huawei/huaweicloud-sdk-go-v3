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

type CheckpointPlanCreate struct {
	// 存储库id
	Id string `json:"id"`
	// 存储库名称
	Name string `json:"name"`
	// 备份对象
	Resources *[]CheckpointResourceResp `json:"resources,omitempty"`
	// 备份时跳过的资源列表
	SkippedResources *[]CheckpointCreateSkippedResource `json:"skipped_resources,omitempty"`
}

func (o CheckpointPlanCreate) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CheckpointPlanCreate", string(data)}, " ")
}
