/*
 * EVS
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

type VolumeType struct {
	// 云硬盘类型的ID。
	Id string `json:"id"`
	// 云硬盘类型名称。
	Name       string                `json:"name"`
	ExtraSpecs *VolumeTypeExtraSpecs `json:"extra_specs,omitempty"`
	// 云硬盘类型的描述信息。
	Description string `json:"description,omitempty"`
	// 预留属性。
	QosSpecsId string `json:"qos_specs_id,omitempty"`
	// 预留属性。
	IsPublic bool `json:"is_public,omitempty"`
}

func (o VolumeType) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"VolumeType", string(data)}, " ")
}
